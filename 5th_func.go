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
}