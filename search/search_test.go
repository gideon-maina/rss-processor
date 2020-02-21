package search

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "From the search module."
	actual := Hello()
	if actual != expected {
		fmt.Println("Test Failed")
	}
}
