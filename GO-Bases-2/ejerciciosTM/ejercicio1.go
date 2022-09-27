package ejerciciosTM
import "fmt"


func impuesto(salario float64) float64 {
	var imp float64
	if salario > 50000.00 {
		imp += salario*0.17
	}
	if salario > 150000{
		imp += salario*0.10
	}
	return imp
}

func Sueldo()  {
	salarioEmp1:=60000.00
	salarioEmp2:=160000.00
	salarioEmp3:=45000.00

	fmt.Printf("El salario es de %f y su impuesto total es %f\n",salarioEmp1,impuesto(salarioEmp1))
	fmt.Printf("El salario es de %f y su impuesto total es %f\n",salarioEmp2,impuesto(salarioEmp2))
	fmt.Printf("El salario es de %f y su impuesto total es %f\n",salarioEmp3,impuesto(salarioEmp3))

}