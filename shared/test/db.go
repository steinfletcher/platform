package test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/steinfletcher/platform/shared/crypto"
	"os"
)

type User struct {
	Username     string
	Password     string
	SessionToken string
}

func init() {
	connStr := "host=localhost port=15432 user=postgres password=postgres dbname=platform sslmode=disable"
	err := os.Setenv("DB_ADDR", connStr)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("SESSION_SECRET", testSessionSecret)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("UI_DOMAIN", "localhost:3000")
	if err != nil {
		panic(err)
	}
}

const testSessionSecret = "notTheRealPasswordLolz"

func DBSetup(setup func(db *sqlx.DB)) *sqlx.DB {
	addr := os.Getenv("DB_ADDR")
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic(err)
	}
	setup(db)
	return db
}

func UserWithSession() *User {
	username := uuid.NewV4().String()
	password := "abcdef"

	DBSetup(func(db *sqlx.DB) {
		q := `INSERT INTO users (username, password) VALUES ('%s', '%s')`
		db.MustExec(fmt.Sprintf(q, username, crypto.HashPassword(password)))
	})

	sessionToken, err := crypto.NewCreateJWT(testSessionSecret)(username)
	if err != nil {
		panic(err)
	}

	return &User{
		Username:     username,
		Password:     password,
		SessionToken: sessionToken,
	}
}
