package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	// This is usually used to do something before run the test, for example connect to DB, prepare the data, etc.
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after unit test
	fmt.Println("AFTER UNIT TEST")
}

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

func TestHelloWorldAssertEqual(t *testing.T) {
	result := HelloWorld("Jon")
	assert.Equal(t, "Hello Jon", result, "They should be equal")
	fmt.Println("Assert done") // it will be executed because assert above using Fail
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Snow")
	require.Equal(t, "Hello Snow", result, "Result Must be Hello Snow")
	fmt.Println("Will not be executed if it fails") // using FailNow()
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		// using t.Skip() to skip the test on certain condition
		t.Skip("Can't run on MacOs")
	}

	result := HelloWorld("Jon")
	assert.Equal(t, "Hello Jon", result, "They should be equal")
}
