package ejerciciosTT

import "fmt"

func EdadEmpleado()  {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es: %d\n",employees["Benjamin"])
	var mayores int
	for _,edad := range employees {
		if edad > 21 {
			mayores++
		}
	}
	fmt.Printf("La cantidad de empleados mayores a 21 es: %d\n",mayores)
	employees["Federico"]=25
	delete(employees,"Pedro")
	fmt.Printf("Empleados: %#v\n",employees)
}

