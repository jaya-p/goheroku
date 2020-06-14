// build: go build driver/cli/main.go
// run: go run driver/cli/main.go
package main

import (
	"fmt"

	"github.com/jaya-p/goheroku"
)

func main() {
	fmt.Println(goheroku.HelloWorld())
}
