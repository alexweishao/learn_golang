package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("working/GolangTraing/27_code-in-process/30_package-io/05_ReadFull/src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst05.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 5) //创建一个容量为5的字节切片

	//ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。
	io.ReadFull(src, bs)

	/*Writer接口用于包装基本的写入方法。
	Write方法len(p) 字节数据从p写入底层的数据流。
	它会返回写入的字节数(0 <= n <= len(p))和遇到的任何导致写入提取结束的错误。
	Write必须返回非nil的错误，
	如果它返回的 n < len(p)。Write不能修改切片p中的数据，即使临时修改也不行。*/
	dst.Write(bs)
}
