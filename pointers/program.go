package main

import "fmt"

func swap(x *int, y *int) (int, int) {
	temp := *x
	*x = *y
	*y = temp
	return *x,*y
}

func main() {
	x := 1
	y := 2

	swap(&x, &y)

	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
}
