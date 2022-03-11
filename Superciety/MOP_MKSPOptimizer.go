package Superciety

import (
	mt "SuperPlasm/SuperMath"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/..MetaOperations../MOP_MKSPOptimizer.go
//		Functions that are used for the MKSP-Optimizer Meta-Operation.
//
//
//[1]		SuperLPMelter		Melts the LP to Super to ascertain if a greater MKSP can be achieved.
//[2]		SuperMelter		Melts the Super SuperLP to ascertain if a greater MKSP can be achieved.
//[3]		SuperLPtoSuperConvertor	Displays Information regarding Super-LP to Super Conversion.
//[4]		SuperToSuperLPConvertor	Displays Information regarding Super to Super-LP Conversion.
//[5]		Optimizer		The Main Function for this Meta-Operation.
//
//======================================================================================================================
//======================================================================================================================
//
//
//[1]		SuperLPMelter
//		Melts the LP to Super to ascertain if a greater MKSP can be achieved.
//
//
func SuperLPMelter(Addy ElrondAddress, CurrentPrice MetaSuperPrice) (OptimalLpReduction, MaxSuperPower *p.Decimal) {
	//Obtaining Address Specifications
	//These represent Super amount, SuperLP Amount, and Meta.
	var (
		MaxIteration int64
	)
	MySuper, MySuperLp := GetAddySuperValues(Addy)
	SFT1Chain := CreateCamelChain()
	GetMeta := IzMeta(Addy, SFT1Chain)

	//Address Meta-kosonic Super-Power
	OriginalMKSP := ConvertSupersToMKSP(MySuper, MySuperLp, GetMeta)

	//Initializing MaxSuperPower as the Address MKSP
	MaxSuperPower = OriginalMKSP

	//Computes how many Melt Steps must be calculated.
	//One Melt Step is the equivalent of 10 USD
	//This is how many iterations a new MKSP will be computed for.
	fmt.Println("My SuperLP is ", MySuperLp)
	fmt.Println("SuperLp Price is", CurrentPrice.SP.LPVC.TotalUSD)
	MeltSteps := mt.TruncateCustom(mt.DIVxc(mt.MULxc(MySuperLp, CurrentPrice.SP.LPVC.TotalUSD), p.NFI(10)), 0)
	fmt.Println(MeltSteps, "Liquidity melt steps must be checked:")

	//Depending on the number of Melt-Steps, the amount of LP that gets subtracted for each step is computed.
	//This amount is broken into Super and EGLD, EGLD is used to "virtually" buy Super.
	//The Super resulted from breaking the LP (Super and Super bought with EGLD)
	//Is added on top of the original Super Amount
	//And a new MKSP is computed.
	SubtractingLP := mt.DIVxc(MySuperLp, MeltSteps)

	//Virtual Values are values that are virtually created by the SuperLPMelter
	//VirtualPrice are the Prices Resulted when LP is melted and used to buy Super.
	//	Melting LP Creates Super and EGLD
	//	Super is added to original Address Super
	//	Egld is virtually swapped to Super, thus virtually increasing Super prices.
	VirtualPrices := CurrentPrice
	//Virtual Super is the Super obtained after a MeltStep
	//It starts as MySuper
	ObtainedVirtualSuper := MySuper
	//RemainingVirtualLp is the SuperLP remaining after a MeltStep
	//It starts as MySuperLP
	RemainingVirtualLP := MySuperLp

	for i := int64(0); i < p.INT64(MeltSteps); i++ {
		var (
			SuperGained1, SuperGained2 = new(p.Decimal), new(p.Decimal)
		)

		BaseStringPoint := "Super-LP Melting Iteration "
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, i+1, "/", MeltSteps)

		RemainingVirtualLP = mt.TruncateCustom(mt.SUBxc(RemainingVirtualLP, SubtractingLP), 18)
		//Subtracting the subtracted LP to the total LP existing in the VirtualPrices.
		VirtualPrices.Liquidity = mt.TruncateCustom(mt.SUBxc(VirtualPrices.Liquidity, SubtractingLP), 18)

		//What is gained from Subtracted Amount
		SuperGained1 = mt.TruncateCustom(mt.MULxc(VirtualPrices.SP.LPVC.SuperHalf, SubtractingLP), 18)
		EgldGained := mt.TruncateCustom(mt.MULxc(VirtualPrices.SP.LPVC.ElrondHalf, SubtractingLP), 18)

		SuperGained2, VirtualPrices = BuySuper(EgldGained, VirtualPrices)
		TotalSuperGained := mt.TruncateCustom(mt.ADDxc(SuperGained1, SuperGained2), 18)

		ObtainedVirtualSuper = mt.TruncateCustom(mt.ADDxc(ObtainedVirtualSuper, TotalSuperGained), 18)

		//Getting the Virtual MKSP used as base for comparison.
		VirtualMKSP := ConvertSupersToMKSP(ObtainedVirtualSuper, RemainingVirtualLP, GetMeta)
		MaxSuperPower = mt.MaxDecimal(VirtualMKSP, MaxSuperPower)
		if mt.DecimalGreaterThan(MaxSuperPower, VirtualMKSP) == false {
			MaxIteration = i + 1
		}
	}
	if MaxIteration == 0 {
		OptimalLpReduction = p.NFI(0)
	} else {
		OptimalLpReduction = mt.TruncateCustom(mt.MULxc(SubtractingLP, p.NFI(MaxIteration)), 18)
	}

	return OptimalLpReduction, MaxSuperPower

}

