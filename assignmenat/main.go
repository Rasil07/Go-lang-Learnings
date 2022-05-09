package main

import (
	"fmt"
)

func main(){

	intSlice := []int{}

	for i:=0; i<11; i++{
		intSlice = append(intSlice, i)
	}

	for _, val:=range intSlice{
		remainder:= val%2
	
		if remainder>0{
			fmt.Println(val,"is  odd")
		}
		if(remainder==0){
			fmt.Println(val," is even")
		}
	}


}