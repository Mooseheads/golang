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
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pp *person) updateUsername(newUsername string) {
	*&pp.username = newUsername

}
