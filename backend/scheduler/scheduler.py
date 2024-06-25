from flask import Flask, request, jsonify
import requests
import time
from queue import PriorityQueue

app = Flask(__name__)

class Scheduler:
    def __init__(self, api_url):
        self.api_url = api_url
        self.tasks = PriorityQueue()

    # `GET /api/tasks` gets all the tasks in the queue
    def fetch_tasks(self):
        response = requests.get(f'{self.api_url}/api/tasks')
        if response.status_code == 200:
            tasks = response.json()
            for task in tasks:
                self.tasks.put((task['priority'], task))


    def get_next_task(self):
        return self.tasks.get()[1] if not self.tasks.empty() else None

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