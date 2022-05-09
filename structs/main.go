package main

import "fmt"

type contactInfo  struct{
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}
func main(){	
	var rasil person

	rasil.firstName = "Rasil"
	rasil.lastName="Baidar"
	rasil.contact.email="rasilgrt@gmail.com"
	rasil.contact.zipCode=14510	

	rasil.updateName("rasilo")
	rasil.getInfo()
}

func (p person) getInfo(){
	fmt.Printf("%+v\n",p)
}

func (p *person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}