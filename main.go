// main.go
package main

import (
	"fmt"

	"github.com/Variz-h264/Go-Microservice/middleware"
	"github.com/Variz-h264/Go-Microservice/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// สร้าง Middleware CORS และกำหนดค่าในการอนุญาตให้เกิด CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Middleware to log request info
	app.Use(middleware.LogMiddleware)

	// middleware.Ready()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// เพิ่ม Route สำหรับการเรียกข้อมูล User
	app.Get("/users", service.GetUserHandler)

	// สั่งให้เซิร์ฟเวอร์รันและจัดการข้อผิดพลาด
	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
