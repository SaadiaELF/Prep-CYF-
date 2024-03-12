package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}

// func main() {
//     // Create a new pointer to an int using new
//     ptr := new(int)

//     // Assign a value to the memory location pointed to by ptr
//     *ptr = 42

//     // Print the value and memory address
//     fmt.Printf("Value: %d\n", *ptr)
//     fmt.Printf("Memory address: %p\n", ptr)
// }
// func main() {
//     // Declare a variable
//     x := 42

//     // Declare a pointer to an int and assign the memory address of x to it
//     var ptr *int
//     ptr  = &x

//     // Print the value and memory address of x
//     fmt.Printf("Original value of x: %d\n", x)
//     fmt.Printf("Memory address of x: %p\n", ptr)

//     // Assign a new value to the memory location pointed to by ptr
//     *ptr = 99

//     // Print the updated value of x
//     fmt.Printf("Updated value of x: %d\n", x)
// 	fmt.Printf("Memory address of x: %p\n", ptr)

// }
