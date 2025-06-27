//3.2_JWTbased_Authentication_Middleware

/*
What Is JWT Authentication Middleware?
            It’s a Gin middleware that:
                      Extracts the JWT from the Authorization header
                      Validates the signature, expiry, and claims
                      Blocks invalid tokens and passes valid ones to the next handler
                      Stores user info in context for use in downstream routes
*/
package main

import (
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key") // Store in env in prod

// Custom claims struct
type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.RegisteredClaims
}

// Middleware: Validate JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, "Bearer ")
        if len(parts) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            c.Abort()
            return
        }

        tokenStr := parts[1]

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Pass claims to next handlers
        c.Set("username", claims.Username)
        c.Set("role", claims.Role)

        c.Next()
    }
}

// Generate JWT for login (normally after DB check)
func GenerateToken(username, role string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    claims := &Claims{
        Username: username,
        Role:     role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/login", func(c *gin.Context) {
        var creds struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.BindJSON(&creds); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        // Fake auth (replace with DB check)
        if creds.Username != "admin" || creds.Password != "1234" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        token, err := GenerateToken(creds.Username, "admin")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    })

    // Protected route group
    authGroup := r.Group("/secure")
    authGroup.Use(JWTAuthMiddleware())
    {
        authGroup.GET("/profile", func(c *gin.Context) {
            username := c.GetString("username")
            role := c.GetString("role")
            c.JSON(http.StatusOK, gin.H{
                "message":  "Welcome to secure profile!",
                "username": username,
                "role":     role,
            })
        })
    }

    r.Run(":8080")
}


/*
testing : 
  curl -X POST http://localhost:8080/login \
 -H "Content-Type: application/json" \
 -d '{"username":"admin", "password":"1234"}'

returns :
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6Ikp..."}


access secure route :
curl -H "Authorization: Bearer <TOKEN>" \
 http://localhost:8080/secure/profile
*/
