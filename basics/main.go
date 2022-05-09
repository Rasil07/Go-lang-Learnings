package main

func main(){
	newD :=newDeckFromFile("My_Cards")
	newD.shuffle()
	newD.print()	
}

