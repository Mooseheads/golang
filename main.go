package main

import "fmt"

type address struct {
	streetName string
	houseNo    int
	pincode    int
}

type person struct {
	firstName     string
	lastName      string
	age           int
	fanof         string
	username      string
	teamsFollowed []string
	address
}

type memaddress struct {
	value int
}

func main() {
	vikram := person{
		firstName:     "vikram",
		lastName:      "subramaniam",
		username:      "subvikr",
		age:           40,
		fanof:         "Arsenal FC",
		teamsFollowed: []string{"CSK", "Barca", "Juve"},
		address: address{
			streetName: "clyde street",
			houseNo:    2247,
			pincode:    450506,
		},
	}

	vikram.updateUsername("vikramvns")
	vikram.print()

	val := memaddress{
		value: 10,
	}

	// Print the memory address / RAM address of a struct
	fmt.Printf("%p\n", &val)

	// Print the value of a struct with the explanation
	fmt.Printf("%+v", val)

	// Print the value of a struct
	fmt.Printf("%v", val)

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pp *person) updateUsername(newUsername string) {
	*&pp.username = newUsername

}
