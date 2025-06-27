/*
Request_Time_Tracking_Middleware

This middleware tracks how long it takes to process each request 
— from when it enters the middleware to when the response is sent back.
*/
package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)

// RequestTimeTrackingMiddleware tracks and logs how long each request takes
func RequestTimeTrackingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now() // Record the start time

        c.Next() // Proceed to the next handler

        duration := time.Since(startTime) // Calculate time taken
        log.Printf("🔁 %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
    }
}

func main() {
    r := gin.Default()

    // Apply the request time tracking middleware globally
    r.Use(RequestTimeTrackingMiddleware())

    // Sample route
    r.GET("/", func(c *gin.Context) {
        // Simulate processing delay
        time.Sleep(100 * time.Millisecond)
        c.String(http.StatusOK, "Request processed!")
    })

    r.Run(":8080")
}

/* 
Client
  ↓
RequestTimeTrackingMiddleware
  - Record start time
  - c.Next() → actual handler runs
  - Record duration
  - Log time taken
  ↓
Handler → returns "Request processed!"



 Purpose                   Description                                 
 ⏱ Performance monitoring  Detect slow routes or bottlenecks           
 🧪 Debugging              See where time is being spent in processing 
 📈 Logging for metrics    Export to tools like Prometheus, ELK, etc.  



To log time into a file instead of console?
A Prometheus-compatible version?
Add timing info into response headers (e.g., X-Response-Time)?
*/
