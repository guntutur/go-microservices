package github

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := GetAuthorizationHeader("myToken")
	assert.EqualValues(t, "token myToken", header)
}

// this is the result of TestNotDefer execution :
/*
=== RUN   TestNotDefer
functions body
--- PASS: TestNotDefer (0.00s)
PASS
*/
func TestNotDefer(t *testing.T) {
	fmt.Println("functions body")
}

// this is the result of TestDefer execution :
/*
=== RUN   TestDefer
functions body
3
2
1
--- PASS: TestDefer (0.00s)
PASS
*/
func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("functions body")
}
// now you know the difference whats defer, how its work, and when to use it
