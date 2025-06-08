package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        string `param:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	userDB := make(map[string]*User)
	e := echo.New()
	e.PUT("/users/:id", updateUser(userDB))
	e.Logger.Fatal(e.Start(":8080"))
}

func updateUser(userDB map[string]*User) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user User
		err := c.Bind(&user)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		dbUser, ok := userDB[user.ID]
		if !ok {
			return c.NoContent(http.StatusNotFound)
		}

		dbUser.FirstName = user.FirstName
		dbUser.LastName = user.LastName

		return c.JSON(http.StatusOK, dbUser)
	}
}
