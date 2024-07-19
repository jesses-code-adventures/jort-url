package urls

import (
	"math/rand"
)

const ALL_RANDOM_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func getRandomString(length int) string {
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = ALL_RANDOM_CHARS[rand.Intn(len(ALL_RANDOM_CHARS))]
	}
	return string(randomString)
}

// returns only the pathname component of the shortened url
func shortenUrl(url string) string {
	// TODO: check for an existing url in the db
	lengths := []int{4, 5, 6, 7}
	targetLength := lengths[rand.Intn(len(lengths))]
	exists := true
	var randomString string
	for exists {
		randomString = getRandomString(targetLength)
		// TODO: check the shortened url doesn't already exist in db
		exists = false
	}
	// TODO: save the shortened url to the db
	return randomString
}
