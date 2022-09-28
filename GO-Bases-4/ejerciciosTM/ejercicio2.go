package ejerciciosTM

import (
	"errors"
)

func SalaryTaxesNewError(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}
