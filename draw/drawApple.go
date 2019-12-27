package draw

import (
	"fmt"
	"strconv"
)

var defalutColor = "black"
var defaultSize = 1

func init() {
	//is called first
	//fmt.Printf("default color = %s size=%d", defalutColor, defaultSize)
}

//Apple draw apple
func Apple(color string, size int) {
	fmt.Println("draw an apple with color=" + color + " size=" + strconv.Itoa(size) + "\n")
}
