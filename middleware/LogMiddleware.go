package middleware

import (
	"fmt"
	"io"
	"log"
	"net"
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
	Region     string
}

func LogMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Get the client IP using X-Real-IP or X-Forwarded-For header
	ip := c.IP()

	reqInfo := RequestInfo{
		IP:         ip,
		Host:       c.Hostname(),
		Method:     c.Method(),
		RequestURI: c.Request().URI().String(),
		UserAgent:  c.Get("User-Agent"),
	}

	c.Locals("request_info", reqInfo)

	// Call the next middleware or handler
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

	// Get the region based on IP
	region, err := getRegionFromIP(reqInfo.IP)
	if err != nil {
		log.Println("Failed to get region from IP:", err)
		region = "unknown"
	}
	reqInfo.Region = region

	// แสดงผลเพิ่มเติมว่าใช้ PC หรือ Mobile
	if strings.Contains(reqInfo.UserAgent, "Mobile") {
		fmt.Println("Device:", green("Mobile")) // แสดงผลเป็นสี Magenta
	} else {
		fmt.Println("Device:", green("PC")) // แสดงผลเป็นสี Magenta
	}
	fmt.Printf("Region: %s\n", strings.ToLower(reqInfo.Region))
	fmt.Printf("<================>\n\n")

	return nil
}

func getRegionFromIP(ip string) (string, error) {
	ips, err := net.LookupIP(ip)
	if err != nil {
		return "", err
	}
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			// Use a public GeoIP API to get region from the IP
			resp, err := http.Get("https://ipapi.co/" + ipv4.String() + "/region")
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			buf := new(strings.Builder)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return "", err
			}

			return buf.String(), nil
		}
	}
	return "", fmt.Errorf("No valid IPv4 address found")
}
