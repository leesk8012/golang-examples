package main

import (
	"fmt"
)

func call_by_value(value string) {
	fmt.Println("========================")
	fmt.Printf("Value: %s\n", value)
	fmt.Printf("Call by Value Address : %p\n", &value)	
}

func call_by_reference(reference * string) {
	fmt.Println("========================")
	fmt.Printf("Value: %s\n", *reference)
	fmt.Printf("Call by Reference Address : %p\n", reference)
}

func main() {
	test_string := "Just check the difference between 'Call by Value' and 'Call by Reference'"
	fmt.Printf("Original Address : %p\n", &test_string)

	call_by_reference(&test_string)
	call_by_value(test_string)
}