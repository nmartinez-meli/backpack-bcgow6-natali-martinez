package main

import (
	"fmt"
	"os"

	"github.com/nmartinez-meli/backpack-bcgow6-natali-martinez/ejerciciosTM"
)

func main() {
	salary := 160000
	fmt.Println("Ejercicio 1 TM")
	err := ejerciciosTM.SalaryTaxesCustomError(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ejercicio 1 TM paso sin error")
	fmt.Println("Ejercicio 2 TM")
	err = ejerciciosTM.SalaryTaxesNewError(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ejercicio 2 TM paso sin error")
	fmt.Println("Ejercicio 3 TM")
	err = ejerciciosTM.SalaryTaxesErrorf(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Ejercicio 3 TM paso sin error")
	fmt.Println("Debe pagar impuesto")
}
