package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1, 4)
	fmt.Printf("Length starts out as %d with a capacity of %d\n",
		len(mySlice), cap(mySlice))

}
