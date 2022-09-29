package main

import (
	"fmt"
)

type Alumn struct {
	Name      string
	Surname   string
	ID        int
	BirthDate string
}

func (a *Alumn) Detail() string {
	return fmt.Sprintf("Name: %s, Surname: %s, ID: %d, Date: %s", a.Name, a.Surname, a.ID, a.BirthDate)
}

func main() {
	alumn := Alumn{Name: "Pepe", Surname: "Perez", ID: 10235426874, BirthDate: "2015-01-01"}
	fmt.Println(alumn.Detail())
}
