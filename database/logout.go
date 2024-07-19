package database

func (db *Database) Logout(userId int, token string) error {
	_, err := db.Exec(`UPDATE user SET session_jwt = NULL WHERE id = ? and session_jwt = ?`, userId, token)
	return err
}
