import requests
from queue import PriorityQueue
import time

class Scheduler:
    def __init__(self, api_url):
        self.api_url = api_url
        self.tasks = PriorityQueue()
    
    def fetch_tasks(self):
        # Fetch tasks from the Golang API
        response = requests.get(f"{self.api_url}/api/tasks")
        if response.status_code != 200:
            print("Failed to fetch tasks")
            return []
        tasks = response.json()
        for task in tasks:
            self.tasks.put((task["priority"], task))
    
    