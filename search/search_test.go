package search

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "Hello"
	actual := Hello()
	if actual != expected {
		fmt.Println("Test Failed")
	}
}
