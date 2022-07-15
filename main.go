package main

import "fmt"

type person struct {
	firstName     string
	lastName      string
	age           int
	fanof         string
	username      string
	teamsFollowed []string
}

func main() {
	vikram := person{
		firstName:     "vikram",
		lastName:      "subramaniam",
		username:      "subvikr",
		age:           40,
		fanof:         "Arsenal FC",
		teamsFollowed: []string{"CSK", "Barca", "Juve"},
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