//
//
//======================================================================================================================
//
//
//[2]		SuperMelter
//		Melts the Super SuperLP to ascertain if a greater MKSP can be achieved.
//
//
func SuperMelter(Addy ElrondAddress, CurrentPrice MetaSuperPrice) (OptimalSuperReduction, MaxSuperPower *p.Decimal) {
	//Obtaining Address Specifications
	//These represent Super amount, SuperLP Amount, and Meta.
	var (
		MaxIteration int64
	)
	MySuper, MySuperLp := GetAddySuperValues(Addy)
	SFT1Chain := CreateCamelChain()
	GetMeta := IzMeta(Addy, SFT1Chain)

	//Address Meta-kosonic Super-Power
	OriginalMKSP := ConvertSupersToMKSP(MySuper, MySuperLp, GetMeta)

	//Initializing MaxSuperPower as the Address MKSP
	MaxSuperPower = OriginalMKSP

	//Computes how many Melt Steps must be calculated.
	//One Melt Step is the equivalent of 10 USD
	//This is how many iterations a new MKSP will be computed for.
	fmt.Println("My Super is ", MySuper)
	fmt.Println("Super Price is", CurrentPrice.SP.SV.USDperSUPER)
	MeltSteps := mt.TruncateCustom(mt.DIVxc(mt.MULxc(MySuper, CurrentPrice.SP.SV.USDperSUPER), p.NFI(10)), 0)
	fmt.Println(MeltSteps, "Super melt steps must be checked:")

	//Depending on the number of Melt-Steps, the amount of Super that gets subtracted for each step is computed.
	//This amount is broken into 2 equal Super Parts. Half is kept, half is swapped for EGLD.
	//The EGLD is paired with the half of the Super, any remaining Super is added with the Base Super
	//There is a small amount of remaining Super that cant be paired with EGLD, because we cant swap for the
	//	perfect Amount of EGLD so that no Super remains
	//With the resulted LP a new MKSP is computed
	SubtractingSuper := mt.DIVxc(MySuper, MeltSteps)

	//Virtual Values are values that are virtually created by the SuperMelter
	//VirtualPrice are the Prices Resulted when half of the Super melted is swapped to EGLD
	//	Melting Super is used in equal Parts to create LP
	//	LP is added to the original Address LP
	//	Super is virtually swapped to EGLD, thus virtually decreasing Super prices.
	VirtualPrices := CurrentPrice
	//Virtual LP is the SuperLP obtained after a MeltStep
	//It starts as MySuperLP
	ObtainedVirtualSuperLP := MySuperLp
	//RemainingVirtualSuper is the Super remaining after a MeltStep
	//It starts as MySuper
	RemainingVirtualSuper := MySuper

	for i := int64(0); i < p.INT64(MeltSteps); i++ {
		var (
			EgldGained    = new(p.Decimal)
			ConsumedSuper = new(p.Decimal)
			ResultedLP    = new(p.Decimal)
			RestSuper     = new(p.Decimal)
		)
		BaseStringPoint := "Super    Melting Iteration "
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, i+1, "/", MeltSteps)

		RemainingVirtualSuper = mt.SUBxc(RemainingVirtualSuper, SubtractingSuper)
		//No Subtraction step must be taken here as opposed to the SuperLPMelter

		//What is gained from Subtracted Amount
		HalfSuper := mt.TruncateCustom(mt.DIVxc(SubtractingSuper, p.NFI(2)), 18)
		EgldGained, VirtualPrices = SellSuper(HalfSuper, VirtualPrices)
		ConsumedSuper, ResultedLP = AddLiquidity(EgldGained, VirtualPrices)

		//Computing remains after Liq add
		//	This is the SuperAmount used in next iteration
		RestSuper = mt.TruncateCustom(mt.SUBxc(HalfSuper, ConsumedSuper), 18)
		RemainingVirtualSuper = mt.ADDxc(RemainingVirtualSuper, RestSuper)
		//	This is the Increased VirtualLP
		ObtainedVirtualSuperLP = mt.TruncateCustom(mt.ADDxc(ObtainedVirtualSuperLP, ResultedLP), 18)

		VirtualMKSP := ConvertSupersToMKSP(RemainingVirtualSuper, ObtainedVirtualSuperLP, GetMeta)
		MaxSuperPower = mt.MaxDecimal(VirtualMKSP, MaxSuperPower)
		if mt.DecimalGreaterThan(MaxSuperPower, VirtualMKSP) == false {
			MaxIteration = i + 1
		}
	}
	if MaxIteration == 0 {
		OptimalSuperReduction = p.NFI(0)
	} else {
		OptimalSuperReduction = mt.TruncateCustom(mt.MULxc(SubtractingSuper, p.NFI(MaxIteration)), 18)
	}
	return OptimalSuperReduction, MaxSuperPower

}

