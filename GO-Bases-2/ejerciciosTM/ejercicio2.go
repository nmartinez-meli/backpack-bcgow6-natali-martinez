package ejerciciosTM

import (
	"errors"
	"fmt"
)

func promedio(notas ...float64) (float64, error) {
	var sumaNotas float64
	for _, nota := range notas {
		if nota > 0 {
			sumaNotas += nota
		} else {
			return 0.0, errors.New("La nota no puede ser negativa")
		}
	}
	return sumaNotas / float64(len(notas)), nil
}

func PromedioAlumno() {
	prom, err := promedio(2, 4, -6, 7, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio del alumno es: %f\n", prom)
	}
}
