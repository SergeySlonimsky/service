package main

import (
	"encoding/json"
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

type JwtCustomClaims struct {
	User string `json:"user"`
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
		Claims:        &JwtCustomClaims{},
		SigningKey:    key,
		SigningMethod: "RS256",
		TokenLookup:   "header:token",
		AuthScheme:    "x-auth",
	}))

	e.GET("/", func(c echo.Context) error {
		user, err := getJWTUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()} )
		}
		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getJWTUser(c echo.Context) (User, error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JwtCustomClaims)
	user := User{}
	err := json.Unmarshal([]byte(claims.User), &user)
	return user, err
}
