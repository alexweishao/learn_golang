package main

import (
	"fmt"
)

//先执行 import，然后执行const，再执行var，再执行init()，最后执行main()
func init() {
	fmt.Println("Who ran first?", x)
}

func main() {
	fmt.Println("Hello world.")
}

var x int = 17
