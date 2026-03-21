package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. สร้าง Context (บังคับใช้ในไลบรารีเวอร์ชันใหม่ๆ ของ Go เพื่อจัดการเรื่อง Timeout)
	ctx := context.Background()

	// 2. สร้าง Client เพื่อเชื่อมต่อไปที่ Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // ที่อยู่ของ Redis
		Password: "",               // รหัสผ่าน (รันผ่าน Docker พื้นฐานจะไม่มีรหัส)
		DB:       0,                // ใช้ Database ช่องที่ 0 (Default)
	})

	// 3. ทดสอบการเชื่อมต่อด้วยคำสั่ง Ping
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("เชื่อมต่อ Redis ไม่สำเร็จ: %v", err)
	}
	fmt.Println("เชื่อมต่อ Redis สำเร็จ! ตอบกลับ:", pong)

	// ==========================================
	// ทดลองใช้งาน Set และ Get
	// ==========================================

	// 4. Set ข้อมูล: กำหนด Key="username", Value="farypor" และตั้งเวลาหมดอายุ 10 วินาที
	err = rdb.Set(ctx, "username", "farypor", 10*time.Second).Err()
	if err != nil {
		log.Fatalf("บันทึกข้อมูลไม่สำเร็จ: %v", err)
	}
	fmt.Println("บันทึกข้อมูล 'username' ลง Redis เรียบร้อยแล้ว (จะหมดอายุใน 10 วินาที)")

	// 5. Get ข้อมูลออกมาดู
	val, err := rdb.Get(ctx, "username").Result()
	if err != nil {
		log.Fatalf("ดึงข้อมูลไม่สำเร็จ: %v", err)
	}
	fmt.Printf("ดึงข้อมูล 'username' ได้ค่าเป็น: %s\n", val)
}
