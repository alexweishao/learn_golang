package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	fmt.Println(f)
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke again")
	}

	str := string(bs)
	fmt.Println(str)
}
