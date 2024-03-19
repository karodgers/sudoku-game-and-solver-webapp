package main

import (
	"fmt"
	"strconv"

	"github.com/01-edu/z01"
)

// func main() {

// 	x := 'z'
// 	z01.PrintRune(x)
// 	z01.PrintRune('\n')
// 	y := "2373"
// 	fmt.Println(atoi(y))
// 	slice1()
// 	zero()

// }

func atoi(s string) int {

	result := 0
	sign := 1

	for len(s) > 0 && (s[0] == '-' || s[0] == '+') {

		for s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')

		if digit < 0 || digit > 9 {
			return 0
		}
		result = result*10 + digit
	}
	result = result * sign
	return result

}

func slice1() {
	sliceNames := []string{
		"mark",
		"madison",
		"mane",
		"brook",
	}
	sliceNames = append(sliceNames, "molly")
	fmt.Println(sliceNames)

	for i, v := range sliceNames {
		if i == 0 {
			for _, j := range v {
				z01.PrintRune(j)
			}
			z01.PrintRune('\n')
		}
		// for i, e := range v {
		// 	if i == 0 {
		// 		z01.PrintRune(e)
		// 	}

		// }

	}
	// vop:= []int{}
	// slice2 := make([]int, 20)
	// for i:=0; i<20;i++{
	// 	return slice2[i]
	// }
	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] =
				strconv.Itoa(row) + "," +
					strconv.Itoa(column)
		}

	}
	fmt.Println(table)

}

func zero() {
	// Creating a 2D array with 5 rows and 6 columns
	rows, cols := 5, 6
	array2D := make([][]int, rows)
	for i := range array2D {
		array2D[i] = make([]int, cols)
	}

	// Filling the array with some values (for demonstration)
	counter := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			array2D[i][j] = counter
			counter++
		}
	}

	// Printing the 2D array
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", array2D[i][j])
		}
		fmt.Println()
	}
}
