package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-async/sandbox/draf1/async"
)

func main() {
	async.Async(func() {
		fmt.Println("Hello, World!")
	})
}
