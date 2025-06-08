package transport

import (
	"errors"
	"fmt"
	"net/http"

	appErr "example/internal/errors"
	"example/internal/users"

	"github.com/labstack/echo/v4"
)

type HTTPConfig struct {
	Port         int
	UsersService users.Service
}

type HTTPServer struct {
	e    *echo.Echo
	port int
}

func NewHTTPServer(config *HTTPConfig) *HTTPServer {
	e := echo.New()
	registerHTTPRoutes(e, config.UsersService)
	return &HTTPServer{
		e:    e,
		port: config.Port,
	}
}

func (server *HTTPServer) Start() {
	server.e.Logger.Fatal(server.e.Start(fmt.Sprintf(":%d", server.port)))
}

func registerHTTPRoutes(e *echo.Echo, usersService users.Service) {
	e.PUT("/users/:id", updateUser(usersService))
}

type updateUserRequest struct {
	ID        string `param:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type updateUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func updateUser(usersService users.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request updateUserRequest
		err := c.Bind(&request)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		resp, err := usersService.UpdateUser(users.UpdateUserRequest{
			ID:        request.ID,
			FirstName: request.FirstName,
			LastName:  request.LastName,
		})

		if err != nil {
			return c.NoContent(MapErrorToHTTPStatus(err))
		}
		return c.JSON(http.StatusOK, updateUserResponse{
			ID:        resp.ID,
			FirstName: resp.FirstName,
			LastName:  resp.LastName,
		})
	}
}

func MapErrorToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, appErr.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, appErr.ErrBadRequest):
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
