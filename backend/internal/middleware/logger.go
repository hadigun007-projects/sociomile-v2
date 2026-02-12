package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(logDir string) gin.HandlerFunc {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()
		latency := time.Since(startTime)

		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		logEntry := fmt.Sprintf(
			"[%s] %s | %3d | %13v | %15s | %-7s %s | %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			getStatusColor(statusCode),
			statusCode,
			latency,
			clientIP,
			method,
			path,
			userAgent,
		)

		logFileName := fmt.Sprintf("requests_%s.log", time.Now().Format("2006-01-02"))
		logFilePath := filepath.Join(logDir, logFileName)

		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Failed to open log file: %v\n", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(logEntry); err != nil {
			fmt.Printf("Failed to write to log file: %v\n", err)
		}
	}
}

func getStatusColor(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return "SUCCESS"
	case statusCode >= 300 && statusCode < 400:
		return "REDIRECT"
	case statusCode >= 400 && statusCode < 500:
		return "CLIENT_ERR"
	case statusCode >= 500:
		return "SERVER_ERR"
	default:
		return "UNKNOWN"
	}
}
