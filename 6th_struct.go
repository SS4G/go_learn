package main

import "fmt"

type School struct {
	scname string 
	location string 
	year int
}

type Student struct {
	name 	string
	id 	    int
    school  School
}

type StudentAnamous struct {
	name 	string
	id 	    int
    School
}


//Student
func PrintSchool(scPtr *School) {//指针的类型在左边 真是奇葩
	fmt.Printf("school_name=%s school_year=%d\n", scPtr.scname, scPtr.year)
}

func main6th() {
	fmt.Println("this is 6th main function")
	//结构体的定义
	xdu := School{scname: "XDU", location: "XIAN", year: 1931}
	fmt.Println(xdu)
	xduPtr := &xdu
	//结构体的指针操作 
	xduPtr.location = "Beijing"
	xdu.year = 1981
	fmt.Println(xdu)
	//结构体作为参数 指针式传入和 值传递
	PrintSchool(xduPtr)	
	//结构体的绑定成员方法？？
	var studentana StudentAnamous
	studentana.name = "zihengs"
	studentana.id = 140402
	studentana.scname = "XDU"
	studentana.location = "Xian"
	studentana.year = 1934
    fmt.Println(studentana)
}