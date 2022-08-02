package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	myPhrase := "i love you!"
	rdr := strings.NewReader(myPhrase) //NewReader入参是string类型。出参为*Reader

	bs, err := ioutil.ReadAll(rdr) //输入一个 io.Reader类型的r，输出[]byte类型和error
	if err != nil {
		log.Fatalln("my program broke again")
	}
	fmt.Println(bs)   //[105 32 108 111 118 101 32 121 111 117 33]
	str := string(bs) //将[]byte类型转化为string类型
	fmt.Println(str)
}
