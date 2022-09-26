package ejerciciosTT
import "fmt"
func Prestamos()  {
	edad := 25
	isEmpledo := true
	antiguedad := 3
	sueldo:= 120000
	
	if edad > 22 && isEmpledo && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se le otorgara un credito tasa 0")
			
		}else {

			fmt.Println("Se le otorgara un credito con interese")
		}
	}else{
		fmt.Println("No cumple con los requisitos para obtener un prestamo")
	}
}