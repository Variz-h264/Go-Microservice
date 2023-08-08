package service

import (
	"github.com/Variz-h264/Go-Microservice/core"
	"github.com/gofiber/fiber/v2"
)

type demo struct {
	PID string `json:"id"`
	// Name string `json:"name"`
}

func ApiKeyHandler(c *fiber.Ctx) error {
	// สร้าง Response ใหม่
	response := core.NewResponse()

	// ตั้งค่าข้อมูลใน Response
	demo := demo{
		PID: "1",
		// Name: "John Doe",
	}
	response.SetResponse(demo)

	// ตอบกลับ JSON
	return c.JSON(response.Stack())
}