//
//
//======================================================================================================================
//
//
//[3]		SuperLPtoSuperConvertor
//		Displays Information regarding LP to Super Conversion.
//
//
func SuperLPtoSuperConvertor(InitialLP, UsedLP *p.Decimal, Prices MetaSuperPrice) {
	Super1 := mt.TruncateCustom(mt.MULxc(UsedLP, Prices.SP.LPVC.SuperHalf), 18)
	Elrond := mt.TruncateCustom(mt.MULxc(UsedLP, Prices.SP.LPVC.ElrondHalf), 18)
	SuperBought, NewPrices := BuySuper(Elrond, Prices)
	SuperGained := mt.ADDxc(Super1, SuperBought)
	RemainingLP := mt.SUBxc(InitialLP, UsedLP)

	//P1 := Prices.SP.SV.USDperSUPER
	//P2 := NewPrices.SP.SV.USDperSUPER
	//PP := PercentSwing(P1,P2)

	//Printing Data
	fmt.Println("STEP0:          from the:", InitialLP, "SUPER-EGLD-LP")
	fmt.Println("")
	fmt.Println("STEP1:            remove:", UsedLP, "SUPER-EGLD-LP which")
	fmt.Println("STEP2:           creates:", Super1, "SUPER")
	fmt.Println("STEP3:               and:", Elrond, "EGLD; use this EGLD")
	fmt.Println("STEP4:            to buy:", SuperBought, "SUPER;")
	fmt.Println("")
	fmt.Println("STEP5:        you gained:", SuperGained, "SUPER anew.")
	fmt.Println("STEP6:          and have:", RemainingLP, "SUPER-EGLD-LP left.")
	fmt.Println("")
	fmt.Println("=======")
	//fmt.Println("Increased the price by ",PP,"%")
	fmt.Println("Price movement from:")
	fmt.Println(Prices.SP.SV.USDperSUPER, "USD to:")
	fmt.Println(NewPrices.SP.SV.USDperSUPER, "USD.")
}

