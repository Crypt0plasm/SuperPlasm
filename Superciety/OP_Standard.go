package Superciety

import (
	mt "SuperPlasm/SuperMath"
	p "github.com/Crypt0plasm/Firefly-APD"
	"io/ioutil"
	"log"
	"net/http"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/Operations/OP_Standard.go
//		Basic Functions, Blockchain related Functions and math.
//
//
//              [A]Basic Functions
//[A]01         PercentSwing            Computes the % difference between 2 values.
//[A]02         OnPage                  Basic Snapshot Function
//
//              [B]AtomicUnit String Converter Functions
//[B]01         ConvertAU18             Converts a string of numbers as AU to decimals, considering 18 decimals.
//[B]02         ConvertAU06             Converts a string of numbers as AU to decimals, considering  6 decimals.
//
//              [C]Blockchain Operation Functions
//[C]01         BuySuper                Computes the Super Amount that can be bought with a given amount of EGLD.
//[C]02         SellSuper               Computes the EGLD Amount that can be gained by selling a given amount of Super.
//[C]03         AddLiquidity            Computes the ESDT Amount req. to add Liq. (given EGLD and Input Prices).
//
//======================================================================================================================
//======================================================================================================================
//
//
//[A]           Basic Functions
//
//
//[A]01         PercentSwing
//              Computes the % difference between Value2 and Value1.
//
//
func PercentSwing(Value1, Value2 *p.Decimal) *p.Decimal {
	M1 := mt.MULxc(Value2, p.NFS("100"))
	D1 := mt.DIVxc(M1, Value1)
	S1 := mt.SUBxc(D1, p.NFS("100"))
	PP := mt.TruncateCustom(S1, 6)
	return PP
}

//======================================================================================================================
//
//
//[A]02         OnPage
//              Basic Snapshot Function
//              Snapshots Link and returns string
//
//
func OnPage(Link string) string {
	res, err := http.Get(Link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

//
//
//======================================================================================================================
//
//
//[B]           AtomicUnit String Converter Functions
//
//
//[B]01         ConvertAU18 Converts a string of raw numbers as atomic units, to its respective Decimal
//	        Usage for 18 Decimals
//
//
func ConvertAU18(Number string) *p.Decimal {
	Value := p.NFS(Number)
	Result := mt.DIVxc(Value, mt.POWxc(p.NFI(10), p.NFI(18)))
	return Result
}

//======================================================================================================================
//
//
//[B]02         ConvertAU06 Converts a string of raw numbers as atomic units, to its respective Decimal
//	        Usage for 6 Decimals
//
//
func ConvertAU06(Number string) *p.Decimal {
	Value := p.NFS(Number)
	Result := mt.DIVxc(Value, mt.POWxc(p.NFI(10), p.NFI(6)))
	return Result
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//[C]           Blockchain Operation Functions
//
//
//[C]01         BuySuper computes the Super Amount that can be bought with a given amount of EGLD.
//              And computes the resulting Prices given the input Prices
//
//
func BuySuper(ElrondAmount *p.Decimal, InputPrice MetaSuperPrice) (SuperBought *p.Decimal, OutputPrice MetaSuperPrice) {
	NewPoolEgld := mt.TruncateCustom(mt.ADDxc(InputPrice.SP.SuperPool.EgldAmount, ElrondAmount), 18)
	NewPoolSuper := mt.TruncateCustom(mt.DIVxc(mt.MULxc(InputPrice.SP.SuperPool.EsdtAmount, InputPrice.SP.SuperPool.EgldAmount), NewPoolEgld), 18)
	SuperBought = mt.TruncateCustom(mt.SUBxc(InputPrice.SP.SuperPool.EsdtAmount, NewPoolSuper), 18)
	//Applying Maiar Fee on SuperBought
	SuperBought = mt.TruncateCustom(mt.MULxc(SuperBought, p.NFS("0.997")), 18)

	OutputPrice.SP.DollarPool = InputPrice.SP.DollarPool
	OutputPrice.SP.SuperPool = ESDTPoolAmounts{EsdtAmount: NewPoolSuper, EgldAmount: NewPoolEgld}

	OutputUSDperEGLD := InputPrice.SP.SV.USDperEGLD
	OutputSuperperEGLD := mt.TruncateCustom(mt.DIVxc(NewPoolSuper, NewPoolEgld), 18)
	OutputUSDperSUPER := mt.TruncateCustom(mt.DIVxc(OutputUSDperEGLD, OutputSuperperEGLD), 18)
	OutputPrice.SP.SV = SuperValue{USDperEGLD: OutputUSDperEGLD, SUPERperEGLD: OutputSuperperEGLD, USDperSUPER: OutputUSDperSUPER}

	OutputPrice.Liquidity = InputPrice.Liquidity
	OutputSuperLP1 := mt.TruncateCustom(mt.DIVxc(NewPoolSuper, OutputPrice.Liquidity), 18)
	OutputSuperLP2 := mt.TruncateCustom(mt.DIVxc(NewPoolEgld, OutputPrice.Liquidity), 18)
	OutputSuperLPValue := mt.TruncateCustom(mt.ADDxc(mt.MULxc(OutputSuperLP1, OutputUSDperSUPER), mt.MULxc(OutputSuperLP2, OutputUSDperEGLD)), 18)
	OutputPrice.SP.LPVC = SuperLPValueComposition{SuperHalf: OutputSuperLP1, ElrondHalf: OutputSuperLP2, TotalUSD: OutputSuperLPValue}

	return SuperBought, OutputPrice

}

//======================================================================================================================
//
//
//[C]02         SellSuper computes the EGLD Amount that can be gained by selling a given amount of Super.
//              And computes the resulting Prices given the input Prices
//
//
func SellSuper(SuperAmount *p.Decimal, InputPrice MetaSuperPrice) (EgldBought *p.Decimal, OutputPrice MetaSuperPrice) {
	NewPoolSuper := mt.TruncateCustom(mt.ADDxc(InputPrice.SP.SuperPool.EsdtAmount, SuperAmount), 18)
	NewPoolEgld := mt.TruncateCustom(mt.DIVxc(mt.MULxc(InputPrice.SP.SuperPool.EsdtAmount, InputPrice.SP.SuperPool.EgldAmount), NewPoolSuper), 18)
	EgldBought = mt.TruncateCustom(mt.SUBxc(InputPrice.SP.SuperPool.EgldAmount, NewPoolEgld), 18)
	//Applying Maiar Fee on SuperBought
	EgldBought = mt.TruncateCustom(mt.MULxc(EgldBought, p.NFS("0.997")), 18)

	OutputPrice.SP.DollarPool = InputPrice.SP.DollarPool
	OutputPrice.SP.SuperPool = ESDTPoolAmounts{EsdtAmount: NewPoolSuper, EgldAmount: NewPoolEgld}

	OutputUSDperEGLD := InputPrice.SP.SV.USDperEGLD
	OutputSuperperEGLD := mt.TruncateCustom(mt.DIVxc(NewPoolSuper, NewPoolEgld), 18)
	OutputUSDperSUPER := mt.TruncateCustom(mt.DIVxc(OutputUSDperEGLD, OutputSuperperEGLD), 18)
	OutputPrice.SP.SV = SuperValue{USDperEGLD: OutputUSDperEGLD, SUPERperEGLD: OutputSuperperEGLD, USDperSUPER: OutputUSDperSUPER}

	OutputPrice.Liquidity = InputPrice.Liquidity
	OutputSuperLP1 := mt.TruncateCustom(mt.DIVxc(NewPoolSuper, OutputPrice.Liquidity), 18)
	OutputSuperLP2 := mt.TruncateCustom(mt.DIVxc(NewPoolEgld, OutputPrice.Liquidity), 18)
	OutputSuperLPValue := mt.TruncateCustom(mt.ADDxc(mt.MULxc(OutputSuperLP1, OutputUSDperSUPER), mt.MULxc(OutputSuperLP2, OutputUSDperEGLD)), 18)
	OutputPrice.SP.LPVC = SuperLPValueComposition{SuperHalf: OutputSuperLP1, ElrondHalf: OutputSuperLP2, TotalUSD: OutputSuperLPValue}

	return EgldBought, OutputPrice
}

//
//
//======================================================================================================================
//
//
//[C]03         AddLiquidity computes the ESDT amount req. to add Liq. using a given amount of EGLD and Input Prices.
//              And computes the amount of Resulted LP (which is dependent on Input Prices)
//
//
func AddLiquidity(EgldAmount *p.Decimal, InputPrice MetaSuperPrice) (RequiredESDT, ResultedLP *p.Decimal) {
	ResultedLP = mt.TruncateCustom(mt.DIVxc(EgldAmount, InputPrice.SP.LPVC.ElrondHalf), 18)
	RequiredESDT = mt.TruncateCustom(mt.MULxc(ResultedLP, InputPrice.SP.LPVC.SuperHalf), 18)
	return RequiredESDT, ResultedLP
}

//
//
//======================================================================================================================
//======================================================================================================================
