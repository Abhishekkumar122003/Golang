package main

import (
	"fmt"
	"net/smtp"
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

	// uint16
	var smallPositiveValue2 uint16
	smallPositiveValue2 = 65535
	fmt.Println(smallPositiveValue2, " :0 to 65535 it can store");
	//uint32
	var   smallPositiveValue3 uint32
	smallPositiveValue3=429467295;
	fmt.Println(smallPositiveValue3, " :0 to 429467295 it ca store");
	//uint64
	var smallPositiveValue4 uint64
	smallPositiveValue4 = 18446744073709551615
	fmt.Println(smallPositiveValue4, " 0 to 18446744073709551615 it can store");
}