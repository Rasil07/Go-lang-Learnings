package main

import (
	"fmt"
	"net/http"
)

func main(){
	sites := []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
	}

	c:=make(chan string)

	for _,url := range sites{
		go checkLink(url,c)
	}

	for l:=range c {
		go func(link string){
			checkLink(link,c)
		}(l)
	}

}

func checkLink(url string,c chan string){

	_,err := http.Get(url)

	if err!=nil{
		fmt.Println(url + " might be down!")
		c <- url
		return
	}

	fmt.Println(url + " is up!")
	c <- url

}