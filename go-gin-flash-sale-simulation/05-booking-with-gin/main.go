package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import driver แบบไม่เรียกใช้ตรงๆ
)

var db *sql.DB

func initDB() {
	var err error
	// เชื่อมต่อ MySQL (ใช้ User/Pass ตามที่ตั้งใน docker-compose)
	dsn := "root:root@tcp(127.0.0.1:3306)/flashsale"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// ตั้งค่า Connection Pool ให้รับโหลดได้นิดหน่อย
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
}

func bookItem(c *gin.Context) {
	// จำลองชื่อ User แบบสุ่ม (ใช้เวลาปัจจุบันมารวม)
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	// 1. อ่านค่าสต็อกปัจจุบัน
	var stock int
	err := db.QueryRow("SELECT stock FROM products WHERE id = 1").Scan(&stock)
	if err != nil {
		c.String(http.StatusInternalServerError, "Database error")
		return
	}

	// 2. เช็คว่าของหมดหรือยัง
	if stock > 0 {
		// ⚠️ จำลองความหน่วงของ Server (เพื่อให้เกิด Race Condition ชัดเจนขึ้น)
		time.Sleep(50 * time.Millisecond)

		// 3. ถ้ายังมีของ ให้ลดสต็อกลง 1
		_, err = db.Exec("UPDATE products SET stock = stock - 1 WHERE id = 1")
		if err != nil {
			c.String(http.StatusInternalServerError, "Update error")
			return
		}

		// 4. บันทึกประวัติการสั่งซื้อ
		_, err = db.Exec("INSERT INTO orders (product_id, user_id) VALUES (1, ?)", userID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Insert error")
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("✅ จองสำเร็จ! (User: %s)\n", userID))
	} else {
		c.String(http.StatusBadRequest, "❌ ของหมดแล้ว!\n")
	}
}

func main() {
	initDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/book", bookItem)

	fmt.Println("API Server (No Redis + Gin) running on port 8080...")
	r.Run(":8080")
}
