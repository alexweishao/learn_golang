package main

import "fmt"

func main() {
	a := 42
	address := &a
	fmt.Printf("address is %T", address) //打印出 address is *int
}
