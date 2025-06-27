//2.6_Response_Compression_Middleware 
/*
This middleware compresses HTTP responses using Gzip, 
which reduces payload size and improves performance — especially over slow networks.
*/
package main

import (
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Enable Gzip compression middleware (default: level = -1 which means best speed)
    r.Use(gzip.Gzip(gzip.DefaultCompression))

    r.GET("/", func(c *gin.Context) {
        // This large response will be compressed before being sent to the client
        data := "Hello, this is a response that will be compressed using Gzip.\n"
        c.String(http.StatusOK, data + data + data + data) // repeating to make it compressible
    })

    r.Run(":8080")
}
/*
Test: curl --compressed http://localhost:8080
✅ The server will return a compressed Gzip response, and curl will decompress it before displaying.


Headers seen in response
Content-Encoding: gzip
Content-Type: text/plain; charset=utf-8

why we use the Gzip compression middleware.
 Benefit                   Explanation                                   
 📉 Reduces response size  Often 50–90% smaller                          
 🚀 Faster loading times   Especially useful for APIs and mobile clients 
 🔒 Bandwidth efficient    Saves data on limited networks                
 ✅ Widely supported       All major browsers & HTTP clients support it  



 compression levels
 r.Use(gzip.Gzip(gzip.BestCompression)) // max compression (slower)
r.Use(gzip.Gzip(gzip.BestSpeed))       // less compression, faster
r.Use(gzip.Gzip(gzip.NoCompression))   // disables it

 Step     Description                         
 Install  `github.com/gin-contrib/gzip`       
 Use      `r.Use(gzip.Gzip(...))`             
 Purpose  Compress HTTP responses using Gzip  
 Client   Should send `Accept-Encoding: gzip` 

*/
