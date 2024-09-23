package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc {
	godotenv.Load()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
	})
}

func CreateToken(id, email string) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractToken(e echo.Context) (string, string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		Id := claims["id"].(string)
		email := claims["email"].(string)
		return Id, email, nil
	}
	return "", "", errors.New("invalid token")
}