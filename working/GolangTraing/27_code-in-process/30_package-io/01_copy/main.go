package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rdr := strings.NewReader("test") //NewReader创建一个从s读取数据的Reader。返回的是一个指针类型*Reader
	fmt.Println(rdr)                 //&{test 0 -1}
	io.Copy(os.Stdout, rdr)          //将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误

}
