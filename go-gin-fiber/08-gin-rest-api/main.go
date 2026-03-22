package main

import (
	"fmt"

	"github.com/farypor/go-journey/go-gin-fiber/08-gin-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. สร้าง Router พื้นฐาน (มี Logger กับ Recovery มาให้ในตัวเลย)
	r := gin.Default()

	// 2. เรียกใช้ Routes ที่เราแยกไฟล์ไว้
	routes.SetupRoutes(r)

	// 3. สตาร์ท Server (Gin มักจะใช้ Port 8080 เป็นค่าเริ่มต้น)
	fmt.Println("🍸 Gin Server is running on port 8080...")
	r.Run(":8080")
}
