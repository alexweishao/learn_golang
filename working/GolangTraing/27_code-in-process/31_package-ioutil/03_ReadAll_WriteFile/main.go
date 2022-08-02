package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	myStr := "Hello Everybody"

	err := ioutil.WriteFile("hey.txt", []byte(myStr), 0777) //创建文件hey.txt并写入字符串Hello Everybody
	if err != nil {
		log.Fatalln("something went wrong!", err.Error())
	}
	//Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据
	f, err := os.Open("hey.txt")
	if err != nil {
		log.Fatalln("couldn't read file!", err.Error())
	}
	defer f.Close()

	//ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Readall crashed!", err.Error())
	}

	fmt.Println(string(bs))

}
