package main

import (
	"fmt"
	"hackaton/internal/file"
	"hackaton/internal/service"
)

func main() {

	var myFile file.File = file.File{Path: "tickets.csv"}

	data, _ := myFile.Read()

	for _, v := range data {
		fmt.Println(v.Date, v.Price, v.Id)
	}

	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
