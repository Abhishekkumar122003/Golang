package main

import (
	"fmt"

)

func main() {
	//Variable of type uint8 - unsigned integer with 8bytes signed int value which has only positive integer.
	var smallPositiveValue uint8
	fmt.Println(smallPositiveValue , " is the minimum value that uint8 can store");
	smallPositiveValue = 255;
	fmt.Println(smallPositiveValue , " is the maximum value that uint8 can store");
	if smallPositiveValue > 0 {
		fmt.Println("yes it worked")
	}
}