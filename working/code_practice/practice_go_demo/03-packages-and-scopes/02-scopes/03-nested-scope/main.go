// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// I didn't talk about this in a lecture
// As a side-note, I wanted to put it here
// Please review it.

var declareMeAgain = 10 //全局变量

func nested() { // block scope starts

	// declares the same variable
	// they both can exist together
	// this one only belongs to this scope
	// package's variable is still intact
	var declareMeAgain = 5 //局部变量
	fmt.Println("inside nested:", declareMeAgain)

} // block scope ends

func main() { // block scope starts

	fmt.Println("inside main:", declareMeAgain) //10

	nested() //5

	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	fmt.Println("inside main:", declareMeAgain) //10

} // block scope ends
