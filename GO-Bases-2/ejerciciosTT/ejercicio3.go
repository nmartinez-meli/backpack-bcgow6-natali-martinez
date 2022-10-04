package ejerciciosTT

import "fmt"

var (
	pequenio string = "PEQUEÑO"
	mediano  string = "MEDIANO"
	grande   string = "GRANDE"
)

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}
type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

type tienda struct {
	produtos []Producto
}

func nuevoProducto(tipoProducto, nombre string, precio float64) Producto {
	return &producto{precio: precio, nombre: nombre, tipo: tipoProducto}
}
func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (p producto) CalcularCosto() float64 {
	// Retorna precio del producto
	//más el costo de mantenerlo en tienda
	//más el costo de envío
	var mantencion float64
	var envio float64
	switch p.tipo {
	case mediano:
		mantencion = p.precio * 0.03
	case grande:
		mantencion = p.precio * 0.06
		envio = 2500
	}
	return p.precio + mantencion + envio
}

func (t tienda) Total() (total float64) {
	for _, producto := range t.produtos {
		total += producto.CalcularCosto()
	}
	return total
}
func (t *tienda) Agregar(p Producto) {
	t.produtos = append(t.produtos, p)
}

func EcommerceProductos() {
	insOficina := nuevaTienda()
	prod1 := nuevoProducto(pequenio, "lapicera", 125.00)
	insOficina.Agregar(prod1)
	prod2 := nuevoProducto(mediano, "mochila", 325.00)
	insOficina.Agregar(prod2)
	prod3 := nuevoProducto(grande, "escritorio", 4500.00)
	insOficina.Agregar(prod3)
	fmt.Printf("Precio total del producto: 💰 %.2f\n", insOficina.Total())
}
