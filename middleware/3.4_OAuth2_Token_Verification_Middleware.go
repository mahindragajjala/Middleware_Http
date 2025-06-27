// 3.4_OAuth2_Token_Verification_Middleware
/*
OAuth2 Token Verification Middleware checks if the incoming request 
has a valid access token (usually in the header) before allowing 
access to protected resources.

Think of Google Drive API:
When your app tries to access a user's files, Google requires an OAuth2 access token.
This token is verified using middleware before any business logic is executed.
*/

/* 
🔄 Common Flow
Client sends a request with an Authorization: Bearer <access_token> header.
Middleware intercepts request.
Token is validated (e.g., signature, expiry, issuer).
If valid → continue.
If invalid → return 401 Unauthorized.
*/

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Dummy token validator — replace with real OAuth2 logic (e.g., introspection)
func isValidOAuth2Token(token string) bool {
	// In real-world: validate with OAuth2 provider or JWKS
	return token == "valid-oauth2-token"
}

func OAuth2Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if !isValidOAuth2Token(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
func main() {
	r := gin.Default()

	// Apply to specific route or group
	protected := r.Group("/secure", OAuth2Middleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Secure data accessed"})
		})
	}

	r.Run(":8080")
}


//test:curl -H "Authorization: Bearer valid-oauth2-token" http://localhost:8080/secure/profile
//# Output: {"message": "Secure data accessed"}
