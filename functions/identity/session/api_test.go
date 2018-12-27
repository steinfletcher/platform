package session_test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/steinfletcher/api-test"
	"github.com/steinfletcher/platform/functions/identity/app"
	"github.com/steinfletcher/platform/shared/crypto"
	"github.com/steinfletcher/platform/shared/test"
	"net/http"
	"testing"
)

func TestCreateSession(t *testing.T) {
	username := uuid.NewV4()
	password := "abcdef"

	test.DBSetup(func(db *sqlx.DB) {
		q := `INSERT INTO users (username, password) VALUES ('%s', '%s')`
		db.MustExec(fmt.Sprintf(q, username, crypto.HashPassword(password)))
	})

	apitest.New(app.New().Router).
		Post("/v1/session").
		Body(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)).
		Expect(t).
		Status(http.StatusOK).
		CookiePresent("Session-Token").
		End()
}

func TestCreateSession_BadRequestIfWrongPassword(t *testing.T) {
	username := uuid.NewV4()
	password := "abcdef"

	test.DBSetup(func(db *sqlx.DB) {
		q := `INSERT INTO users (username, password) VALUES ('%s', '%s')`
		db.MustExec(fmt.Sprintf(q, username, crypto.HashPassword(password)))
	})

	apitest.New(app.New().Router).
		Post("/v1/session").
		Body(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, "wrong password")).
		Expect(t).
		Status(http.StatusBadRequest).
		CookieNotPresent("Session-Token").
		Body(`{"code":"BAD_REQUEST","description":"Invalid credentials"}`).
		End()
}
