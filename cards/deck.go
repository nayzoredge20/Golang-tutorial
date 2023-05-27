package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// creating a new type of deck which is a slice of strings
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(FileName string) error {
	return ioutil.WriteFile(FileName, []byte(d.toString()), 0644)
	//io.Copy([]byte(FileName)  , []byte(d.toString()))
}

// iooutil.WriteFile(filename, []byte(string())) copies the contents of a string converted to a byte type into the filename andd also return a single error
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	var s deck // easy to understand as compared to s:=strings.Split(string(bs), ",") and then return deck(s)
	s = strings.Split(string(bs), ",")
	return s
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newp := r.Intn(len(d) - 1)
		d[index], d[newp] = d[newp], d[index]
		// by defalut slices are passed by reference so the value will be changed in the main deck
	}
}
