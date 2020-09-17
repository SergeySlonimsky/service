package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/project-mate/user-sdk"
	"log"
	"net/http"
)

func main() {
	/*	if err := DbConnect(); err != nil {
		log.Fatal(err)
	}*/
	jwtMiddleware, err := user_sdk.UserJWTMiddleware()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(jwtMiddleware)

	e.GET("/", func(c echo.Context) error {
		user, err := user_sdk.GetJWTUser(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()} )
		}
		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
