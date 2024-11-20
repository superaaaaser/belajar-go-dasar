package main

import "fmt"

var width int = 30
height := 30

func main(){
	// membuat variabel dengan keyword var dan mendeskripsikan tipe datanya
	var myName string = "Super Aser"

	// membuat variabel dengan keyword var tanpa mendeskripsikan tipe data
	var nickName = "superaser"

	// membuat variabel dengan tanda :=
	myAge := 34

	fmt.Println(myName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	fmt.Println(width)
	fmt.Println(height)
}