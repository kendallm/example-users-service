package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var ()

func TestUpdateUser(t *testing.T) {

	t.Run("should update user when user exists", func(t *testing.T) {

		mockDB := map[string]*User{
			"test_user_id": &User{
				ID:        "test_user_id",
				FirstName: "Test",
				LastName:  "User",
			},
		}
		userJSON := `{"first_name":"UpdatedTest", "last_name":"UpdatedUser"}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("test_user_id")
		h := updateUser(mockDB)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var response User
			if !assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response), "unable to unmarshal body to user") {
				t.FailNow()
			}
			assert.Equal(t, mockDB["test_user_id"], &response)
			assert.Equal(t, "UpdatedTest", mockDB["test_user_id"].FirstName)
			assert.Equal(t, "UpdatedUser", mockDB["test_user_id"].LastName)
		}
	})

	t.Run("should return 404 when user not found", func(t *testing.T) {

		mockDB := map[string]*User{
			"test_user_id": &User{
				ID:        "test_user_id",
				FirstName: "Test",
				LastName:  "User",
			},
		}
		userJSON := `{"first_name":"UpdatedTest", "last_name":"UpdatedUser"}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("bad_user_id")
		h := updateUser(mockDB)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("should return 400 when request is invalid", func(t *testing.T) {

		mockDB := map[string]*User{
			"test_user_id": &User{
				ID:        "test_user_id",
				FirstName: "Test",
				LastName:  "User",
			},
		}
		userJSON := `{"first_name":"UpdatedTest", "last_name":300}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("bad_user_id")
		h := updateUser(mockDB)

		if assert.NoError(t, h(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
