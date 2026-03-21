# stress_test.py
import threading
import requests

def book_ticket(user_id):
    res = requests.post("http://localhost:8080/book")
    if res.status_code == 200:
        print(f"User {user_id}: 🎉 กดทัน!")
    else:
        print(f"User {user_id}: 😭 นก (ของหมด)")

# สร้าง 100 Threads รุมยิงพร้อมกัน
threads = []
for i in range(100):
    t = threading.Thread(target=book_ticket, args=(i,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()