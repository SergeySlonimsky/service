package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type jwtCustomClaims struct {
	User
	jwt.StandardClaims
}

func main() {
	/*	if err := DbConnect(); err != nil {
		log.Fatal(err)
	}*/
	key, err := GetPublic()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &jwtCustomClaims{},
		SigningKey:    key,
		SigningMethod: "RS256",
		TokenLookup:   "header:token",
		AuthScheme:    "x-auth",
	}))

	e.GET("/", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		log.Println(claims.Email)
		return c.String(http.StatusOK, "Hello")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
