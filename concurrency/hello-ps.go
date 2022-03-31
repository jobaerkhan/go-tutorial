package main

import (
	"fmt"
	"time"
)

func main() {

	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	func() {
		fmt.Println("Pluralsight")
	}()
}
