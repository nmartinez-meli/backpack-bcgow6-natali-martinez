package ejerciciosTT

import "fmt"

func ObtenerMes1()  {
	numMes:=3
	switch (numMes){
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	}
	
}
func ObtenerMes2()  {
	numMes:=3
	mesMap := map[int]string{1:"Enero",2:"Febrero",3:"Marzo",4:"Abril",5:"Mayo",6:"Junio",7:"Julio",8:"agosto",9:"Septiembre",10:"Octubre",11:"Noviembre",12:"Diciembre"}
	fmt.Println(mesMap[numMes])
}
func ObtenerMes3()  {
	numMes:=3
	mesArr := [12]string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","agosto","Septiembre","Octubre","Noviembre","Diciembre"}
	fmt.Println(mesArr[numMes])
}