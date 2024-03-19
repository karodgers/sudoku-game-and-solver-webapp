// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func first() {
// 	samples := []string{"hello", "apple_π!"}
// outer:
// 	for _, sample := range samples {
// 		for i, r := range sample {
// 			fmt.Println(i, r, string(r))
// 			if r == 'l' {
// 				continue outer
// 			}
// 		}
// 		fmt.Println()
// 	}

// }

// func prob() {
// 	samples := []string{"hello", "apple_π!"}
// 	for _, sample := range samples {
// 		for i, r := range sample {
// 			fmt.Println(i, r, string(r))
// 		}
// 		fmt.Println()
// 	}

// 	evenVals := []int{2, 4, 6, 8, 10}
// 	for i := 1; i < len(evenVals)-1; i++ {
// 		fmt.Println(i, evenVals[i])
// 	}

// 	words := []string{"a", "cow", "smile", "gopher",
// 		"octopus", "anthropologist"}
// 	for _, word := range words {
// 		switch size := len(word); size {
// 		case 1, 2, 3, 4:
// 			fmt.Println(word, "is a short word!")
// 		case 5:
// 			wordLen := len(word)
// 			fmt.Println(word, "is exactly the right length:", wordLen)
// 		case 6, 7, 8, 9:
// 		default:
// 			fmt.Println(word, "is a long word!")
// 		}
// 		fmt.Println()
// 	}

// 	names := []string{
// 		"ma",
// 		"john",
// 		"peter",
// 		6: "luke",
// 		"a",
// 	}

// 	fmt.Println(names, len(names), cap(names))
// 	fmt.Println("the index at six is", names[6])
// 	fmt.Println(len(names))
// 	fmt.Println(names)
// 	names = append(names, "lea", "toy", "maembe")
// 	fmt.Println(names)

// 	for _, value := range names {
// 		switch valueLen := len(value); {
// 		case valueLen < 2:
// 			fmt.Println("what a short name!", valueLen)
// 		case valueLen < 3:
// 			fmt.Println("still a short name", valueLen)
// 		default:
// 			fmt.Println("your name qualifies", valueLen)
// 		}
// 	}

// 	switch n := rand.Intn(10); {
// 	case n == 0:
// 		fmt.Println("That's too low")
// 	case n > 5:
// 		fmt.Println("That's too big:", n)
// 	default:
// 		fmt.Println("That's a good number:", n)
// 	}

// }

// func main() {
// 	first()
// 	prob()
// }
