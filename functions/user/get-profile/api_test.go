package get_profile_test

import (
	"github.com/steinfletcher/api-test"
	"github.com/steinfletcher/platform/functions/user/app"
	"github.com/steinfletcher/platform/shared/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetProfile_Success(t *testing.T) {
	user := test.UserWithSession()

	apitest.New(app.New().Router).
		Get("/v1/user").
		Cookies(map[string]string{"Session-Token": user.SessionToken}).
		Expect(t).
		Status(http.StatusOK).
		JSONPath("$.username", func(values interface{}) {
			assert.Equal(t, user.Username, values)
		}).
		End()
}

func TestGetProfile_UnauthorizedIfNoSession(t *testing.T) {
	apitest.New(app.New().Router).
		Get("/v1/user").
		Expect(t).
		Status(http.StatusUnauthorized).
		Body(`{"code": "UNAUTHORIZED", "description":"You must be logged in to perform this action"}`).
		End()
}
