package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a JWT for a given user
// @Description Generates a JSON Web Token (JWT) using the provided secret key and claims.
// @Param secretKey string The secret key used to sign the token.
// @Param claims jwt.MapClaims The claims to be embedded in the token.
// @Return string The signed JWT as a string.
// @Return error An error if the token generation fails.
func GenerateToken(secretKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT and returns the claims if valid
// @Description Validates a JSON Web Token (JWT) using the provided secret key.
// Returns the claims if the token is valid, or an error if validation fails.
// @Param secretKey string The secret key used to validate the token's signature.
// @Param tokenString string The JWT string to be validated.
// @Return jwt.MapClaims The claims extracted from the valid token.
// @Return error An error if the token is invalid or validation fails.
func ValidateToken(secretKey string, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// AuthMiddleware is a middleware that validates JWT tokens in HTTP requests
// @Description Middleware to validate JSON Web Tokens (JWT) for protected routes.
// @Param secretKey string The secret key used to validate the token.
// @Return gin.HandlerFunc The middleware function for Gin.
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Ensure the token is prefixed with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Validate the token
		tokenString := parts[1]
		claims, err := ValidateToken(secretKey, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Store claims in the context for further use
		c.Set("claims", claims)

		// Continue to the next middleware/handler
		c.Next()
	}
}
