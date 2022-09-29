package ejerciciosTM

import (
	"fmt"
	"os"
)

type clientes struct {
	Legajo         int
	Nombre         string
	Apellido       string
	DNI            int
	NumeroTelefono string
	Domicilio      string
}

func datosClientes(archivo string) []byte {
	datos, err := os.ReadFile(archivo)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return datos
}

func LeerClientes() {
	defer fmt.Println("Ejecucion finalizada")
	clietes := datosClientes("./Go-Bases-5/ejerciciosTM/clientses.txt")
	println(string(clietes))
}
