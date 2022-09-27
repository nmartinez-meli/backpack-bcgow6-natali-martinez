package ejerciciosTM

import (
	"errors"
	"fmt"
)

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func salario(minutos float64, categoria string) (float64, error) {
	horas := minutos / 60
	var s float64
	var err error
	switch categoria {
	case categoriaA:
		s = horas * 3000 * 1.5
	case categoriaB:
		s = horas * 1500 * 1.2
	case categoriaC:
		s = horas * 1000
	default:
		err = errors.New("La categoria no es correcta")
	}
	return s, err
}

func CalcularSueldo() {
	var minutosTrabajados float64 = 9600 // 8 hrs * 60 minutos * 5 dias a la semana * 4 semanas al mes
	var categoria string = "A"

	salario, err := salario(minutosTrabajados, categoria)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El Salario es: %f\n", salario)
	}
}
