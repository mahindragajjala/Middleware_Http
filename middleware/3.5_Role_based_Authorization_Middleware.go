//3.5_Role_based_Authorization_Middleware

/* 
Role-based authorization middleware ensures that only users with specific 
roles (like admin, user, or manager) can access certain routes or perform specific actions.

🔐 
Authentication ≠ Authorization
Authentication: Who are you?
Authorization: Are you allowed to do this?
*/
//Middleware Code: RoleAuthorizationMiddleware.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// This would normally be extracted during JWT parsing
type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// AcceptableRoles → e.g., "admin", "user"
func RoleAuthorizationMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole") // set earlier by auth middleware
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
			return
		}

		// Check if user's role is allowed
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
	}
}
//JWT Parsing Middleware (Extracts role into context)
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		claims := &CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Save role into Gin context for next middlewares
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
//✅ Route Setup Example
func main() {
	r := gin.Default()

	// Apply JWT authentication middleware globally or per group
	protected := r.Group("/admin")
	protected.Use(JWTMiddleware())
	protected.Use(RoleAuthorizationMiddleware("admin"))
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!"})
		})
	}

	user := r.Group("/user")
	user.Use(JWTMiddleware())
	user.Use(RoleAuthorizationMiddleware("user", "admin"))
	{
		user.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "User Profile Accessed"})
		})
	}

	r.Run(":8080")
}
