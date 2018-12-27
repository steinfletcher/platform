package session

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/steinfletcher/platform/shared/test"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/lib/pq"
)

func TestSelectUserPassword(t *testing.T) {
	username := uuid.NewV4()
	conn := test.DBSetup(func(db *sqlx.DB) {
		q := `INSERT INTO users (username, password) VALUES ('%s', 'abcdef')`
		db.MustExec(fmt.Sprintf(q, username))
	})

	password, err := NewSelectUserPassword(conn)(username.String())

	assert.Nil(t, err)
	assert.Equal(t, "abcdef", password)
}
