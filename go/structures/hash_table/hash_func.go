package hash_table

import (
	"math"
)

func HashFuncDivision(value int, m int) int {
	return value % m
}

func HashFuncMultiplicacion(value int, m int) int {
	a := (math.Sqrt(5) - 1) / 2
	va := float64(value) * a
	v := va - math.Floor(va)               // Tomar la parte decimal
	return int(math.Floor(float64(m) * v)) // Multiplicar por m y tomar la parte entera
}
