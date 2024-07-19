package database

type InvalidCredentialsError struct{}

func (e InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

func (db *Database) Authenticate(userId int, jwt string) error {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM user WHERE id = ? AND session_jwt = ?`, userId, jwt).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return InvalidCredentialsError{}
	}
	return nil
}
