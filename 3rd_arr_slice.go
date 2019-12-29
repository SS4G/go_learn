package main

import (
	"fmt"
	"sort"
	"strings"
	//"reflect"
)


type intSlice []int
//定义一个intSlice 类型 然后定义关联的方法(接口) 包括 Len Less Swap 三个方法 这样才可以后续被调用
func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//类比于python slice类似于list 数组是长度不可以变的list 但是数组可以直接==比较但是slice不行
//数组之间赋值是拷贝了一遍 但是slice只是传递了引用
//定义的区别是 数组创建时指定了长度 但是slice没有 
func main3rd() {
	fmt.Println("this is 3rd main function")
	//1 数组和slice创建
	fmt.Println("1--------------------------------------") //带有默认换行符
	var arr1 = [5]int{1, 2, 3}  //正常的定义长度初始化 没有初始化的位置被设置为对应类型的默认值
	var arr2 = [...]int{1, 2, 3, 4} // 根据初始化的情况来确定数组的长度 类似于c的 int arr[] = {1, 2, 3, 4}
	var arr3 = [5]int{1:1, 2:2} //初始化一部分元素
	var slice1 = []int{1:10, 3:40} //直接创建slice 没有指定尺寸就是slice
	var slice2 = arr1[2:]  //通过切片操作来创建
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3) 
	fmt.Println(slice1)
	//2 数组和slice内部的读写
	fmt.Println("2--------------------------------------") //带有默认换行符
	arr3[4] = 444
	fmt.Println(arr3) 
	//slice1[4] = 34 如果不扩容直接写会panic
	//3 数组的迭代
	fmt.Println("3--------------------------------------") //带有默认换行符
	//var i = 5
	for idx, val := range arr1 {
		fmt.Printf("idx=%d, val=%d\n", idx, val)
	}
	fmt.Println("=======================")
	for idx, val := range slice1 {
		fmt.Printf("idx=%d, val=%d\n", idx, val)
	}
	//4 数组的合并 只有第一个参数是slice才可以合并
	fmt.Println("4--------------------------------------") //带有默认换行符
    fmt.Println(append(slice1, 1232, 123)) //单独添加一个数
    fmt.Println(append(slice1, slice2...)) //加3个点是什么屁操作。。。 类似于python的可变参数 *？ //添加另一个slice
	//5 数组的翻转 只有第一个参数是slice才可以翻转
	//var slice3 = make([]int, 0)  //通过切片操作来创建
	//slice3 = append(slice3, 1, 2, 3, 4, 5)
	//fmt.Println(sort.Reverse(slice3)) //单独添加一个数

	//6 数组和slice的拷贝 
	fmt.Println("6--------------------------------------") //带有默认换行符
	arr1 = [5]int{1,2,3}
	arr3 = arr1
	arr3[2] = 1000
	fmt.Println("arr3", arr3)
	fmt.Println("arr1", arr1) //还是原来的数值 说明里面被copy了一遍
	slice3 := slice1
	slice3[2] = 1000
	fmt.Println("slice3", slice3)
	fmt.Println("slice1", slice1) //slice 1 和 3的数值是一样的 说明 slice1 和slice3是一个对象 
	//7 数组的排序 key 函数
	fmt.Println("7--------------------------------------") //带有默认换行符
	slice4ext := make(intSlice, 0)
	slice4ext = append(slice4ext, 3, 2, 1, 4, 21,6, 13)
	sort.Sort(slice4ext)
	fmt.Println("sorted slice", slice4ext)
	//8 数组字符串互转
	var stringArr = []string{"aa", "bb", "cc"}
	joinedString := strings.Join(stringArr, "#")
	fmt.Println("joined string", joinedString)
	//9 数组的比较
	var arrA = [3]string{"aa", "bb", "ccc"}
	var arrB = [3]string{"aa", "bb", "ccc"}
	var arrC = [5]string{"aa", "bb", "ccc", "dd", "eee"}
	if arrA == arrB {
		fmt.Println(arrA, "==", arrB)
	} else {
		fmt.Println(arrA, "!=", arrB)
	}
	fmt.Println("arr C", arrC)
	//if arrC == arrB { //different can't be 
	//	fmt.Println(arrC, "==", arrB)
	//} else {
	//	fmt.Println(arrC, "!=", arrB)
	//}
    
}
