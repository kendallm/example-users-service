package main

import (
	"example/internal/transport"
	"example/internal/users"
)

type User struct {
	ID        string `param:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	userDB := make(map[string]*users.User)
	usersService := users.Service{
		DB: userDB,
	}

	httpService := transport.NewHTTPServer(&transport.HTTPConfig{
		Port:         8080,
		UsersService: usersService,
	})

	httpService.Start()
}
