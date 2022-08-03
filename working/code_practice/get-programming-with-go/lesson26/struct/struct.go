package main

import "fmt"

func main() {
	superpowers := &[3]string{"test", "ok", "error"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1])
	fmt.Println(superpowers[:2])
}
