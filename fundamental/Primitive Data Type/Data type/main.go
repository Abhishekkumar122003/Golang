package main

import "fmt"

func main() {

	var myString string // here we just create the data type and not assigned any value means it has "zero" value stroed in string case it's an empty string 

	fmt.Println(myString);  // op:- "" empty string
	myString = "wellcome to Golang"
	fmt.Println((myString))
}