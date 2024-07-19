package database

import (
	"database/sql"
	"time"
)

type UrlData struct {
	Id            int
	CreatedAt     time.Time
	UserId        int
	Url           string
	ShortenedPath string
	Clicks        int
}

func (db *Database) GetUrls(userId int) ([]UrlData, error) {
	rows, err := db.Query(`SELECT id, created_at, url, shortened_path, clicks FROM url WHERE user_id = ? ORDER BY created_at desc`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var urls []UrlData
	for rows.Next() {
		var url UrlData
		err := rows.Scan(&url.Id, &url.CreatedAt, &url.Url, &url.ShortenedPath, &url.Clicks)
		if err != nil {
			return nil, err
		}
		url.UserId = userId
		urls = append(urls, url)
	}
	return urls, nil
}

func (db *Database) IncrementUrlClicks(userId int, shortenedPath string) error {
	_, err := db.Exec(`UPDATE url SET clicks = clicks + 1 WHERE user_id = ? AND shortened_path = ?`, userId, shortenedPath)
	return err
}

// returns nil with nil error if the shortened path is not found
func (db *Database) GetExistingShortenedPath(userId int, url string) (*string, error) {
	var shortenedPath string
	err := db.QueryRow(`SELECT shortened_path FROM url WHERE user_id = ? AND url = ?`, userId, url).Scan(&shortenedPath)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &shortenedPath, nil
}

// returns nil with nil error if the url is not found
func (db *Database) GetUrl(shortenedPath string) (*string, error) {
	var url string
	err := db.QueryRow(`SELECT url FROM url WHERE shortened_path = ?`, shortenedPath).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (db *Database) ShortPathHasBeenUsed(shortenedPath string) (bool, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM url WHERE shortened_path = ?`, shortenedPath).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (db *Database) CreateUrl(userId int, url string, shortenedPath string) error {
	ctime := time.Now()
	_, err := db.Exec(`INSERT INTO url (user_id, created_at, url, shortened_path) VALUES (?, ?, ?, ?)`, userId, ctime, url, shortenedPath)
	if err != nil {
		return err
	}
	return nil
}
