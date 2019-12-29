package main
import "fmt"

func main4th() {
	fmt.Println("this is 4th main function")
	//字典的创建和初始化
	var map1 = map[string]int{"aa":1, "bb":2}
	fmt.Println(map1)
	map1["hk"] = 3
	fmt.Println(map1)
	fmt.Println(map1["hk"])
	jk, ok := map1["hk"] //如果有两个返回值值
	fmt.Println(jk, ok)
	//字典的迭代 按照 key value item
	for k, v := range map1 {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
 
	//字典的读写 if 返回值
	key := "sk"
	if jk, ok := map1[key]; ok {
		fmt.Println(jk)
	} else {
		fmt.Println("do not have this key", key)
	}
	//字典和json的互相转化

	key = "hk"
	if jk, ok := map1[key]; ok {
		fmt.Println(jk)
	} else {
		fmt.Println("do not have this key", key)
	}
}   
