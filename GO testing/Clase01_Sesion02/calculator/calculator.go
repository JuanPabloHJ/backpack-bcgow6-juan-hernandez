package calculator

import "errors"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func SortAscending(numbers []int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				num := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = num
			}
		}

	}

	return numbers
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Divided by zero")
	}

	return num1 / num2, nil
}
