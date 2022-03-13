package Superciety

import (
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

func TraderBuyer(ElrondAmount *p.Decimal) {
	ScannedSuperPrices := GetAllSuperPrices()

	SuperBought, ResultingPrice := BuySuper(ElrondAmount, ScannedSuperPrices)
	fmt.Println("Using                   :", KO(ElrondAmount), "EGLD to buy Super,")
	fmt.Println("To buy $uper        buys:", KO(SuperBought), "$uper,")
	fmt.Println("")
	fmt.Println("OLD Prices:")
	PricePrinter(ScannedSuperPrices)
	fmt.Println("")
	fmt.Println("NEW Prices:")
	PricePrinter(ResultingPrice)
}
