package main
import (
	"fmt"
	"net/http"
	"time"
)


func main(){
	links:=[]string{
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}
	// chan ch bool
	c:=make(chan string)
	for _, link:= range links{
		go checkLink(link,c)
		
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	//Repitative check of links 
	for l:= range c{ 
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link,c)
		}(l)
		// go checkLink(l,c)
		
	}
	
}

func checkLink(link string,c chan string ){
	_,err:=http.Get(link)
	if err!=nil{
		fmt.Println(link,"might be down \n Error:",err)
		c <-  link
		return
	}
	c <-link
	fmt.Println(link, "is up!")
}