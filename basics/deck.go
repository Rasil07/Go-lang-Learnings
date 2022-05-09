package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards:= deck{}
	cardSuits := []string{"Spades","Daimonds","Hearts","Clubs"}
	cardValues := []string{"Ace","2","3","4","5","6","7","8","9","10","J","Q","K"}

	for _,cardSuit:= range cardSuits{
		for _,cardValue := range cardValues{
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

func (d deck) print(){
	for _,card:=range d{
		fmt.Println((card))
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize] ,d[handSize:]
}

func (d deck) toString() string{ 
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {	
	return ioutil.WriteFile(fileName,[]byte(d.toString()), 0666 )
}

func newDeckFromFile(fileName string) deck{
	content,err:= ioutil.ReadFile(fileName)

	if err != nil {
	log.Fatal(err)
	os.Exit(1)
	}	

	return strings.Split(string(content),",")
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)

	for i:= range d{
		newPosition := r.Intn(len(d)-1)
		d[i],d[newPosition] = d[newPosition],d[i]
	}	
}