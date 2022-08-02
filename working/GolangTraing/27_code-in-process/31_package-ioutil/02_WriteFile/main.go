package main

import (
	"io/ioutil"
)

func main() {
	//WriteFile()函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
	err := ioutil.WriteFile("hello.txt", []byte("Hello world"), 0777)
	if err != nil {
		panic("something went wrong")
	}
}
