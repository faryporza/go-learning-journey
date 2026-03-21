package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// เรียกใช้ฟังก์ชันจาก handlers.go ได้เลย เพราะอยู่ package main เหมือนกัน
	mux.HandleFunc("GET /items", getItems)
	mux.HandleFunc("POST /items", createItem)
	mux.HandleFunc("DELETE /items/{id}", deleteItem)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
