package main

import "fmt"

var str string = "WELCOME...here!" // global variable
func main(){
	var str1 string = "prity" // full variable
	str2 := "ps27" //short variable
	fmt.Println("Hello Go!")
	const greet string ="Hello, welcome for  Go-lang training!"
	fmt.Println(greet)
	fmt.Println(str1)
	fmt.Println(str2)
	testing()

}

func testing(){
	fmt.Println("hello from testing func!")
	fmt.Println(str)
}