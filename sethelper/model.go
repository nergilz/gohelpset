package sethelper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaimsModel
type JWTClaimsModel struct {
	jwt.StandardClaims
	ID   string
	Data string
	Time time.Time
}
