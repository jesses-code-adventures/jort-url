package urls

import (
	"math/rand"
	"slices"
)

const ALL_RANDOM_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generatePathOption() string {
	lengths := []int{1, 2, 3, 4, 5, 6, 7}
	targetLength := lengths[rand.Intn(len(lengths))]
	randomString := make([]byte, targetLength)
	for i := range randomString {
		randomString[i] = ALL_RANDOM_CHARS[rand.Intn(len(ALL_RANDOM_CHARS))]
	}
	return string(randomString)
}

func GetRandomPath(staticPaths []string) string {
	for {
		path := generatePathOption()
		if slices.Contains(staticPaths, path) {
			continue
		}
		return path
	}
}
