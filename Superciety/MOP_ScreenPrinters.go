package Superciety

import (
	mt "SuperPlasm/SuperMath"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"strings"
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
	fmt.Println("==========================Scanned Input=================================================")
	fmt.Println("ADDRESS:", Addy)
	fmt.Println("Address         Super is:", KO(Super))
	fmt.Println("Address Super-Egld-LP is:", KO(SuperLP))
	fmt.Println("Address          MKSP is:", KO2(mt.MKSP2Print(SP)))
	fmt.Println("======================END-Scanned Input=================================================")
	fmt.Println("")
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
	fmt.Println("==========================SUPER--Prices=================================================")
	fmt.Println("===========================Super//EGLD===============================================")
	fmt.Println("Price of 1 EGLD       is:", KO(Prices.SP.SV.USDperEGLD), "USD")
	fmt.Println("Price of 1 EGLD       is:", KO(Prices.SP.SV.SUPERperEGLD), "SUPER")
	fmt.Println("Price of 1 SUPER      is:", KO(Prices.SP.SV.USDperSUPER), "USD")
	fmt.Println("============================SUPER--LP================================================")
	fmt.Println("One SUPER-EGLD-LP equals:", KO(Prices.SP.LPVC.SuperHalf), "Pool SUPER")
	fmt.Println("One SUPER-EGLD-LP equals:", KO(Prices.SP.LPVC.ElrondHalf), "Pool EGLD")
	fmt.Println("One SUPER-EGLD-LP equals:", KO(Prices.SP.LPVC.TotalUSD), "USD")
	fmt.Println("==========================SUPER--Prices=================================================")
	fmt.Println("")
}

//======================================================================================================================
//======================================================================================================================

func TwoMKSPrinter(MKSP1, MKSP2 *p.Decimal) {
	fmt.Println("Will increase your Meta-Kosonic Super-Power from:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(MKSP1)), "to:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(MKSP2)), "representing:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(mt.SUBxc(MKSP2, MKSP1))), "gain.")
}

//PriceDisplayOffset
func PDO(Amount *p.Decimal) string {
	var Result string
	Length := len(mt.KosonicDecimalConversion(p.NFI(10000000000))) - len(mt.KosonicDecimalConversion(Amount))
	Result = strings.Repeat(" ", Length-1)
	return Result
}

//PDO for strings
func PDO2(Amount string) string {
	var Result string
	Length := len(mt.KosonicDecimalConversion(p.NFI(10000000000))) - len(Amount)
	Result = strings.Repeat(" ", Length-1)
	return Result
}

//KosonicOffset
func KO(Number *p.Decimal) string {
	String1 := PDO(Number)
	String2 := mt.KosonicDecimalConversion(Number)
	Result := String1 + String2
	return Result
}

//KosonicOffset for strings
func KO2(Value string) string {
	String1 := PDO2(Value)
	Result := String1 + Value
	return Result
}
