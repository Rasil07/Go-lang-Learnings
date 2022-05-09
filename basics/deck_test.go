package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d:=newDeck()

	if len(d) != 52{
		t.Errorf("Expected deck length of 52, but we got %d",len(d))
	}
	if(d[0]!="Ace of Spades"){
		t.Errorf("Expected first card to be Ace of Spades, but we got %v",d[0])
	}

	if(d[len(d)-1]!="K of Clubs"){
		t.Errorf("Expected last card to be K of Clubs, but we got %v",d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
os.Remove("_decktesting")

deck:=newDeck()
deck.saveToFile("_decktesting")

loadedDeck := newDeckFromFile("_decktesting")

if len(loadedDeck)!=52{
	t.Errorf("Expected 52 cards in a dec, got %v",len(loadedDeck))
}

os.Remove("_decktesting")


}