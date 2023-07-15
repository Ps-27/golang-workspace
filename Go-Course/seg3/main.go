//library package
// go get github.com/gorilla/mux     // third party library
package main

import (
	// "github.com/gorilla/mux" //third party library
	"fmt"  //standard library
)

func main(){
	yourNmae := "prity"

	fmt.Println(&yourNmae)

	fmt.Println(yourNmae)

	

	// r :=mux.NewRouter()
	// r.HandleFunc("/",Home)

	name("psinha")
	// name(0)

	combine("psinha",25)

	floatTest(1.2,1.5)
}
// func Home(){
	
// }

func name(name string)string{
	fmt.Println(name)
	return name
	// return 0

}

func combine(name string, age int){
	fmt.Println(name)
	fmt.Println(age)
	// fmt.Println(name + "is" + age + "years old")
}
func floatTest(f1,f2 float32){
	fmt.Println(f1+f2)
}