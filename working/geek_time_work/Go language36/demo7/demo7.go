package main

import (
	"flag"
	"fmt"
)

func main() {
	var name1 = flag.String("name", "everyone", "The greeting object.") //声明+赋值
	name2 := flag.String("name", "everyone", "The greeting object.")    //类型推断
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name1)
	fmt.Printf("Hello, %v!\n", name2)
}
