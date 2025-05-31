package main

import (
	"fmt"
	"time"

	. "github.com/fengdotdev/golibs-async/sandbox/draf1/async"
)

func main() {

	someValue := Await(func() int {
		time.Sleep(2 * time.Second) // Simulate a delay
		fmt.Println("Calculating some value...")
		return 42
	})

	fmt.Println("The value is:", someValue)

	someValue.Then(func(value int) {
		fmt.Println("The value is:", value)
	}).Catch(func(err error) {
		fmt.Println("An error occurred:", err)
	}).Finally(func() {
		fmt.Println("Future operation completed.")
	})
}
