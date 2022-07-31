package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	rdr := strings.NewReader("hello world")

	io.Copy(dst, rdr)
}
