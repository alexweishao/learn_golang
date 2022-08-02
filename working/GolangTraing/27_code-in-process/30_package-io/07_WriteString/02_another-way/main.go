package main

import (
	"io"
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
	//	bs := []byte(str)

	_, err = io.WriteString(f, str)
	//	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
