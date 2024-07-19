package database

func (db *Database) Logout(username []byte) error {
	_, err := db.Exec(`UPDATE user SET session_jwt = NULL WHERE username = ?`, username)
	return err
}
