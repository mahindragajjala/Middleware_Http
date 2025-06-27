//2.3_Logging_Middleware_Example
package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware logs the incoming request method, path, and processing time
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Log request method and path before processing
        log.Printf("Started %s %s", c.Request.Method, c.Request.URL.Path)

        c.Next() // Pass control to the next handler

        // After handler completes
        duration := time.Since(startTime)
        log.Printf("Completed %s in %v", c.Request.URL.Path, duration)
    }
}

func main() {
    r := gin.Default()

    // Apply LoggingMiddleware globally
    r.Use(LoggingMiddleware())

    // Example route
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to Gin Logging Example!")
    })

    log.Println("🚀 Server is running at http://localhost:8080")
    r.Run(":8080")
}
