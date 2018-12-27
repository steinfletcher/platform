package session

import "github.com/jmoiron/sqlx"

type SelectPassword func(username string) (string, error)

const query = `SELECT password FROM users WHERE username=$1 LIMIT 1`

func NewSelectUserPassword(db *sqlx.DB) SelectPassword {
	return func(username string) (string, error) {
		var password []string
		err := db.Select(&password, query, username)
		if err != nil {
			return "", err
		}
		if len(password) == 0 {
			return "", nil
		}
		return password[0], nil
	}
}
