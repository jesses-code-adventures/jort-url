package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type IncorrectPasswordError struct{}

func (e IncorrectPasswordError) Error() string {
	return "incorrect password"
}

type UserNotFoundError struct{}

func (e UserNotFoundError) Error() string {
	return "user not found"
}

func createJwtToken(username string, timestamp time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte(fmt.Sprintf("%s-%d", username, timestamp.UnixNano())))
}

func (db *Database) setUserSessionJwt(username string, jwt string) (int, error) {
	var userId int
	err := db.QueryRow(`UPDATE user SET session_jwt = ? WHERE username = ? RETURNING id`, jwt, username).Scan(&userId)
	return userId, err
}

// returns a jwt that should be set as a cookie on success
// can throw a UserNotFoundError on non-existent users
// can throw an IncorrectPasswordError on password mismatches
// all other errors are genuine errors
func (db *Database) Login(username, password string) (int, string, error) {
	var encodedPassword []byte
	err := db.QueryRow(`SELECT password FROM user WHERE username = ?`, username).Scan(&encodedPassword)
	if err == sql.ErrNoRows {
		return 0, "", UserNotFoundError{}
	}
	if err != nil {
		return 0, "", err
	}
	verified, err := db.PasswordHandler.Verify([]byte(password), encodedPassword)
	if err != nil {
		return 0, "", err
	}
	if !verified {
		return 0, "", IncorrectPasswordError{}
	}
	jwt, err := createJwtToken(username, time.Now())
	if err != nil {
		return 0, "", err
	}
	userId, err := db.setUserSessionJwt(username, jwt)
	if err != nil {
		return 0, "", err
	}
	return userId, jwt, nil
}
