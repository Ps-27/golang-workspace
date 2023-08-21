// package main
 
// import (
//  "fmt"
// )
 
// func main() {
//  greeting := "Hi There!"
 
//  go (func() {
//      fmt.Println(greeting) 
//  })()
// }
// func main() {
// 	c := make(chan string)
// 	c <- "Hi there!"
// }

package main
 
import "fmt"
// var  deckSize=20
deckSize := 20

func main() {
 fmt.Println(deckSize)
 c := make(chan string)
 for i := 0; i < 4; i++ {
     go printString("Hello there!", c)
 }
 
 for i:=0;i<4;i++{
	s:=<-c
     fmt.Println(s)
 }
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
	// fmt.Println(<- c)
}
 
func printString(s string, c chan string) {
 fmt.Println(s)
 c <- "Done printing." 
}