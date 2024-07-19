package database

import (
	"github.com/golang-jwt/jwt/v5"
)

type IncorrectPasswordError struct{}

func (e IncorrectPasswordError) Error() string {
	return "incorrect password"
}

func createJwtToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte("secret"))
}

func (db *Database) setUserSessionJwt(username []byte, jwt string) error {
	_, err := db.Exec(`UPDATE user SET session_jwt = ? WHERE username = ?`, jwt, username)
	return err
}

// returns a jwt that should be set as a cookie on success
// can throw an IncorrectPasswordError on password mismatches
// all other errors are genuine errors
func (db *Database) Login(username, password []byte) (string, error) {
	var encodedPassword []byte
	err := db.QueryRow(`SELECT password FROM user WHERE username = ?`, username).Scan(&encodedPassword)
	if err != nil {
		return "", err
	}
	verified, err := db.PasswordHandler.Verify(password, encodedPassword)
	if err != nil {
		return "", err
	}
	if !verified {
		return "", IncorrectPasswordError{}
	}
	jwt, err := createJwtToken()
	if err != nil {
		return "", err
	}
	err = db.setUserSessionJwt(username, jwt)
	return jwt, nil
}
