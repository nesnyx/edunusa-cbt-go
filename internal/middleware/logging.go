// internal/middleware/simple_logging.go
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ANSI Color Codes untuk terminal
const (
	ColorGreen   = "\033[92m"
	ColorWhite   = "\033[97m"
	ColorYellow  = "\033[93m"
	ColorRed     = "\033[91m"
	ColorBlue    = "\033[94m"
	ColorMagenta = "\033[95m"
	ColorCyan    = "\033[96m"
	ColorReset   = "\033[0m"
)

// SimpleLoggingMiddleware adalah middleware untuk log sederhana dan berwarna.
func SimpleLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Waktu mulai request
		start := time.Now()

		// Proses request ke handler selanjutnya
		c.Next()

		// Waktu selesai request
		latency := time.Since(start)

		// Kumpulkan data-data penting
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Tentukan warna berdasarkan status code
		var statusColor string
		switch {
		case statusCode >= 200 && statusCode < 300:
			statusColor = ColorGreen
		case statusCode >= 400 && statusCode < 500:
			statusColor = ColorYellow
		case statusCode >= 500:
			statusColor = ColorRed
		default:
			statusColor = ColorWhite
		}

		var methodColor string
		switch method {
		case "GET":
			methodColor = ColorBlue
		case "POST":
			methodColor = ColorCyan
		case "PUT":
			methodColor = ColorYellow
		case "DELETE":
			methodColor = ColorRed
		case "PATCH":
			methodColor = ColorMagenta
		default:
			methodColor = ColorWhite
		}

		// Format output log
		logMessage := fmt.Sprintf("[GIN-LOG] |%s %3d %s| %13v | %15s |%s %-7s %s %s",
			statusColor, statusCode, ColorReset,
			latency,
			clientIP,
			methodColor, method, ColorReset,
			path,
		)

		// Cetak log ke console
		fmt.Println(logMessage)
	}
}
