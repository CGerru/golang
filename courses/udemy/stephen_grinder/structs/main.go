package main

import "fmt"

// Here is the definition of the struct. Could be a nice idea
// to make this in different file
type person struct {
	firstName string
	lastName  string
	// This struct has been added later on the lecture
	contact contactInfo
	// A struct property can be declared only by it's type
	// contactInfo <- Here the property has the same name than the struct
}
type contactInfo struct {
	zipCode int
	email   string
}

func main() {
	// This struct construction is not recommended, because relies
	// everithing in the order of struct properties
	// alex := person{"Alex", "Anderson"}

	// Another way to implement a struct. In this case Go assigns zero value
	// to the struct properties
	//var alex person

	// This is a more apropiate approach, a JSON like construction type
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		// Added later. Every final propertie must have a final comma
		contact: contactInfo{
			email:   "false@email.com",
			zipCode: 00000,
		},
	}

	fmt.Println(alex)
	// You can use printf to write avery property name and its value. This
	// might be usefull in the future, I don't know where how
	fmt.Printf("%+v", alex)
	fmt.Println()
	// This is the way to assign or modify the properties of a struct. Like
	// setters or JS properties
	alex.firstName = "Harry"
	alex.lastName = "de Tal"
	alex.updateName("Test") // This assignment has no effectD
	fmt.Println("--- Printing with no pointers")
	alex.print()

	// Let's play with pointers. Hell overloose

	/*
		Here the & operator gives the memory address of the value this
		variable is pointing at. TL; DR "Give me the address of this variable"
		So, this & gets the memory address and is assigned to a pointer (alexPointer)
		This pointer will be converted in the function below using the * operator inside
		the function
	*/
	alexPointer := &alex
	fmt.Println("This is a pointer", alexPointer)
	alexPointer.updateNamePointer("Jimmy")
	fmt.Println("--- Printing with pointers")
	alexPointer.print()
	/*
		Let's use a shortcut to pointers and stuff.
		The pointer function below can be called with a pointer (like alexPointer above) but
		with a person type too. This is a shortcut that can save some code writing, removing the
		need of using a pointer to call pointer functions
	*/
	alex.updateNamePointer("Now with more pointers!!!!")
}

/*
This function is supposed to update the person name, but ir does not
It happens because the Go behavior with parameters passed as reference
or value. I never figured it out in Java, I'm not going to with Go
Too old for this shit.
Update. Doing some experiments if you declare the receiver like this
p *person: the struct (or whatever) is passed as reference; but this way
p person: the struct (or something instead) is passed as value
*/
func (p person) updateName(firstName string) {
	p.firstName = firstName
}

/*
Implementation of the function above but this time we are messing things
up with pointers. Ok the * in front of a type means a very different thing
than * in front of a variable. Pay attention please!!!!

Whenever a * is seen where a type should by means that this is a description
of a type. So this function can only be called with a receiver to a pointer
p is a pointer then a value
*/
func (p *person) updateNamePointer(firstName string) {
	// This * takes this pointer and turns it into an actual value
	(*p).firstName = firstName
}

/*
Ok, this might help:
		ADDRESS -> VALUE using *address
		VALUE -> ADDRESS using &value
*/

// This function has a person struct as a receiver. This function can be
// called from any person struct
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
