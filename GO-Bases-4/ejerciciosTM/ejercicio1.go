package ejerciciosTM

import (
	"fmt"
)

type SalaryError struct {
	menssage string
}

func (err *SalaryError) Error() string {
	return fmt.Sprintf("error: %s", err.menssage)
}
func SalaryTaxesCustomError(salary int) error {
	if salary < 150000 {
		return &SalaryError{
			menssage: "el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}
