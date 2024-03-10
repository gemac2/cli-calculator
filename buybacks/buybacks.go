package buybacks

import (
	"fmt"

	"trading/cli-calculator/data"
)

type Recompra struct {
	Precio             float64
	Cantidad           float64
	Margen             float64
	SumaMargen         float64
	Promedio           float64
	PrecioStopLoss     float64
	Valor              float64
	TotalMonedas       float64
	PrecioLiquidacion  float64
	SumaValor          float64
	UltimaDistancia    float64
	CoberturaDistancia float64
	TakeProfit         float64
}

func GenerarRecompras(d data.Data) []Recompra {
	recompras := make([]Recompra, 11)
	porcentajeRecompras := d.PorcentajeRecompras
	if d.TipoDeCompra == "long" {
		porcentajeRecompras = -d.PorcentajeRecompras
	}

	if d.PorcentajeMonedas == 100 {
		precioPrimeraRecompra := d.PrecioCompra * ((porcentajeRecompras / 100) + 1)
		recompras[0] = Recompra{Precio: d.PrecioCompra, Cantidad: d.CantidadMonedas}
		recompras[1] = Recompra{Precio: precioPrimeraRecompra, Cantidad: d.CantidadMonedas}

		for i := 2; i < 11; i++ {
			recompras[i].Precio = recompras[i-1].Precio * (1 + porcentajeRecompras/100)
			recompras[i].Cantidad = recompras[i-1].Cantidad + recompras[i-1].Cantidad*(d.PorcentajeMonedas/100)
		}

		poblarMargen(recompras, d.Apalancamiento)
		obtenerPrecioFinaldeCompras(recompras)
		recompras = calcularPrecioStopLoss(d, recompras)
		calcularDistanciaData(recompras)
		calcularTakeProfit(d, recompras)

		return recompras
	}

	recompras[0] = Recompra{Precio: d.PrecioCompra, Cantidad: d.CantidadMonedas}

	for i := 1; i < 11; i++ {
		recompras[i].Precio = recompras[i-1].Precio * (1 + porcentajeRecompras/100)
		recompras[i].Cantidad = recompras[i-1].Cantidad + recompras[i-1].Cantidad*(d.PorcentajeMonedas/100)
	}

	poblarMargen(recompras, d.Apalancamiento)
	obtenerPrecioFinaldeCompras(recompras)
	calcularPrecioStopLoss(d, recompras)

	return recompras
}

func ImprimirRecompras(recompras []Recompra, tipoDeCompra string) {
	for _, r := range recompras {
		if tipoDeCompra == "short" {
			if r.Precio < r.PrecioStopLoss {
				fmt.Printf("Precio: %.4f, Cantidad: %.4f, Valor:%.4f, TotalMonedas:%.4f, Promedio: %.4f, PrecioStopLoss: %.4f, UltimaDistancia:%.2f, CoberturaDistancia:%.2f, TakeProfit: %.2f\n", r.Precio, r.Cantidad, r.Valor, r.TotalMonedas, r.Promedio, r.PrecioStopLoss, r.UltimaDistancia, r.CoberturaDistancia, r.TakeProfit)
			}
		}

		if tipoDeCompra == "long" {
			if r.Precio > r.PrecioStopLoss {
				fmt.Printf("Precio: %.4f, Cantidad: %.4f, Valor:%.4f, TotalMonedas:%.4f, Promedio: %.4f, PrecioStopLoss: %.4f, UltimaDistancia:%.2f, CoberturaDistancia:%.2f,  TakeProfit: %.2f\n", r.Precio, r.Cantidad, r.Valor, r.TotalMonedas, r.Promedio, r.PrecioStopLoss, (r.UltimaDistancia * -1), (r.CoberturaDistancia * -1), r.TakeProfit)
			}
		}
	}
	fmt.Println()
}

