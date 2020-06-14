// run test: go test -v
package goheroku

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello World"
	if got := HelloWorld(); got != want {
		t.Errorf("HelloWorld() return wrong output: got %q , want %q", got, want)
	}
}
