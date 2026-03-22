package routes

import (
	"github.com/farypor/go-journey/go-gin-fiber/07-fiber-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// สร้าง Group จัดหมวดหมู่ API (URL จะขึ้นต้นด้วย /api/v1)
	api := app.Group("/api/v1")

	// ผูก URL เข้ากับ Controller
	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
}
