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

	// uint16
	var smallPositiveValue2 uint16
	smallPositiveValue2 = 65535
	fmt.Println(smallPositiveValue2, " uint16: 0 to 65535 it can store");
	//uint32
	var   smallPositiveValue3 uint32
	smallPositiveValue3=429467295;
	fmt.Println(smallPositiveValue3, " uint32:0 to 429467295 it ca store");
	//uint64
	var smallPositiveValue4 uint64
	smallPositiveValue4 = 18446744073709551615
	fmt.Println(smallPositiveValue4, " uint64: 0 to 18446744073709551615 it can store");

	var myInt int = 2343243242323424
	myInt = int(smallPositiveValue);
	fmt.Println(myInt , " this is how we do type casting in Golang");
	myInt = int(smallPositiveValue3)
	// myInt++;
	fmt.Println(myInt)

}