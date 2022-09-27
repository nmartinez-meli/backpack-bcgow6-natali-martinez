package ejerciciosTM

import (
	"fmt"
	"os"
)

func escribirArchivo(nombre string, dato string) {
	f, err := os.OpenFile(nombre, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = f.WriteString(dato + "\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GuardarProductosComprados() {
	productos := []string{"111223;30012.00;1", "444321;1000000.00;4", "434321;50.50;1"}
	for _, producto := range productos {
		escribirArchivo("productos.csv", producto)
	}
}
