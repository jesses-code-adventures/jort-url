package urls

import (
	"math/rand"
)

const ALL_RANDOM_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GetRandomPath() string {
	lengths := []int{1, 2, 3, 4, 5, 6, 7}
	targetLength := lengths[rand.Intn(len(lengths))]
	randomString := make([]byte, targetLength)
	for i := range randomString {
		randomString[i] = ALL_RANDOM_CHARS[rand.Intn(len(ALL_RANDOM_CHARS))]
	}
	return string(randomString)
}