//
//
//======================================================================================================================
//
//
//[4]		SuperToSuperLPConvertor
//		Displays Information regarding Super to Super-LP Conversion.
//
//
func SuperToSuperLPConvertor(InitialSuper, UsedSuper *p.Decimal, Prices MetaSuperPrice) {
	HalfSuper := mt.TruncateCustom(mt.DIVxc(UsedSuper, p.NFI(2)), 18)
	EgldGained, NewPrices := SellSuper(HalfSuper, Prices)
	ConsumedSuper, ResultedLP := AddLiquidity(EgldGained, NewPrices)
	RestSuper := mt.TruncateCustom(mt.SUBxc(HalfSuper, ConsumedSuper), 18)
	RemainingSuper := mt.ADDxc(mt.SUBxc(InitialSuper, UsedSuper), RestSuper)

	//P1 := Prices.SP.SV.USDperSUPER
	//P2 := NewPrices.SP.SV.USDperSUPER
	//PP := PercentSwing(P1,P2)

	//Printing Data
	fmt.Println("STEP0,           from the:", InitialSuper, "SUPER")
	fmt.Println("")
	fmt.Println("STEP1:               use:", HalfSuper, "SUPER to swap to EGLD;")
	fmt.Println("STEP2:               use:", EgldGained, "EGLD to add Liquidity;")
	fmt.Println("STEP3:    you now gained:", ResultedLP, "SUPER-EGLD-LP anew.")
	fmt.Println("")
	fmt.Println("STEP4: you are left with:", RemainingSuper, "SUPER")
	fmt.Println("")
	fmt.Println("========")
	//fmt.Println("Dropped the price by ",PP,"%")
	fmt.Println("Price movement from:")
	fmt.Println(Prices.SP.SV.USDperSUPER, "USD to:")
	fmt.Println(NewPrices.SP.SV.USDperSUPER, "USD.")
}

//
//
//======================================================================================================================
//
//
//[5]		Optimizer
//		The Main Optimizer Function.
//
//
func Optimizer(Addy ElrondAddress) {
	ScannedSuperPrices := GetAllSuperPrices()
	PricePrinter(ScannedSuperPrices)
	AddySuper, AddyLP, MKSP := AddySpecs(Addy)
	AddySpecsPrinter(Addy)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("==========================LiquidityMelter=================================")
	LPReduction, Value := SuperLPMelter(Addy, ScannedSuperPrices)
	if mt.DecimalEqual(MKSP, Value) == false {
		MKSPGain := mt.SUBxc(Value, MKSP)
		fmt.Println("======================END-LiquidityMelter=================================")
		fmt.Println("")

		fmt.Println("==========================Results=========================================")
		SuperLPtoSuperConvertor(AddyLP, LPReduction, ScannedSuperPrices)
		fmt.Println("")
		fmt.Println("If you reduce your LP by", LPReduction)
		fmt.Println("You will increase your MKSP from", MKSP, "to", Value)
		fmt.Println("which means a gain of", MKSPGain)
		fmt.Println("======================END-Results=========================================")
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Println("======================END-LiquidityMelter=================================")
		fmt.Println("")
		fmt.Println("==========================SuperMelter=====================================")
		SuperReduction, Value2 := SuperMelter(Addy, ScannedSuperPrices)
		if mt.DecimalEqual(MKSP, Value2) == false {
			MKSPGain2 := mt.SUBxc(Value2, MKSP)
			fmt.Println("")
			fmt.Println("======================END-SuperMelter=====================================")
			fmt.Println("")

			fmt.Println("==========================Results=========================================")
			SuperToSuperLPConvertor(AddySuper, SuperReduction, ScannedSuperPrices)
			fmt.Println("")
			fmt.Println("If you use ", SuperReduction, "SUPER to buy LP")
			fmt.Println("You will increase your MKSP from", MKSP, "to", Value2)
			fmt.Println("which means a gain of", MKSPGain2)
			fmt.Println("======================END-Results=========================================")
			fmt.Println("")
		} else {
			fmt.Println("")
			fmt.Println("======================END-SuperMelter=====================================")
			fmt.Println("")

			fmt.Println("==========================Results=========================================")
			fmt.Println("Your Super and LP amounts are optimal for Maximum MKSP")
			fmt.Println("======================END-Results=========================================")
			fmt.Println("")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

//======================================================================================================================
//======================================================================================================================
