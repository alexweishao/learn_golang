package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func cp(srcName, dstName string) error {

	src, err := os.Open(srcName) //Open打开一个文件并从中读取一些数据,返回一个*File和error
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	/*Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，
	如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；
	对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。*/
	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v ", err)
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalln("Usage: 04_io-copy <SRC> <DST>")
	}

	srcName := os.Args[1]
	dstName := os.Args[2]

	err := cp(srcName, dstName)
	if err != nil {
		log.Fatalln(err)
	}
}
