package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("./test1.txt")

	if err != nil {
		fmt.Println("This is the error code:", err)
	}

}
