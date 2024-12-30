package jwt

import (
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
