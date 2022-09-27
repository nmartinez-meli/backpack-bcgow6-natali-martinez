package ejerciciosTM

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leerArchivo(nombre string) {
	data, err := os.ReadFile(nombre)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-15s %15s %10s\n", "ID", "Precio", "Cantidad")
	var total float64
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			campo := strings.Split(string(line), ";")
			precio, err := strconv.ParseFloat(campo[1], 64)
			cantidad, err := strconv.ParseFloat(campo[2], 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			total += precio * cantidad
			fmt.Printf("%-15s %15s %10s\n", campo[0], campo[1], campo[2])
		}
	}
	fmt.Printf("%-20s %.2f\n", "Total:", total)
}

func DetallarProductosComprados() {
	leerArchivo("productos.csv")
}
