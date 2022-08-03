package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//Walk函数对每一个文件/目录都会调用WalkFunc函数类型值。
	//调用时path参数会包含Walk的root参数作为前缀；就是说，如果Walk函数的root为"dir"，
	//该目录下有文件"a"，将会使用"dir/a"调用walkFn参数。
	//walkFn参数被调用时的info参数是path指定的地址（文件/目录）的文件信息，类型为os.FileInfo。
	filepath.Walk("27_code-in-process/90_append-to-file", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.Contains(path, ".go") { //判断path目录中是否具有包含.go的文件
			return nil
		}

		fmt.Println(path) //打印path路径
		return nil

	})
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
