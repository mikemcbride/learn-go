package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
// e.g., it has all the same behaviors as a slice of strings,
// but we can refer to it as "deck"
type deck []string

func newDeck() deck {
    cards := deck{}
    cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }
    return cards
}

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) toByteSlice() []byte {
    return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    // turn byte slice into a deck
    s := strings.Split(string(bs), ",")
    return deck(s)
}

func (d deck) shuffle () {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
    
    for i := range d {
        newPosition := r.Intn(len(d) - 1)
        d[i], d[newPosition] = d[newPosition], d[i]
    }
}
