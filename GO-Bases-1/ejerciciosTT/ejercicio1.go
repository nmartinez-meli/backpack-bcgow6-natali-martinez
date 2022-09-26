package ejerciciosTT

import "fmt"

func CantLetras()  {
	palabra := "Onomatopeya"
	fmt.Printf("La longitud de la palabra %s es:%d\n",palabra,len(palabra))

	deletrear(palabra)
}

func deletrear(p string)  {
	fmt.Printf("La palabra %s se deletrea:",p)

	for i,_ := range p {
		if i==0 {
			fmt.Printf("%c",p[i])
		} else if i < len(p) - 1{
			fmt.Printf(" - %c",p[i])
		}else{
			fmt.Printf(" - %c\n",p[i])
		}
	}	
}