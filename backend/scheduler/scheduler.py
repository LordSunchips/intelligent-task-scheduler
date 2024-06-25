from flask import Flask, request, jsonify
import requests
import time
from queue import PriorityQueue
from worker import WorkerNode

app = Flask(__name__)

class Scheduler:
    def __init__(self, api_url):
        self.api_url = api_url
        self.tasks = PriorityQueue()
        self.worker_nodes = []


    def fetch_tasks(self):
        # `GET /api/tasks` fetches all tasks from the API and adds them to the queue
        response = requests.get(f'{self.api_url}/api/tasks')
        
        if response.status_code == 200:
            tasks = response.json()
            for task in tasks:
                self.tasks.put((task['priority'], task))

    def get_next_task(self):
        return self.tasks.get()[1] if not self.tasks.empty() else None
    
    def get_best_worker(self, task):
        # linear search for the best worker
        best_worker, closest_resources = None, float('inf')
        for worker in self.worker_nodes:
            resources_available = worker.get_available_resources()
            if resources_available == task["resource_needed"]:
                return worker
            elif task["resource_needed"] < resources_available < closest_resources:
                best_worker = worker
                closest_resources = resources_available
        return best_worker

    def add_worker(self, worker):
        self.worker_nodes.append(worker)
                
    
    def assign_task(self, task):
        worker_node = self.get_best_worker(task)
        worker_node.execute_task(task)

scheduler = Scheduler(api_url='http://localhost:5001')
scheduler.fetch_tasks()

# `GET /request_task` gets the next task from the queue
@app.route('/request_task', methods=['GET'])
def request_task():
    task = scheduler.get_next_task()
    if task:
        return jsonify(task), 200
    return jsonify({"message": "No tasks available"}), 200

if __name__ == "__main__":
    app.run(port=5000)
    
# `POST /connect_worker` adds a worker node to the list of available workers
@app.route('/connect_worker', methods=['POST'])
def connect_worker():
    response = request.get_json()
    worker_id, total_resources = response['id'], response['total_resources']
    if worker_id in [worker.id for worker in scheduler.worker_nodes]:
        return jsonify({"message": "Worker already connected"}), 200
    scheduler.add_worker(WorkerNode(worker_id, total_resources, scheduler.api_url))
    