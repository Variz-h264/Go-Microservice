package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

type RequestInfo struct {
	IP         string
	Host       string
	Method     string
	RequestURI string
	UserAgent  string
}

func LogMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	reqInfo := RequestInfo{
		IP:         c.IP(),
		Host:       c.Hostname(),
		Method:     c.Method(),
		RequestURI: c.Request().URI().String(),
		UserAgent:  c.Get("User-Agent"),
	}

	c.Locals("request_info", reqInfo)

	err := c.Next()

	statusText := http.StatusText(c.Response().StatusCode())

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("<================>\n")
	fmt.Printf("%s\n", green("Request"))
	fmt.Printf("%s: %s\n", cyan("IP"), reqInfo.IP)
	fmt.Printf("%s: %s\n", cyan("Host"), reqInfo.Host)
	fmt.Printf("%s: %s\n", cyan("Method"), reqInfo.Method)
	fmt.Printf("%s: %s\n", cyan("RequestURI"), reqInfo.RequestURI)
	fmt.Printf("%s: %s\n", cyan("UserAgent"), reqInfo.UserAgent)

	fmt.Printf("%s %d %s\n", green("Response Status:"), c.Response().StatusCode(), statusText)
	fmt.Printf("%s: %v\n", green("Response Time"), time.Since(start))

	// ตรวจสอบ User-Agent ที่ส่งมาใน request header
	isMobile := strings.Contains(reqInfo.UserAgent, "Mobile")

	// แสดงผลเพิ่มเติมว่าใช้ PC หรือ Mobile
	if isMobile {
		fmt.Println("Device:", green("Mobile")) // แสดงผลเป็นสี Magenta
		fmt.Printf("<================>\n\n")

	} else {
		fmt.Println("Device:", green("PC")) // แสดงผลเป็นสี Magenta
		fmt.Printf("<================>\n\n")

	}

	return err
}
