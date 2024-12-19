package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type Claims struct {
	User_id   string `json:"user_id"`
	User_Type string `json:"user_type"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id string, user_Type string) (string, string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	refreshTokenTime := time.Now().Add(7 * 24 * time.Hour)

	var user_type string

	claims := &Claims{
		User_id:   user_id,
		User_Type: user_type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SECRET_KEY)

	refreshClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshTokenTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(SECRET_KEY)

	if err != nil {
		panic(err)
	}

	return signedToken, signedRefreshToken, nil

}
