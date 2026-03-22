package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	db  *sql.DB
	rdb *redis.Client
	ctx = context.Background()
)

func initConnections() {
	// 1. เชื่อมต่อ MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/flashsale"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("MySQL Error:", err)
	}

	// 2. เชื่อมต่อ Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 3. **ตั้งค่าเริ่มต้น:** เอาสต็อก 10 ชิ้น ไปใส่ใน Redis ซะก่อน!
	err = rdb.Set(ctx, "iphone_stock", 10, 0).Err()
	if err != nil {
		log.Fatal("Redis Set Error:", err)
	}
	fmt.Println("📦 เอาสินค้า 10 ชิ้น เข้าโกดัง Redis เรียบร้อย!")
}

func bookItemWithRedis(c *gin.Context) {
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	// ==========================================
	// ⚡️ พระเอกอยู่ตรงนี้: ตัดสต็อกใน Redis ก่อนเลย! ⚡️
	// ==========================================
	stockLeft, err := rdb.Decr(ctx, "iphone_stock").Result()
	if err != nil {
		c.String(http.StatusInternalServerError, "Redis error")
		return
	}

	// ถ้าตัดสต็อกแล้วได้เลขติดลบ (-1, -2, -3...) แปลว่าของหมดไปแล้ว
	if stockLeft < 0 {
		c.String(http.StatusBadRequest, "❌ ของหมดแล้ว!\n")
		return // จบการทำงานทันที ไม่ต้องไปยุ่งกับ MySQL!
	}

	// ==========================================
	// ถ้าผ่านมาถึงตรงนี้ได้ แปลว่า "ได้ของชัวร์ๆ" (มีแค่ 10 คนเท่านั้นที่รอดมาได้)
	// ==========================================

	time.Sleep(50 * time.Millisecond) // จำลองความหน่วงของ Server นิดนึง

	// อัปเดตยอดจริงใน MySQL
	_, err = db.Exec("UPDATE products SET stock = stock - 1 WHERE id = 1")
	if err != nil {
		log.Println("MySQL Update Error:", err)
	}

	// บันทึกออเดอร์ลง MySQL
	_, err = db.Exec("INSERT INTO orders (product_id, user_id) VALUES (1, ?)", userID)
	if err != nil {
		log.Println("MySQL Insert Error:", err)
	}

	c.String(http.StatusOK, fmt.Sprintf("✅ จองสำเร็จ! (User: %s) คิวที่รอด: เหลือ %d ชิ้น\n", userID, stockLeft))
}

func main() {
	initConnections()
	defer db.Close()

	r := gin.Default()
	r.POST("/book", bookItemWithRedis)

	fmt.Println("🚀 API Server (With Redis Gatekeeper + Gin) running on port 8080...")
	r.Run(":8080")
}
