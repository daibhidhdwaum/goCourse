package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// * Since contact info has already been assigned a type, we could use an ES6 style way of defining it in our person struct and just call it contactInfo
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// ! This way of defining a new person struct assumes that you are entering in data in the order that it appears in the type definition.
	// alex := person{"Alex", "Anderson"}

	// * This way means we can change the order of fields, add more in and not affect our data
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// * This is a third way to define a struct
	// * When we define a struct this way, each field contained within the struct is assigned a falsey value.
	// ! Printing this out without any values using println will only print an empty object, it does contain the keys defined, though. We can see them by using printf instead
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// * The %+v represents the property + value
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Partey",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim)

	// ! This doesn't work, because Go is a pass by value language. It makes a copy of the original when we create the updateName function
	// jim.updateName("Jimmy")

	// * The apersand operator tells Go to give us acces to the memory address that this struct is sitting at. It is a reference to the memory address rather that the object
	// * it turn a value into a pointer
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// * This is a shortcut for the above
	jim.updateName("jimmy")

	jim.print()
}

// * Without a pointer this would be definded as (p person)
// *  The *pointer gives us the value that this memory address is point at
// * A star in front of a type is different from one in front of a value
// Q? Pointers are confusing. Study them more.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
