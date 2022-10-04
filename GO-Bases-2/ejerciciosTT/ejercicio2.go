package ejerciciosTT

import (
	"fmt"
	"os"
)

type Matrix struct {
	valores [][]float64
	alto    int
	ancho   int
}

func (matriz *Matrix) Set(valores []float64, alto, ancho int) error {
	if len(valores) != ancho*alto {
		return fmt.Errorf("no se pueden asignar %d valores a una matris de dimension %d X %d", len(valores), alto, ancho)
	}
	matriz.alto = alto
	matriz.ancho = ancho
	for fila := 0; fila < alto; fila++ {
		columna := []float64{}
		for col := 0; col < ancho; col++ {
			columna = append(columna, valores[fila*ancho+col])
		}
		matriz.valores = append(matriz.valores, columna)
	}
	return nil

}
func (matriz *Matrix) Cuadratica() bool {
	if (matriz.alto == matriz.ancho) && matriz.alto != 0 {
		return true
	}
	return false
}
func (matriz Matrix) Max() float64 {
	maximo := matriz.valores[0][0]
	for _, fila := range matriz.valores {
		for _, valor := range fila {
			if valor > maximo {
				maximo = valor
			}
		}
	}
	return maximo
}
func (m Matrix) Print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Printf("\t%v\n", m.valores[fila])
	}
}
func Matriz() {
	var m Matrix
	err := m.Set([]float64{1, 2, 3, 4, 54, 65, 76, 87, 87}, 3, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// m.Print()
	isCuadratica := m.Cuadratica()
	if isCuadratica {
		fmt.Println("La matriz dada es cuadratica")
	}
	fmt.Printf("%.2f es el valor maximo de la matriz:\n", m.Max())
	m.Print()

}
