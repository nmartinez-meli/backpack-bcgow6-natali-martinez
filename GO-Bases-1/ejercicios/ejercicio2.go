package ejercicios

import "fmt"

func Meteorologo()  {
	var temperatura, humedad,presion float64
	humedad = 60
	presion = 1200
	temperatura =19
	
	fmt.Printf("Humedad: %%%v\n",humedad)
	fmt.Printf("Presion: %v\n",presion)
	fmt.Printf("Temperatura: %v\n",temperatura)
}

