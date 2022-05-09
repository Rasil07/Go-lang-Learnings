package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{

}

func main(){

	resp,err := http.Get("https://www.google.com/search?q=cats&oq=cats&aqs=chrome..69i57.937j0j1&sourceid=chrome&ie=UTF-8")
	

	if(err!=nil){
	fmt.Println("Error:",err)
	os.Exit(1)
	}	

	lw:=logWriter{}
	io.Copy(lw,resp.Body)


}

func (logWriter) Write(bs []byte)(int,error){
fmt.Println(string(bs))

return len(bs),nil
}