package main

import (
	"fmt"

	"trading/cli-calculator/buybacks"
	"trading/cli-calculator/data"
)

func main() {
	d, err := data.GettingData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recompras := buybacks.GenerarRecompras(d)
	buybacks.ImprimirRecompras(recompras, d.TipoDeCompra)
}
