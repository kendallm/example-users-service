package users

import (
	appErr "example/internal/errors"
	"fmt"
)

type UpdateUserRequest struct {
	ID        string
	FirstName string
	LastName  string
}

type UpdateUserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

type User struct {
	ID        string
	FirstName string
	LastName  string
}

type Service struct {
	DB map[string]*User
}

func (s *Service) UpdateUser(request UpdateUserRequest) (UpdateUserResponse, error) {
	user, ok := s.DB[request.ID]
	if !ok {
		return UpdateUserResponse{}, fmt.Errorf("unable to find user in db: %w", appErr.ErrNotFound)
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}

	if request.LastName != "" {
		user.LastName = request.LastName
	}

	return UpdateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
