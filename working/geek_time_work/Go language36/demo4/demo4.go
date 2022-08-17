package main

import "flag"

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
} //StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。

func main() {
	flag.Parse() //从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	hello(name)  //调用同一个包下面的函数
}
