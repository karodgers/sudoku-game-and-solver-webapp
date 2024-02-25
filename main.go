package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {

	x := 'z'
	z01.PrintRune(x)
	z01.PrintRune('\n')
	y := "2373"
	fmt.Println(atoi(y))
}

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