func poblarMargen(recompras []Recompra, apalancamiento int) {
	for i := range recompras {
		recompras[i].Margen = (recompras[i].Precio * recompras[i].Cantidad) / float64(apalancamiento)
	}

	recompras[0].SumaMargen = recompras[0].Margen

	for i := 1; i < len(recompras); i++ {
		recompras[i].SumaMargen = recompras[i-1].SumaMargen + recompras[i].Margen
	}
}

func obtenerPrecioFinaldeCompras(recompras []Recompra) {
	recompras[0].Valor = recompras[0].Precio * recompras[0].Cantidad
	recompras[0].TotalMonedas = recompras[0].Cantidad
	for i := 1; i < len(recompras); i++ {
		recompras[i].Valor = recompras[i].Precio * recompras[i].Cantidad
		recompras[i].TotalMonedas = recompras[i].Cantidad + recompras[i-1].TotalMonedas
	}

	recompras[0].SumaValor = recompras[0].Valor
	for i := 1; i < len(recompras); i++ {
		recompras[i].SumaValor = recompras[i].Valor + recompras[i-1].SumaValor
	}

	recompras[0].Promedio = recompras[0].SumaValor / recompras[0].TotalMonedas

	for i := 1; i < len(recompras); i++ {
		recompras[i].Promedio = (recompras[i].Valor + recompras[i-1].SumaValor) / recompras[i].TotalMonedas
	}
}

func calcularPrecioStopLoss(d data.Data, recompras []Recompra) []Recompra {
	porcentajeLiquidacion := porcentajeLiquidacion(d.Apalancamiento)
	for i := range recompras {
		posicion := (recompras[i].Promedio * recompras[i].TotalMonedas)

		if d.TipoDeCompra == "short" {
			recompras[i].PrecioLiquidacion = recompras[i].Promedio * (1 + (float64(porcentajeLiquidacion) / 100))
			recompras[i].PrecioStopLoss = (posicion + d.StopLoss) / recompras[i].TotalMonedas
		}

		if d.TipoDeCompra == "long" {
			recompras[i].PrecioLiquidacion = (recompras[i].Promedio * ((float64(porcentajeLiquidacion) / 100) - 1)) * (-1)
			recompras[i].PrecioStopLoss = (posicion - d.StopLoss) / recompras[i].TotalMonedas
		}
	}
	return recompras
}

func porcentajeLiquidacion(apalancamiento int) (porcentajeLiquidacion float64) {
	switch apalancamiento {
	case 5:
		porcentajeLiquidacion = 20
	case 10:
		porcentajeLiquidacion = 10
	case 20:
		porcentajeLiquidacion = 5
	case 25:
		porcentajeLiquidacion = 4
	case 50:
		porcentajeLiquidacion = 2
	case 75:
		porcentajeLiquidacion = 1.33
	case 100:
		porcentajeLiquidacion = 1
	default:
		fmt.Println("\nOpción inválida")
	}

	return porcentajeLiquidacion
}

func calcularDistanciaData(recompras []Recompra) {

	for i := range recompras {
		recompras[i].UltimaDistancia = ((recompras[i].PrecioStopLoss / recompras[i].Precio) * 100) - 100
		recompras[i].CoberturaDistancia = ((recompras[i].PrecioStopLoss / recompras[i].Promedio) * 100) - 100
	}
}

func calcularTakeProfit(d data.Data, recompras []Recompra) {

	for i := range recompras {
		costoInicial := recompras[i].TotalMonedas * recompras[i].Promedio
		var precioTomarGanancias float64
		if d.TipoDeCompra == "long" {
			precioTomarGanancias = (((d.TakeProfitPorcentaje * recompras[i].Promedio) / 100) - recompras[i].Promedio) * -1
		}

		if d.TipoDeCompra == "short" {
			precioTomarGanancias = ((d.TakeProfitPorcentaje * recompras[i].Promedio) / 100) + recompras[i].Promedio
		}
		costoFinal := precioTomarGanancias * recompras[i].TotalMonedas

		recompras[i].TakeProfit = costoInicial - costoFinal

		if recompras[i].TakeProfit < 0 {
			recompras[i].TakeProfit = recompras[i].TakeProfit * -1
		}
	}
}
