package users_test

import (
	"testing"

	appErr "example/internal/errors"
	"example/internal/users"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {

	t.Run("should update user when user exists", func(t *testing.T) {

		mockDB := map[string]*users.User{
			"test_user_id": &users.User{
				ID:        "test_user_id",
				FirstName: "Test",
				LastName:  "User",
			},
		}

		request := users.UpdateUserRequest{
			ID:        "test_user_id",
			FirstName: "UpdatedTest",
			LastName:  "UpdatedUser",
		}
		expectedResponse := users.UpdateUserResponse{
			ID:        "test_user_id",
			FirstName: "UpdatedTest",
			LastName:  "UpdatedUser",
		}
		s := users.Service{DB: mockDB}

		if response, err := s.UpdateUser(request); assert.NoError(t, err) {
			user := mockDB["test_user_id"]
			assert.Equal(t, expectedResponse, response)
			assert.Equal(t, "UpdatedTest", user.FirstName)
			assert.Equal(t, "UpdatedUser", user.LastName)
		}
	})

	t.Run("should return ErrNotFound when user not found", func(t *testing.T) {
		mockDB := map[string]*users.User{
			"test_user_id": &users.User{
				ID:        "test_user_id",
				FirstName: "Test",
				LastName:  "User",
			},
		}

		request := users.UpdateUserRequest{
			ID:        "bad_user_id",
			FirstName: "UpdatedTest",
			LastName:  "UpdatedUser",
		}

		s := users.Service{DB: mockDB}

		if response, err := s.UpdateUser(request); assert.ErrorIs(t, err, appErr.ErrNotFound) {
			assert.Equal(t, users.UpdateUserResponse{}, response)
		}
	})
}
