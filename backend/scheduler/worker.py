import requests
import time

SCHEDULER_URL = 'http://localhost:5000'

class WorkerNode:
    def __init__(self, id, total_resources, scheduler_url):
        self.id = id
        self.total_resources = total_resources
        self.resources_available = total_resources
        self.scheduler_url = scheduler_url

    def is_available(self):
        return self.resources_available == self.total_resources

    def connect_to_scheduler(self, scheduler_url):
        payload = {"total_resources": self.total_resources, "id": self.id}
        response = requests.post(f'{scheduler_url}/connect_worker', json=payload)
        if response.status_code == 200:
            print(f"Worker {self.id} connected to scheduler")
        else:
            print(f"Worker {self.id} failed to connect to scheduler")
            

    def execute_task(self, task):
        if task['resource_needed'] <= self.resources_available:
            self.resources_available -= task['resource_needed']
            print(f"Worker {self.id} is executing task {task['title']} with {task['resource_needed']} resources")
            time.sleep(task["resource_needed"]) # simulate task execution
            self.resources_available += task['resource_needed']
            print(f"Worker {self.id} completed task {task['title']}")

    def request_task(self):
        response = requests.get(f'{self.scheduler_url}/request_task')
        if response.status_code == 200:
            task = response.json()
            if "resource_needed" in task:
                self.execute_task(task)
            else:
                print(f"Worker {self.id} received message: {task['message']}")
        else:
            print(f"Worker {self.id} failed to request task")

if __name__ == "__main__":
    worker = WorkerNode(1, total_resources=10, scheduler_url=SCHEDULER_URL)
    worker.connect_to_scheduler(SCHEDULER_URL)
    