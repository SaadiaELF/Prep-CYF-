package main

import "fmt"

func swap(px *int, py *int) (int, int) {
	temp := *px
	*px = *py
	*py = temp
	return *px,*py
}

func main() {
	x := 1
	y := 2

	swap(&x, &y)

	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
}
