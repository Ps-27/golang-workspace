package main

import (
	"fmt"
	"runtime"
)

func main(){
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
		 
	}
}