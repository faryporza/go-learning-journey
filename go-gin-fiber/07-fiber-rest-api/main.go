package main

import (
	"fmt"
	"log"

	"github.com/farypor/go-journey/go-gin-fiber/07-fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. สร้าง App Fiber
	app := fiber.New()

	// 2. ใส่ Middleware ช่วยโชว์ Log สวยๆ ใน Terminal เวลามีคนเรียก API
	app.Use(logger.New())

	// 3. เรียกใช้ SetupRoutes เพื่อเชื่อมต่อเส้นทางทั้งหมด
	routes.SetupRoutes(app)

	// 4. สตาร์ท Server ที่ Port 3000 (ปกติ Fiber นิยมใช้ 3000 ครับ)
	fmt.Println("Fiber Server is running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
