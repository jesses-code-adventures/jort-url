package database

type UserExistsError struct{}

func (e UserExistsError) Error() string {
	return "user already exists"
}

func (db *Database) userExists(username []byte) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM user WHERE username = ?`, username).Scan(&count)
	return count > 0, err
}

func (db *Database) CreateUser(username, password []byte) error {
	exists, err := db.userExists(username)
	if err != nil {
		return err
	}
	if exists {
		return UserExistsError{}
	}
	encodedPassword, err := db.PasswordHandler.Hash(password)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO user (username, password) VALUES (?, ?)`, username, encodedPassword)
	return err
}
