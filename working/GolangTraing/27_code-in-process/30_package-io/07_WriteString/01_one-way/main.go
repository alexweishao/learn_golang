package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "Put some phrase here."
	bs := []byte(str)
	fmt.Println(str) //Put some phrase here.

	res, err := f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
	fmt.Println(res)
}
