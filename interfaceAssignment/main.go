package main

import (
	"fmt"
	"io"
	"os"
)

/*

// Assignment 1 of interface

type shape interface{
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}

type square struct{
	sideLength float64
}

func main(){


	tr := triangle{
		base: 10,
		height: 15.4,
	}
	sq := square{
		sideLength: 15.4,
	}

	printArea(tr)
	printArea(sq)



}

func  printArea(sh shape){
	fmt.Println(sh.getArea())
}


func (tr triangle) getArea() float64{
	return tr.base*tr.height*0.5
}

func (sq square) getArea() float64{
	return sq.sideLength*sq.sideLength

}
*/

// Assignment 2 of Interface

type logWriter struct{}

func main(){
	fmt.Println(os.Args)

	file,err := os.Open(os.Args[1])
	if(err!=nil){
		fmt.Println("Error:",err)
		os.Exit(1)
		}
	lw := logWriter{}
	io.Copy(lw,file)
}

func (logWriter) Write(bs []byte)(int,error){
	fmt.Println(string(bs))
	
	return len(bs),nil
}