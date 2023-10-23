package main

import "fmt"

func swap(n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
func main() {
	a := 1
	b := 2
	c := &a // var c *int = &a
	//fmt.Printf("%d %d %d\n", a, b, c)
	fmt.Printf("%d %d %d\n", a, b, *c)
}
