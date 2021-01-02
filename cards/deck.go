package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of "deck"
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// ! Go doesn't allow you to declare variables and not use them. To stop receiving errors and allow the program to compile, we replace the variables with an underscore. This lets Go know that we don't care about them.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// # This function states that any variable inside of our application if type deck is assigned the print method

// * The "deck" argument is the type that we want to attach the method to
// * The "d" is the the actual copy of the deck we are working with
// * d is essentially a reference to the cards variable that we had in our func main if we replaced the letter d with the word cards, it would look very similar to that original loop. See starred section in main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// * Go has the ability to return multiple values from a function. In this function, we are returning two instances of "deck".
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// ----------------------------------------------------
// * To write files to a hard drive, go needs us to convert our data to a []byte
// * Here we convert the initail deck slice to a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// * bs argument is byte slice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
