package controllers

import (
	"github.com/farypor/go-journey/go-gin-fiber/07-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

// จำลอง Database ไปก่อน (Mock Data)
var users = []models.User{
	{ID: 1, Username: "farypor", Email: "farypor@example.com"},
	{ID: 2, Username: "golang_dev", Email: "dev@example.com"},
}

// Handler สำหรับดึงข้อมูล User ทั้งหมด (GET)
func GetUsers(c *fiber.Ctx) error {
	// Fiber ช่วยแปลง Struct เป็น JSON ให้ง่ายๆ แค่นี้เลย!
	return c.JSON(fiber.Map{
		"success": true,
		"message": "ดึงข้อมูลสำเร็จ",
		"data":    users,
	})
}

// Handler สำหรับเพิ่ม User ใหม่ (POST)
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// อ่าน JSON จาก Body มาใส่ใน Struct
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
		})
	}

	// กำหนด ID ให้ใหม่ (จำลองการทำงานของ Database)
	user.ID = len(users) + 1
	users = append(users, *user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "สร้าง User สำเร็จ",
		"data":    user,
	})
}
