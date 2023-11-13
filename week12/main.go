package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	aS := a[:2]
	aS[1] = "Z"
	fmt.Println(a, aS)

	b := [4]int{4, 3, 2, 1}
	bS := b[1:3]
	fmt.Println(b, bS)
}
