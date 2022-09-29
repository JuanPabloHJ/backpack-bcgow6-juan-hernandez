package main

import (
	"fmt"
	// "os"
	"errors"
)

type Customer struct {
	Legajo,
	Name,
	Surname,
	PhoneNumber,
	ID,
	Address string
}

func (c *Customer) SetLegajo(legajo string) (err error) {

	if c.ID == "0" {
		err = errors.New("ID invalid, can't set legajo")

		panic(err)
	}

	c.Legajo = c.ID

	return
}

func main() {
	fmt.Print("")

}
