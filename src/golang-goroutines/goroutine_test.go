package golanggoroutines

import (
	"fmt"
	"testing"
)

func HelloWord() {
	fmt.Println("Hello")
}

func testingCreateGoroutine(t *testing.T) {
	go HelloWord()
	fmt.Println("Hello")
}
