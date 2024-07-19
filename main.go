package main

import (
	"github.com/jesses-code-adventures/jort-url/database"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
