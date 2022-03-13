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
//[1]a		AddySpecsPrinter	Prints Super, SuperLP and MKSP Values, given an Elrond Address.
//[1]b          SpecsPrinterCore	Prints Super, SuperLP and MKSP Values.
//[2]           PricePrinter            Prints the information inside a MetaSuperPrice Structure.
//[3]           TwoMKSPrinter		Prints the information inside a MetaSuperPrice Structure.
//
//		[4]Offset String Functions
//[4]a          PDO			creates a specific length "empty" string, needed for aligned printing purposes.
//[4]b          PDO2			same as PDO, but uses a string an Input Amount.
//[4]c          KO			creates a string from a Decimal Number, needed for aligned printing purposes.
//[4]d          KO2			same as KO, but uses a string an Input Amount.
//
//======================================================================================================================
//======================================================================================================================
//
//
//[1]a          AddySpecsPrinter
//              Prints Super, SuperLP and MKSP Values, given an Elrond Address.
//
//
func AddySpecsPrinter(Addy ElrondAddress) {
	fmt.Println("")
	fmt.Println("ADDRESS:", Addy)

	SFT1Chain := CreateCamelChain()
	GetMeta := IzMeta(Addy, SFT1Chain)
	Super, LP := GetAddySuperValues(Addy)

	SpecsPrinterCore(Super, LP, GetMeta)
	return
}

//
//
//======================================================================================================================
//[1]b          SpecsPrinterCore
//              Prints Super, SuperLP and MKSP Values.
//
//
func SpecsPrinterCore(SuperAmount, LPAmount *p.Decimal, Meta bool) {
	MKSP := ConvertSupersToMKSP(SuperAmount, LPAmount, Meta)
	fmt.Println("==========================Scanned Input=================================================")
	fmt.Println("Amount          Super is:", KO(SuperAmount))
	fmt.Println("Amount  Super-Egld-LP is:", KO(LPAmount))
	fmt.Println("Amount           MKSP is:", KO2(mt.MKSP2Print(MKSP)))
	fmt.Println("======================END-Scanned Input=================================================")
	fmt.Println("")
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

//
//
//======================================================================================================================
//
//
//[3]           TwoMKSPrinter
//              Prints the information pertaining to a MKSP increase.
//		MKSP1 is the smaller value, MKSP2 is the greater value.
//
//
func TwoMKSPrinter(MKSP1, MKSP2 *p.Decimal) {
	fmt.Println("Will increase your Meta-Kosonic Super-Power from:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(MKSP1)), "to:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(MKSP2)), "representing:")
	fmt.Println("                        :", KO2(mt.MKSP2Print(mt.SUBxc(MKSP2, MKSP1))), "gain.")
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//		[4]Offset String Functions
//
//
//[4]a          PDO
//              PDO - Price Display Offset; creates a specific length "empty" string,
//             	needed for aligned printing purposes. Its length is dependent of the Input Amount
//
//
func PDO(Amount *p.Decimal) string {
	var Result string
	Length := len(mt.KosonicDecimalConversion(p.NFI(10000000000))) - len(mt.KosonicDecimalConversion(Amount))
	Result = strings.Repeat(" ", Length-1)
	return Result
}

//
//
//======================================================================================================================
//
//
//[4]b          PDO2
//              PDO - Price Display Offset 2; same as PDO, but uses a string an Input Amount.
//
//
func PDO2(Amount string) string {
	var Result string
	Length := len(mt.KosonicDecimalConversion(p.NFI(10000000000))) - len(Amount)
	Result = strings.Repeat(" ", Length-1)
	return Result
}

//
//
//======================================================================================================================
//
//
//[4]c          KO
//              KO - Kosonic Offset; creates a concatenated string from a Decimal Number
//		that has a proper offset (computed via PDO) needed for aligned printing purposes.
//
//
//
func KO(Number *p.Decimal) string {
	String1 := PDO(Number)
	String2 := mt.KosonicDecimalConversion(Number)
	Result := String1 + String2
	return Result
}

//
//
//======================================================================================================================
//
//
//[4]d          KO2
//              KO2 - Kosonic Offset 2; same as KO, but uses a string an Input Amount.
//
//
func KO2(Value string) string {
	String1 := PDO2(Value)
	Result := String1 + Value
	return Result
}

//======================================================================================================================
//======================================================================================================================
