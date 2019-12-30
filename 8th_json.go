package main

import (
	"fmt"
	"encoding/json"
)

//Company opera
type Company struct {
	Name string
	IP []string
    Cash int
}

func main8th() {
	companys := []Company{Company{Name:"opera", IP:[]string{"192.168.0.1", "192.168.0.2"}, Cash:100},Company{Name:"baidu", IP:[]string{"192.168.0.3", "192.168.0.4"}, Cash:300}}
	//fmt.Println(companys)
	jsonString, _ := json.Marshal(companys)//json.MarshalIndent(companys, "", "    ")
	//生成一个byte数组
	fmt.Printf("%s\n", jsonString)
}
