package main

// ! Variables can be initialized outside of a function but cannot be assigned.

func main() {
	// * Go is statically typed, therefore we need to assign a type to our variable declaration

	// * This is the explicit way to declare a variable
	// # var card string = newCard()

	// * An  alternative way to define the same variable would be:
	// # card := "Ace of Spades"
	// * It should only be used for initalization. If the variable is being re-assigned it would look like this
	// # card = "Five of Diamonds"

	// ! Remember it is "string" and not "strings"
	// * To specify the elements that exist within the slice we add a set of curly braces after the type
	// cards := []string{"Ace of Diamonds", newCard()}

	// <=> When we created our deck.go package we initailized a new type called "deck" that is of the datatype string. This means that the code above can now be expressed as below

	// cards := deck{"Ace of Diamonds", newCard()}

	// <=> Once we build our newDeck function the cards variable then assigned the return value of a function call to newDeck

	cards := newDeck()

	// * To add a new card to our slice, we call the append method, pass in the slice we want to append to and the element we want to append
	// * It is then assigned to cards. It is important to be aware that the append function does not modify the original slice. It instead return a new slice that assigned back to the variable of "cards"
	// cards = append(cards, "Six of Spades")

	// * range cards = Take the slice of "cards" and loop over it. This is not dissimilar to a forEach is JS, but the syntax is rearranged slightly
	// # cards.range((i, card) => { })

	// <=> We declared a receiver in our deck.go package that loops over our deck and so the below code
	// * *****************************
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// * *******************************
	// <=> can be replace with

	// cards.print()

	// * Since we can have multiple return values from a function, we also need a way to assign them to variables.
	// * in this scenario, cards is assigned to hand and 5 is assigned to remainingCards
	hand, remainingCards := deal(cards, 3)

	hand.print()
	remainingCards.print()

}

// * Go needs us to be explicit about what kind of datatype is going to returned from any given function. We do this by adding the datatype after the parens
// <=> The newCard function is no longer needed once we build out the newDeck function
// func newCard() string {
// 	return "Five of Diamonds"
// }

// * Go has two list datatypes.They both must be defined with a type and all elements must be of the same type
// * An array - fixed length list of things
// * A slice - An array that can shrink and grow

// receiver == method()
