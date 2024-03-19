package main

import (
	"fmt"
)

func tumbili() {
	x := []int{1, 2, 3, 4}
	y := make([]int, len(x))
	copy(y, x)
	fmt.Println(y)
	fmt.Println(y)
	y = append(y, 10, 12, 45)
	fmt.Println(y)
	fmt.Println(x)

}

func main() {
	tumbili()
}
