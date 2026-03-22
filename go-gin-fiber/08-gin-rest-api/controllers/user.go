package controllers

import (
	"net/http"

	"github.com/farypor/go-journey/go-gin-fiber/08-gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Username: "farypor", Email: "farypor@example.com"},
	{ID: 2, Username: "gin_master", Email: "master@example.com"},
}

// Handler สำหรับดึงข้อมูล (GET)
func GetUsers(c *gin.Context) {
	// Gin ใช้ c.JSON และต้องใส่ HTTP Status Code กำกับด้วยเสมอ
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ดึงข้อมูลสำเร็จ",
		"data":    users,
	})
}

// Handler สำหรับเพิ่มข้อมูล (POST)
func CreateUser(c *gin.Context) {
	var user models.User

	// ใช้ ShouldBindJSON เพื่อดึงข้อมูลจาก Body มาใส่ Struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "สร้าง User สำเร็จ",
		"data":    user,
	})
}
