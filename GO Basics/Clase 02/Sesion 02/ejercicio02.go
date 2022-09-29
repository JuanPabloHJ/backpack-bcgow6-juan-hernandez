package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	width, height int
	data          []float64
	isSquare      bool
	maximum       float64
}

func (m *Matrix) PrintMatrix() (matrix string) {
	for i := 0; i < len(m.data); i++ {
		matrix += fmt.Sprintf("%.0f ", m.data[i])
		if (i+1)%m.width == 0 {
			matrix += "\n"
		}
	}
	return
}

func (m *Matrix) GetMaximum() float64 {
	m.maximum = math.Inf(-1)
	for _, value := range m.data {
		if value > m.maximum {
			m.maximum = value
		}
	}
	return m.maximum
}

func main() {
	m := &Matrix{
		width:  3,
		height: 3,
		data: []float64{
			1, 2, 3,
			4, 5, 6,
			7, 8, 9}}

	fmt.Print(m.PrintMatrix())
	fmt.Println(m.GetMaximum())
}
