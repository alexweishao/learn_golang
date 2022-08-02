package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("working/GolangTraing/27_code-in-process/30_package-io/04_TeeReader/02/src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst1, err := os.Create("dst1.txt")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	dst2, err := os.Create("dst2.txt")
	if err != nil {
		panic(err)
	}
	defer dst2.Close()

	dst3, err := os.Create("dst3.txt")
	if err != nil {
		panic(err)
	}
	defer dst3.Close()

	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)
	rdr = io.TeeReader(rdr, dst2)

	io.Copy(dst3, rdr)

}
