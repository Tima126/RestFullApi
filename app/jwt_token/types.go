
package jwt
import "github.com/golang-jwt/jwt/v5"


// Общий тип Claims -> jwt
type CustomClaims struct {
	Login string `json:"login"`
	Role string `json:"role,omitempty"`
	jwt.RegisteredClaims
}