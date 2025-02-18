package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test/database"
	_ "test/docs"
	"test/routes"
)



// @title Todo List Social Media Post Management API
// @version 1.0
// @description A ToDo List Management API for Social Media Posts
// @host localhost:8080
// @BasePath /api/
func main() {

	// Membuat instance gin untuk menjalankan server
	r := gin.Default()

	// Menghubungkan ke database dan memanggil setup routes yang telah dibuat
	database.ConnectDB()
	routes.SetupRoutes(r)

	// Menggunakan ginSwagger dari module swaggo untuk menampilkan dokumentasi swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Menjalankan server pada port 8080
	r.Run(":8080")
}
