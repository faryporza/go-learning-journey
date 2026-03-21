package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// ใช้ Context และตัวแปร Client เป็น Global เพื่อให้เข้าถึงได้จากทุกฟังก์ชัน
var ctx = context.Background()
var rdb *redis.Client

// สร้าง Struct สำหรับรับข้อมูล JSON ตอนยิง POST
type UserRequest struct {
	Username string `json:"username"`
}

func main() {
	// 1. เชื่อมต่อ Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 2. สร้าง Goroutine (ใส่คำว่า 'go' นำหน้าฟังก์ชัน)
	// เพื่อให้มันทำงานเป็น Background แยกออกไป ไม่บล็อกการทำงานหลัก
	go func() {
		for {
			// ดึงค่า username จาก Redis
			val, err := rdb.Get(ctx, "username").Result()

			if err == redis.Nil {
				fmt.Println("[Loop] ยังไม่มีข้อมูล username ใน Redis")
			} else if err != nil {
				fmt.Println("[Loop] Error:", err)
			} else {
				fmt.Printf("[Loop] Username ปัจจุบันคือ: %s\n", val)
			}

			// สั่งให้หยุดพัก 1 วินาทีก่อนเริ่มรอบใหม่
			time.Sleep(1 * time.Second)
		}
	}()

	// 3. สร้าง API Endpoint รับ POST /user
	http.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		var req UserRequest
		// แปลง JSON ที่ส่งมาให้อยู่ในรูปแบบ Struct
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}

		// บันทึกค่าลง Redis (เลข 0 ด้านหลังสุดคือการบอกว่า 'ไม่มีวันหมดอายุ')
		err := rdb.Set(ctx, "username", req.Username, 10*time.Second).Err()
		if err != nil {
			http.Error(w, "Failed to save to Redis", http.StatusInternalServerError)
			return
		}

		fmt.Printf("\n>>> อัปเดต Username เป็น: %s เรียบร้อยแล้ว! <<<\n\n", req.Username)

		// ตอบกลับ HTTP Status 201 Created
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Username updated!",
		})
	})

	// เพิ่ม API Endpoint สำหรับลบข้อมูล (DELETE /user)
	http.HandleFunc("DELETE /user", func(w http.ResponseWriter, r *http.Request) {
		// สั่งลบ Key ที่ชื่อ "username" ออกจาก Redis
		err := rdb.Del(ctx, "username").Err()
		if err != nil {
			http.Error(w, "ลบข้อมูลไม่สำเร็จ", http.StatusInternalServerError)
			return
		}

		fmt.Println("\n>>> 🗑️ เคลียร์ค่า Username เรียบร้อยแล้ว! <<<\n")

		// ตอบกลับผู้ใช้ว่าทำงานสำเร็จ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Username cleared!",
		})
	})

	// 4. สตาร์ท HTTP Server ที่ Port 8080 (ทำงานควบคู่ไปกับ Loop ด้านบน)
	fmt.Println("API Server เริ่มทำงานที่ Port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
