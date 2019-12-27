package main

import (
	"fmt"
	"strconv"
	"unsafe"
	"strings"
	"regexp"
	//"unicode/utf8"
)

func main2nd() {
	//1 打印字符串
	fmt.Println("1--------------------------------------") //带有默认换行符

	var str0 = "str0_string"
	var str1 = "str1_string"
	fmt.Print(str0) // 没有换行符
	fmt.Printf(" A got string %s %02d", str1, 1) //没有默认换行符 和c的一样
	fmt.Println(str0) //带有默认换行符
	fmt.Printf(" B got string %s", str1)
	
	fmt.Println("2--------------------------------------") //带有默认换行符

    fmt.Println("length of " + str0 + " is " + strconv.Itoa(len(str0))) //得使用srconv 来转换 挺不方便 所以还是用逗号分割方便一些
	//按照字符迭代字符串
	str0 = "abc中国" //字符串中存放的实际上是utf-8编码 但是迭代的时候 已经解码为定长的单个字符 rune
	fmt.Printf("origin string is %s has byte %d\n", str0, len(str0))
    for idx, c := range str0 {
		// 对字符串迭代 返回的是 idx 和 rune值 idx 从0开始 rune占有4个byte
		fmt.Println("got char ", idx, c, unsafe.Sizeof(c))
	}

	//for c := range str0 { //这样会导致只能得到 idx 而得不到rune
	//	fmt.Println("got char ", c, unsafe.Sizeof(c)) // 对字符串迭代 返回的是 idx 和 rune值 idx 从0开始 rune占有4个byte
	//}

	//3 字符串转化为rune数组
	fmt.Println("3--------------------------------------") //带有默认换行符
	str0 = "abc中国"
	fmt.Printf("correct slice by byte %s \n", str0[:3])
	fmt.Printf("wrone slice by byte %s \n", str0[:4])
	fmt.Printf("correct slice by byte %s \n", str0[:6])
	//rune_arr, size := utf8.DecodeRuneInString(str0)
	//fmt.Printf("%s has %d bytes and %d runes\n", str0, len(rune_arr))
	//var char_arr []rune
	//char_arr = str0.

	//4字符串切片
	fmt.Println("4--------------------------------------") //带有默认换行符
	str0 = "abc中国"
	runeArr := []rune(str0) //使用[]rune 类型转换时被转换为了 rune 数组
	fmt.Println("get rune array", runeArr) //可以看到转换后只有5个rune字符 如果要不想吧字符拆坏 就应该使用rune来处理
	//字符串的slice 实际上是对底层的utf-8 字符串按照字节序列进行slice 所以 如果是包含ascii之外的数据可能会导致 
	//一个完整的rune字符被拆分 显示的时候变成一个 � 

	//5  字符串和数字互转
	fmt.Println("5--------------------------------------") //带有默认换行符

	var intstr = "123"
	var nIntstr = "-123"
	var intval = 1239
	fmt.Printf("show string %s\n", strconv.Itoa(intval))
	intvalCasted, _ := strconv.Atoi(nIntstr)
	fmt.Printf("show int %d\n", intvalCasted)
	intvalCasted, _ = strconv.Atoi(intstr)
	fmt.Printf("show int %d\n", intvalCasted)

	//6  字符串utf-8 byte数组 互转
	fmt.Println("6--------------------------------------") //带有默认换行符
	str0 = "abc中国"
	var bytes = []byte(str0)
	fmt.Println("bytes of utf-8 'abc中国'", bytes) 
	//可以看到string转换为[]byte过程中 底层的数据并没有改变 []byte的长度还是9个字节 中文每个字符占据3个字节

	//7  原值字符串 转义 `
	fmt.Println("7--------------------------------------") //带有默认换行符
	str0 = "abc中国\n\n"
	//var bytes = []byte(str0)
	fmt.Println("non origin str of utf-8 ", str0) //可以看到转换后只有5个rune字符 如果要不想吧字符拆坏 就应该使用rune来处理
	strOrigin := `abc中国\n\n\x60` //有点类似于shell里面的单引号的作用 但是单引号本身只能通过 + "`"来表示 但是\不会被解释为转义字符
	fmt.Println("origin str of utf-8 \x60", strOrigin) //普通字符串中的反斜杠被解释为 转义字符
    fmt.Println([]byte("`"))
	//8  字符串的split join replace 
	fmt.Println("8--------------------------------------") //带有默认换行符
	str2 := "hello my name is JJ"
	str3 := "his,name,is,PP,!"
	fmt.Println("splited:", strings.Split(str2, "")) //sep 为空 就是按照空字符切分
	fmt.Println("splited:", strings.Fields(str2))    //按照空白字符切分
	fmt.Println("splited:", strings.Split(str3, ",")) //按照指定的分隔符切分 这里是"," 

	fmt.Println("replaced:", strings.Replace(str3, ",", "#", 2)) // 只替换前两个出现,的位置 
	fmt.Println("replaced:", strings.Replace(str3, ",", "#", -1)) //替换全部出现的位置 
    //go 没有自带的翻转函数

	//9  字符串格式化生成另一个字符串
	fmt.Println("9--------------------------------------") //带有默认换行符
	str1 = "Marry"
	newStr := fmt.Sprintf("%s just need %02d dollar", str1,  232) // 功能类似于python中的字符串format
	fmt.Println(newStr)
    
	//10 字符串正则表达式 工具使用
	fmt.Println("10--------------------------------------") //带有默认换行符
	regexpString := `h(\d+)`
	stringToBeMatch := "peach1231dak123das352ad"
	match, _ := regexp.MatchString(regexpString, stringToBeMatch) //返回的是一个bool 表示内部是否有匹配的部分
	if match {
		fmt.Printf("regexp %s matched %s\n", regexpString, stringToBeMatch)
	} else {
		fmt.Printf("regexp %s not match %s\n", regexpString, stringToBeMatch)
	}

	fmt.Println("regexp:", match)
	//或者使用提前编译好的regexp对象 匹配方式 
	regexpObject, _ := regexp.Compile(regexpString)
    fmt.Println("matched MatchString ", regexpObject.MatchString(stringToBeMatch))
	fmt.Println("matched FindString ", regexpObject.FindString(stringToBeMatch))
	//这里所谓的submatch 就是可以吧正则表达式group中的匹配也匹配出来
	//返回的是一个数组 想要获得哪一个group的值 直接去数组里面索引即可
	fmt.Println("matched FindStringSubmatch ", regexpObject.FindStringSubmatch(stringToBeMatch))
	fmt.Println("matched FindStringIndex ", regexpObject.FindStringIndex(stringToBeMatch))
	fmt.Println("matched FindStringSubmatchIndex ", regexpObject.FindStringSubmatchIndex(stringToBeMatch))
	fmt.Println("matched FindAllString ", regexpObject.FindAllString(stringToBeMatch, -1))
	fmt.Println("matched FindAllStringSubmatch ", regexpObject.FindAllStringSubmatch(stringToBeMatch, -1))
	//11 字符串子串 包含测试
	fmt.Println("11--------------------------------------") //带有默认换行符
	subStr := "123"
	fatherStr := "192381234832813123240"
	if strings.Contains(fatherStr, subStr) {
		fmt.Printf("%s contains %s \n", fatherStr, subStr)
	} else { //注意else的位置 真奇葩
		fmt.Printf("%s not contains %s \n", fatherStr, subStr)
	}

	//subStr := "123"
	//fatherStr := "192381234832813123240"
	if idx := strings.Index(fatherStr, subStr); idx != -1 { 
		//if 语句钱面可以有个 简短的赋值语句 这样就不用吧赋值语句写在if外面了 包含了
		fmt.Printf("%s contains %s @ index % d\n", fatherStr, subStr, idx) 
	} else { //注意else的位置 真奇葩
		fmt.Printf("%s not contains %s \n", fatherStr, subStr)
	}

}
