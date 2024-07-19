package main

import (
	"github.com/jesses-code-adventures/jort-url/database"
)

func main() {
	_, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
}
