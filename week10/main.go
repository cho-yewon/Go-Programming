package main

import (
	"fmt"
	"week10/src/greeting"
	"week10/src/mymath"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	fmt.Println(mymath.MyAbs(-99))
	fmt.Println(mymath.MyAbs(71))
}
