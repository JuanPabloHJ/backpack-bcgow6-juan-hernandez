package file

import (
	"fmt"
	"hackaton/internal/service"
	"io"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	var file []byte
	var id, price int

	file, err = os.ReadFile(f.Path)
	data := strings.Split(string(file), "\n")

	for _, ticket := range data {
		ticketData := strings.Split(ticket, ",")
		id, err = strconv.Atoi(ticketData[0])
		price, err = strconv.Atoi(ticketData[5])

		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       ticketData[1],
			Email:       ticketData[2],
			Destination: ticketData[3],
			Date:        ticketData[4],
			Price:       price,
		})
	}

	return
}

func (f *File) Write(service.Ticket) (err error) {
	io.Open(f.Path, fmt.Sprintf("%s,%s,%s,%s,%s,%s",service.Ticket.Id,service.Ticket.Names,service.Ticket.Email,service.Ticket.Destination,service.Ticket.Date,service.Ticket.Price)

	)
	return nil
}
