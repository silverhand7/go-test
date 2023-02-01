package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Silverhand")
	if result != "Hello Silverhand" {
		// t.Fail()
		t.Error("Result must be 'Hello Silverhand'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldGeralt(t *testing.T) {
	result := HelloWorld("Geralt")
	if result != "Hello Geralt" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Geralt'")
	}

	fmt.Println("TestHelloWorldGeralt Done")
}
