package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// สร้าง Struct เพื่อกำหนดรูปแบบข้อมูล JSON ที่จะตอบกลับ
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	// 1. กำหนด Route และ Handler Function
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// เช็คว่าเป็น GET Method หรือเปล่า
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// กำหนด Header ว่าเราจะตอบกลับเป็น JSON
		w.Header().Set("Content-Type", "application/json")

		// สร้างข้อมูลจาก Struct
		data := Response{
			Message: "Hello from Go API!",
			Status:  "success",
		}

		// แปลง Struct เป็น JSON แล้วส่งกลับไป
		json.NewEncoder(w).Encode(data)
	})

	// 2. สั่งให้ Server เริ่มทำงานที่ Port 8080
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)

	// ถ้า Server พังให้ Log error ออกมา
	if err != nil {
		log.Fatal(err)
	}
}
