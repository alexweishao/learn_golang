// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// Uncomment below code to see the error
// (Just remove the // characters for all 3 lines below)

// This file cannot see main.go's imported names ("fmt").
// Because the imported names belong to file scope.
func byte() {
	fmt.Println("这是byte包下面的函数")
}

// func bye() {
// 	fmt.Println("Bye!")
// }
