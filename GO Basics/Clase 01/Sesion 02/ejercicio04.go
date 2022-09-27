package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	mayor := 0

	for key, element := range employees {
		if element > 21 {
			mayor++

			fmt.Println(key)
		}

		fmt.Println("Mayores a 21", mayor)

		employees["Federico"] = 25

		delete(employees, "pedro")
	}
}
