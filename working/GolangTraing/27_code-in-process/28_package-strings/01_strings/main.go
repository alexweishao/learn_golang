package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"), //判断es是否在test中

		strings.Count("testtt", "t"), //输出一个整型，test里面包含几个s

		strings.HasPrefix("test", "te"), //判断test是否以te开头

		strings.HasSuffix("test", "st"), //判断test是否以st结尾

		strings.Index("test", "e"), //输出字符串test中第一次出现的e的下标

		strings.Join([]string{"a", "b"}, "-"), //将字符串用指定的字符来拼接

		strings.Repeat("a", 5), //重复指定字符count次，并输出

		strings.Replace("aaaa", "a", "b", 3), //将指定字符串的前n个字符用new代替

		strings.Split("a-b-c-d-e", "-"), //将字符串中的字符按照指定形式分割

		strings.ToLower("TEST"), //小写输出

		strings.ToUpper("test"), //大写输出

	)

	arr := []byte("test") //类型转换      将字符串转换成字节形式，并以ASCII码值输出
	fmt.Println(arr)

	str := string([]byte{'t', 'e', 's', 't'}) //[]byte{'t', 'e', 's', 't'}是个字节类型切片  转换成字符串string类型
	fmt.Println(str)
}

/*true
4
true
true
1
a-b
aaaaa
bbba
[a b c d e]
test
TEST
[116 101 115 116]
test
*/
