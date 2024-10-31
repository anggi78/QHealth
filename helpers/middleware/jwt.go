package middleware

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

// func JwtMiddleware() echo.MiddlewareFunc {
//     godotenv.Load()
//     return echojwt.WithConfig(echojwt.Config{
//         SigningKey:    []byte(os.Getenv("JWT_SECRET")),
//         SigningMethod: "HS256",
//         SuccessHandler: func(c echo.Context) {
//             _, _, doctorId, err := ExtractToken(c)
//             if err == nil {
//                 c.Set("doctorId", doctorId)
//             }
//         },
//     })
// }


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
    user, ok := e.Get("user").(*jwt.Token)
    if !ok || !user.Valid {
        return "", "", errors.New("invalid token")
    }

    claims, ok := user.Claims.(jwt.MapClaims)
    if !ok {
        return "", "", errors.New("invalid token claims")
    }

    Id, okId := claims["id"].(string)
    email, okEmail := claims["email"].(string)
    if !okId || !okEmail {
        return "", "", errors.New("invalid token data")
    }

    return Id, email, nil
}

// func ExtractToken(e echo.Context) (string, string, string, error) {
//     user, ok := e.Get("user").(*jwt.Token)
//     if !ok || !user.Valid {
//         return "", "", "", errors.New("invalid token")
//     }

//     claims, ok := user.Claims.(jwt.MapClaims)
//     if !ok {
//         return "", "", "", errors.New("invalid token claims")
//     }

//     Id, okId := claims["id"].(string)
//     email, okEmail := claims["email"].(string)
//     doctorId, okDoctorId := claims["doctorId"].(string) 
    
//     if !okId || !okEmail || !okDoctorId {
//         return "", "", "", errors.New("invalid token data")
//     }

//     return Id, email, doctorId, nil
// }
