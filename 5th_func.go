package main
import (
	"fmt"
	"strconv"
)

func makePrint(name string, value int) int {
	resultString := name + strconv.Itoa(value)
	fmt.Println(resultString)
    return len(resultString)
}

func makeManyPrint(name string, value int) (int, string) {
	resultString := name + strconv.Itoa(value)
	fmt.Println(resultString)
    return len(resultString), resultString//这里像元祖一样返回既可以
}

func createInner() []string {
	resultString := []string{"a", "b", "c"}
    return resultString//这里像元祖一样返回既可以
}

func add(a, b int) int {return a + b}
func mul(a, b int) int {return a * b}
func sub(a, b int) int {return a - b}

func calc(calcFunc func(int, int) (int), a, b int) (int) {
    return calcFunc(a, b)
}

func main5th() {
	fmt.Println("this is 5th main function")

	//函数创建
	//函数参数传入
	//函数返回值 
	bytelength := makePrint("songziheng", 26)
	fmt.Println(bytelength)
	
	//函数返回多个值
	bytelength, stringValue := makeManyPrint("songziheng", 26)
	fmt.Println(bytelength, stringValue)    

	//函数返回的对象作用域验证 (GC 验证)
	innerObject := createInner()//内部的变量可以被当做对象来返回
	fmt.Println(innerObject)
	innerObject[2] = "babs"
	fmt.Println(innerObject)

	//验证函数被当做对象传递的例子
	a := 10
	b := 20
	div := func(a, b int) (int) { return a / b}
	fmt.Println("add a=%d b=%d result=%d\n", a, b, calc(add, a, b))
	fmt.Println("mul a=%d b=%d result=%d\n", a, b, calc(mul, a, b))
	fmt.Println("sub a=%d b=%d result=%d\n", a, b, calc(sub, a, b))
	fmt.Println("div a=%d b=%d result=%d\n", a, b, calc(div, a, b))

}