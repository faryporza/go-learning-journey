package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

func main() {
	var err error
	// เชื่อมต่อฐานข้อมูล MySQL (ค่าจากใน docker-compose.yml)
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/flashsale")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// ทดสอบว่าต่อ db ติดไหม
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	r := gin.Default()

	// สร้าง Endpoint สำหรับดึงข้อมูล
	r.GET("/products", getProducts)
	r.GET("/orders", getOrders)

	log.Println("Server is running on http://127.0.0.1:3000")
	r.Run(":3000")
}

func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, stock FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Stock); err != nil {
			log.Println("Error scanning product:", err)
			continue
		}
		products = append(products, p)
	}

	if len(products) == 0 { // ให้แสดง [] แทน null หากไม่มีข้อมูล
		products = []Product{}
	}
	c.JSON(http.StatusOK, products)
}

func getOrders(c *gin.Context) {
	rows, err := db.Query("SELECT id, product_id, user_id, created_at FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.ProductID, &o.UserID, &o.CreatedAt); err != nil {
			log.Println("Error scanning order:", err)
			continue
		}
		orders = append(orders, o)
	}

	if len(orders) == 0 { // ให้แสดง [] แทน null หากไม่มีข้อมูล
		orders = []Order{}
	}
	c.JSON(http.StatusOK, orders)
}
