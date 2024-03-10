package data

import "fmt"

type Data struct {
	TipoDeCompra         string
	PrecioCompra         float64
	PorcentajeRecompras  float64
	CantidadMonedas      float64
	PorcentajeMonedas    float64
	StopLoss             float64
	Apalancamiento       int
	TakeProfitPorcentaje float64
}

func GettingData() (data Data, err error) {
	var tipoDeCompra string
	var precioCompra float64
	var porcentajeRecompras float64
	var cantidadMonedas float64
	var porcentajeMonedas float64
	var stopLoss float64
	var apalancamiento int
	var takeProfit float64

	fmt.Print("Ingrese tipo de compra: ")
	_, err = fmt.Scanln(&tipoDeCompra)
	if err != nil {
		fmt.Println("Error al ingresar el tipo de compra. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese el precio inicial de compra: ")
	_, err = fmt.Scanln(&precioCompra)
	if err != nil {
		fmt.Println("Error al ingresar el precio inicial de compra. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese el porcentaje de recompra: ")
	_, err = fmt.Scanln(&porcentajeRecompras)
	if err != nil {
		fmt.Println("Error al ingresar el porcentaje de recompra. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese la cantidad inicial de monedas: ")
	_, err = fmt.Scanln(&cantidadMonedas)
	if err != nil {
		fmt.Println("Error al ingresar la cantidad inicial de monedas. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese el porcentaje de monedas para recompras: ")
	_, err = fmt.Scanln(&porcentajeMonedas)
	if err != nil {
		fmt.Println("Error al ingresar el porcentaje de monedas para recompra. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese la cantidad que está dispuesto a perder (Stop Loss): ")
	_, err = fmt.Scanln(&stopLoss)
	if err != nil {
		fmt.Println("Error al ingresar el stoploss. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese el porcentaje donde cobrara ganancias (Take Profit): ")
	_, err = fmt.Scanln(&takeProfit)
	if err != nil {
		fmt.Println("Error al ingresar el takeProfit. Por favor, inténtelo nuevamente.")
		return data, err
	}

	fmt.Print("Ingrese a cuantas x esta apalancado: ")
	_, err = fmt.Scanln(&apalancamiento)
	if err != nil {
		fmt.Println("Error al ingresar el apalancamiento. Por favor, inténtelo nuevamente.")
		return data, err
	}

	data = NewData(tipoDeCompra, precioCompra, porcentajeRecompras, cantidadMonedas, porcentajeMonedas, stopLoss, takeProfit, apalancamiento)

	return data, nil
}

func NewData(tipoDeCompra string, precioCompra, porcentajeRecompras, cantidadMonedas, porcentajeMonedas, stopLoss, takeProfit float64, apalancamiento int) Data {
	return Data{
		TipoDeCompra:         tipoDeCompra,
		PrecioCompra:         precioCompra,
		PorcentajeRecompras:  porcentajeRecompras,
		CantidadMonedas:      cantidadMonedas,
		PorcentajeMonedas:    porcentajeMonedas,
		StopLoss:             stopLoss,
		TakeProfitPorcentaje: takeProfit,
		Apalancamiento:       apalancamiento,
	}
}
