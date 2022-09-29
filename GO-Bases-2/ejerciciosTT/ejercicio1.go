package ejerciciosTT

import "fmt"

type Alumno struct {
	Nombre   string `json:"name,omitempty"`
	Apellido string `json:"lastName,omitempty"`
	DNI      string `json:"dni,omitempty"`
	Fecha    string `json:"dateIN,omitempty"`
}

func (student Alumno) detalle() {
	if student.Nombre !="" {
		
		fmt.Println("El Nombre del alumno es:", student.Nombre)
	}
	if student.Apellido != "" {
		
		fmt.Println("El Apellido del alumno es:", student.Apellido)
	}
	if student.DNI != "" {
		
		fmt.Println("El DNI del alumno es:", student.DNI)
	}
	if student.Fecha != "" {
		
		fmt.Println("La feccha de ingreso del alumno es:", student.Fecha)
	}
}

func RegistroStudiantes() {
	student1 := Alumno{Nombre: "Benjamin", Apellido: "Pregno", DNI: "35.967.234"}
	student1.detalle()
	student2 := Alumno{Nombre: "Natali", Apellido: "Martinez", DNI: "34.707.384"}
	student2.detalle()
}
