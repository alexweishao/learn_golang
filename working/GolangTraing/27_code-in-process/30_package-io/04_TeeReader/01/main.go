package main

import (
	"io"
	"os"
)

func main() {
	//open()指定一个明确的路径
	src, err := os.Open("working/GolangTraing/27_code-in-process/30_package-io/04_TeeReader/01/src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	//creat()在根目录下创建指定文件名
	dst1, err := os.Create("working/GolangTraing/27_code-in-process/30_package-io/04_TeeReader/01/dst1.txt")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	dst2, err := os.Create("dst01.txt")
	if err != nil {
		panic(err)
	}
	defer dst2.Close()

	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)

	io.Copy(dst2, rdr)
}
