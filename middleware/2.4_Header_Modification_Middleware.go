//2.4_Header_Modification_Middleware
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// HeaderModificationMiddleware adds or modifies headers in the response
func HeaderModificationMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Add custom headers BEFORE calling the next handler
        c.Writer.Header().Set("X-App-Version", "1.0.0")
        c.Writer.Header().Set("X-Security-Policy", "strict")
        
        // Continue to the next middleware or final handler
        c.Next()

        // Optionally modify headers AFTER the handler (not recommended for most headers)
    }
}

func main() {
    r := gin.Default()

    // Apply the header middleware globally
    r.Use(HeaderModificationMiddleware())

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Header modified successfully"})
    })

    r.Run(":8080")
}

/*
What This Middleware Does
It sets custom headers before the request reaches your route handler.

These headers will be part of the final response sent back to the client.

 Header Type          Purpose                                    
 `X-App-Version`      Inform client of current API version       
 `X-Security-Policy`  Add CSP, security policies                 
 `Cache-Control`      Control caching behavior                   
 `Access-Control-*`   Support CORS (Access-Control-Allow-Origin) 
 `X-Request-ID`       Add a unique request ID for tracing logs   


Client
  ↓
HeaderModificationMiddleware → adds headers
  ↓
Final Handler → returns JSON
  ↓
Response (includes the added headers)

*/










