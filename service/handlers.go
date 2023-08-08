package service

import (
	"github.com/Variz-h264/Go-Microservice/core"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUserHandler(c *fiber.Ctx) error {
	// สร้าง Response ใหม่
	response := core.NewResponse()

	// ตั้งค่าข้อมูลใน Response
	user := User{
		ID:   "1",
		Name: "John Doe",
	}
	response.SetResponse(user)

	// ตอบกลับ JSON
	return c.JSON(response.Stack())
}
