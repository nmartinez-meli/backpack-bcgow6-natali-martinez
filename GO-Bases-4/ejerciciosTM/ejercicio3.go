package ejerciciosTM

import (
	"fmt"
)

func SalaryTaxesErrorf(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}
