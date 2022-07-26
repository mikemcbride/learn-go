package main

import "fmt"

type contactInfo struct {
    email string
    zipCode int
}

type person struct {
    firstName string
    lastName string
    contactInfo // this is like implicit property values in JS
}

func main () {
    jim := person {
        firstName: "Jim",
        lastName: "Party",
        contactInfo: contactInfo {
            email: "jim@gmail.com",
            zipCode: 90210,
        },
    }
    // &jim says "give me the memory address of jim", rather than the value of jim
    jimPointer := &jim
    jimPointer.updateName("Jimmy")
    // this is quite verbose, so Go has a shortcut we can do:
    // instead of initializing a new pointer and calling the function on the pointer,
    // Go allows us to pass the variable of that type and it will automatically turn it into a pointer for us.
    // therefore we can just do this:
    // jim.updateName("Jimmy") as long as updateName has a pointer type in the receiver.
    jim.print()
}

func (p person) print() {
    fmt.Printf("%+v", p)
}

// this function takes a pointer - a memory address, rather than a value - as the receiver
// a * in front of a type means something totally different than a * in front of a variable.
// * in front of a type means it is a type description - specifies that we want a pointer to an item of that type.
// * in front of a variable says give me the value of this pointer.
func (pointerToPerson *person) updateName(newFirstName string) {
    // pointerToPerson is the memory address, so *pointerToPerson says "give me the value at this memory address"
    (*pointerToPerson).firstName = newFirstName
}

// quick notes on pointers:
// if you have a memory address, you can turn it into a value with *address
// if you have a value, you can turn it into a memory address with &value

// pass by value vs pass by reference
// just like in JavaScript, Go has some types that always pass by value and some that pass by reference.
// when we need to modify a type that passes by value, we need to use a pointer.
// 
// Reference Types | Value Types
// ------------------------------
// slices          | int
// maps            | float
// channels        | string
// pointers        | bool
// functions       | structs
