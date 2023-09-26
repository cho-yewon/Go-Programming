package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// var a int // declaration
	// a = 7     // assign value

	// var a int = 7 //declaration & assign value

	// var a = 7 //declaration & assign value

	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	// b := 8.34 // float64
	// fmt.Println(b, reflect.TypeOf(b))

	var b float32 = 8.34
	fmt.Println(b, reflect.TypeOf(b))

	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)

	var c7 string // 변수명은 알파벳 대소문자로 시작해야 한다
	var d int     // 변수명은 숫자로 시작할 수 없다
	var e bool
	var f float64
	var G = 99 // 대문자로 시작하는 변수, 함수, 타입은 패키지 외부로 공개할 수 있다
	koreanZombie := "정찬성"
	fmt.Println(koreanZombie)

	fmt.Println(c7, d, e, f, G)

	fmt.Println('Z', '2', '\n', '김', '인', '하') // rune literals(int32) 90 50 10 44608 51064 54616
	fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf(2), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))
	// fmt.Println(math.Floor(삼.오), math.Ceil(이백십칠쩜오칠), math.Sqrt(sixteen))
	// fmt.Println(strings.Title(3.141592))
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	fmt.Println(strings.Title("open source\tprogramming\n\"go\"!"))
	//fmt.Println("OpenSource Programming")
}
