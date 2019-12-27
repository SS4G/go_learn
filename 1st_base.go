package main

import (
	"fmt"
	"go_learn1/draw"
)

func stringTest(str0 string) {
	//var str0 = "abc"
	fmt.Printf("length of %s is %d\n", str0, len(str0))
}

//Main1st is test for basic of go
func Main1st() {
	fmt.Println("this is 1st main function")

	//here
	fmt.Println("hello world!")
	var a = 1
	var b = "a"
	stringTest(`你哈`)
	stringTest("你哈")
	fmt.Println(a, b)
	draw.Apple("red", 123)
}
