package main

import "fmt"

func main(){
	s:=0
	for {
		s++
		if s == 2{
			fmt.Println("not done s = ", s)
			continue
			fmt.Println("definitely not done")
		}
		if s == 4 {
			fmt.Println("not done")
			break
		}
	}
	fmt.Println(s)

	
}