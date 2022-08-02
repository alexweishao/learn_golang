package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("working/GolangTraing/27_code-in-process/30_package-io/06_LimitReader/src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst06.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	rdr := io.LimitReader(src, 5)
	io.Copy(dst, rdr)

}
