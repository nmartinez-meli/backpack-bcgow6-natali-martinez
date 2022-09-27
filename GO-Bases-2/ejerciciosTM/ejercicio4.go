package ejerciciosTM

import (
	"errors"
	"fmt"
)

const (
	minimun = "minimun"
	maximun = "maximun"
	average = "average"
)

func calcularMin(valores ...float64) float64 {
	min := valores[0]
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func calcularMax(valores ...float64) float64 {
	max := valores[0]
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

func calcularAverage(valores ...float64) float64 {
	var totalNotas float64

	for _, valor := range valores {
		totalNotas += valor
	}
	return totalNotas / float64(len(valores))
}

func operation(function string) (func(...float64) float64, error) {
	switch function {
	case minimun:
		return calcularMin, nil
	case maximun:
		return calcularMax, nil
	case average:
		return calcularAverage, nil
	default:
		return nil, errors.New("No existe la funcion")
	}
}

func Estadisticas() {
	minFunction, err := operation(minimun)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageFunction, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxFunction, err := operation(maximun)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("La nota minima es: %f\n", minFunction(6, 8, 5, 9, 3, 2, 0))
	fmt.Printf("La nota maxima es: %f\n", maxFunction(6, 8, 5, 9, 3, 2, 0))
	fmt.Printf("La nota promedio es: %f\n", averageFunction(6, 8, 5, 9, 3, 2, 0))
}
