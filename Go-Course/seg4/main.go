package main

import (
	"fmt"
	"log"
)

func main(){
	test :="ps"

	// if test != "ps1"{
	// 	log.Fatalln("Its broken!")
	// } else if test != "ps1"{
	// 	log.Fatalln("its broken .....point2")
	// }
	//  Its broken!
    //  exit status 1

	// if test != "ps"{
	// 	log.Fatalln("Its broken!")
	// } else if test != "ps1"{
	// 	log.Fatalln("its broken .....point2")
	// } 
	//  its broken .....point2
    // exit status 1
    
	if test != "ps"{
		log.Fatalln("Its broken!")
	} else if test != "ps"{
		log.Fatalln("its broken .....point2")
	} 

	if test == "ps1"{
		fmt.Println("matched!")
	} else {
		fmt.Println("wrong path")
	}
}