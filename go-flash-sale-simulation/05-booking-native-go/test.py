import requests
import time
from concurrent.futures import ThreadPoolExecutor, as_completed

# ปิด Warning ของ requests กรณีเปิด Connection เยอะเกินไป
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# สร้าง Session เดียวแล้วใช้ซ้ำ (สำคัญมาก! ช่วยลดภาระการเปิด Port ของเครื่อง Mac คุณ)
session = requests.Session()

def book_ticket(user_id):
    try:
        # ใช้ 127.0.0.1 แทน localhost ป้องกันปัญหา IPv6 บน Mac
        res = session.post("http://127.0.0.1:8080/book", timeout=5)
        if res.status_code == 200:
            return f"User {user_id}: 🎉 กดทัน! ({res.text.strip()})"
        else:
            return f"User {user_id}: 😭 นก (ของหมด)"
    except Exception as e:
         return f"User {user_id}: ⚠️ เครื่องพังยิงไม่เข้า! ({e})"

def main():
    total_users = 5000   # ลดลงมาเหลือ 5,000 Request (แค่นี้ก็เยอะพอให้ระบบกากๆ ล่มแล้วครับ)
    max_workers = 100    # ให้มีคนงานยิงพร้อมกัน 100 คน (RAM 8GB ไหวสบายๆ)
    
    print(f"🔥 เริ่มเปิดศึกแย่งชิงของ {total_users} คน พร้อมกัน...")
    start_time = time.time()

    success_count = 0
    fail_count = 0

    # ใช้ ThreadPoolExecutor จัดการคิวงานให้เครื่องไม่ค้าง
    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        # โยนงานทั้งหมดลงไปใน Pool
        futures = [executor.submit(book_ticket, i) for i in range(total_users)]
        
        # รอรับผลลัพธ์
        for future in as_completed(futures):
            result = future.result()
            if "🎉" in result:
                success_count += 1
                print(result) # โชว์เฉพาะคนที่จองสำเร็จ จะได้ไม่รกจอ
            else:
                fail_count += 1
                # print(result) # เอาคอมเมนต์ออกถ้าอยากดูคนนก

    print("-" * 30)
    print(f"✅ สรุปผล: กดทัน {success_count} คน | ❌ นก {fail_count} คน")
    print(f"⏱️ ใช้เวลาทั้งหมด: {time.time() - start_time:.2f} วินาที")

if __name__ == "__main__":
    main()