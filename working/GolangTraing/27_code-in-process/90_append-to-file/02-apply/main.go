package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk("27_code-in-process/90_append-to-file", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.Contains(path, ".go") { //判断字符串path是否包含子串.go。
			return nil
		}

		fmt.Println(path)
		//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。
		//它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
		//如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		text := `
/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/`

		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
		return nil
	})
}
