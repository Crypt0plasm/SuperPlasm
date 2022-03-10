package Superciety

import (
	"fmt"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/..MetaOperations../MOP_ScreenPrinters.go
//		Functions that are used for printing various data combination on screen.
//
//
//[1]		AddySpecsPrinter	Prints Elrond Address Super, SuperLP and MKSP Values.
//[2]           PricePrinter            Prints the information inside a MetaSuperPrice Structure.
//
//======================================================================================================================
//======================================================================================================================
//
//
//[1]           AddySpecsPrinter
//              Prints Elrond Address Super, SuperLP and MKSP Values.
//
//
func AddySpecsPrinter(Addy ElrondAddress) {
	Super, SuperLP, SP := AddySpecs(Addy)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("===========Scanned Input===============")
	fmt.Println("ADDRESS:", Addy)
	fmt.Println("My Super is", Super)
	fmt.Println("My SuperLP is", SuperLP)
	fmt.Println("My MKSP is", SP)
	return
}

//
//======================================================================================================================
//
//
//[2]           PricePrinter
//              Prints the information inside a MetaSuperPrice Structure.
//
//
func PricePrinter(Prices MetaSuperPrice) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("===========Super/EGLD==================")
	fmt.Println("Price of 1 EGLD is", Prices.SP.SV.USDperEGLD, "USD")
	fmt.Println("Price of 1 EGLD is", Prices.SP.SV.SUPERperEGLD, "SUPER")
	fmt.Println("Price of 1 SUPER is", Prices.SP.SV.USDperSUPER, "USD")

	fmt.Println("===========SUPER-LP====================")
	fmt.Println("1 SUPER LP equals", Prices.SP.LPVC.SuperHalf, "Super from the Super Pool")
	fmt.Println("1 SUPER LP equals", Prices.SP.LPVC.ElrondHalf, "EGLD from the Super Pool")
	fmt.Println("1 SUPER LP equals", Prices.SP.LPVC.TotalUSD, "USD")
	fmt.Println("")
}

//======================================================================================================================
//======================================================================================================================
