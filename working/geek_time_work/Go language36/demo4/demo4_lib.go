package main //在同一个目录下的源码文件都需要被声明为属于同一个代码包

import "fmt"

func hello(name string) {
	fmt.Printf("Hello,%s\n", name)
}
