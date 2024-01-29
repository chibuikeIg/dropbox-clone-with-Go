package generator

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaim struct {
	UserID string
	jwt.RegisteredClaims
}

func CreateAccessToken(user_id string) (string, error) {

	// specify custom claims
	claims := CustomClaim{
		user_id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	// generate a jwt token using a signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// get signed string using secret key
	return token.SignedString([]byte(os.Getenv("APP_KEY")))

}

func ParseToken(t string) (*CustomClaim, error) {

	// parse claims and generate token
	token, err := jwt.ParseWithClaims(t, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {

		// check if token is using the same signing method
		if token.Method != jwt.SigningMethodHS256 {

			return nil, fmt.Errorf("invalid token")

		}

		return []byte(os.Getenv("APP_KEY")), nil
	})

	if err == nil && token.Valid {

		claims, _ := token.Claims.(*CustomClaim)

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token. error %v", err)
}
