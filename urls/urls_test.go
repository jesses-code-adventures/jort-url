package urls

import (
	"fmt"
	"testing"
)

func TestCreateRandomString(t *testing.T) {
	randomString := getRandomString(10)
	if len(randomString) != 10 {
		t.Errorf("Expected random string of length 10, got %d", len(randomString))
	}
	fmt.Println("random string -> ", randomString)
}
