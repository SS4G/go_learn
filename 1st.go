package main

import (
	"fmt"
)

func stringTest(str0 string) {
	//var str0 = "abc"
	fmt.Printf("length of %s is %d", str0, len(str0))
}

func main2() {
	fmt.Println("hello world!")
	//var a = 1
	//var b = "a"
	stringTest(`你哈`)
	//fmt.Println(a, b)
	//draw.DrawApple("red", 123)
}
