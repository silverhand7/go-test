package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Silverhand")
	if result != "Hello Silverhand" {
		panic("Result is not 'Hello Silverhand'")
	}
}
