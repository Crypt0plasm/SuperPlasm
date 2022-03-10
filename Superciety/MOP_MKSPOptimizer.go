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
	MeltSteps := mt.TruncateCustom(mt.DIVxc(mt.MULxc(MySuperLp, CurrentPrice.SP.LPVC.TotalUSD), p.NFI(10)), 0)

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
		fmt.Print("\r", BaseStringPoint, i, "/", mt.SUBxc(MeltSteps, p.NFI(1)))

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
	fmt.Println("")
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
	MeltSteps := mt.TruncateCustom(mt.DIVxc(mt.MULxc(MySuper, CurrentPrice.SP.SV.USDperSUPER), p.NFI(10)), 0)

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
		fmt.Print("\r", BaseStringPoint, i, "/", mt.SUBxc(MeltSteps, p.NFI(1)))

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
	fmt.Println("")
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

	//P1 := Prices.SP.SV.USDperSUPER
	//P2 := NewPrices.SP.SV.USDperSUPER
	//PP := PercentSwing(P1,P2)

	//Printing Data
	fmt.Println("Initial      SUPER-EGLD-LP:", InitialLP, "SUPER")
	fmt.Println("SwapQuantity SUPER-EGLD-LP:", UsedLP, "SUPER-EGLD-LP yields:")
	fmt.Println("Following Liquidity Amount:", Super1, "SUPER and,")
	fmt.Println("Following Liquidity Amount:", Elrond, "EGLD, which buys:")
	fmt.Println("Bought    Liquidity Amount:", SuperBought, "SUPER")
	fmt.Println("For a   TOTAL Liquidity of:", SuperGained, "SUPER")
	fmt.Println("=======")
	fmt.Println("Buying Super with EGLD to increase MKSP")
	//fmt.Println("Increased the price by ",PP,"%")
	fmt.Println("The Price increased from:")
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
	fmt.Println("Initial              SUPER:", InitialSuper, "SUPER")
	fmt.Println("Spending             SUPER:", UsedSuper, "SUPER")
	fmt.Println("SwapQuantity         SUPER:", HalfSuper, "SUPER to EGLD")
	fmt.Println("Add Liquidity with    EGLD:", EgldGained, "EGLD")
	fmt.Println("To create    SUPER-EGLD-LP:", ResultedLP, "SUPER-EGLD-LP")
	fmt.Println("Remaining amount     SUPER:", RemainingSuper, "SUPER")
	fmt.Println("=======")

	fmt.Println("Selling Super for EGLD to add Liquidity")
	//fmt.Println("Dropped the price by ",PP,"%")
	fmt.Println("The Price dropped from:")
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
	AddySuper, AddyLP, MKSP := AddySpecs(Addy)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("===========Optimizer==================")

	LPReduction, Value := SuperLPMelter(Addy, ScannedSuperPrices)
	if mt.DecimalGreaterThan(LPReduction, p.NFS("0")) == true {
		fmt.Println("OldMKSP is", MKSP)
		fmt.Println("NewMKSP is", Value)
		MKSPGain := mt.SUBxc(Value, MKSP)
		fmt.Println("===========Results====================")
		SuperLPtoSuperConvertor(AddyLP, LPReduction, ScannedSuperPrices)
		fmt.Println("")
		fmt.Println("If you reduce your LP by", LPReduction)
		fmt.Println("You will increase MKSP by", MKSPGain)
		fmt.Println("")
	} else {
		SuperReduction, Value2 := SuperMelter(Addy, ScannedSuperPrices)
		fmt.Println("===========Results====================")
		if mt.DecimalGreaterThan(SuperReduction, p.NFS("0")) == true {
			fmt.Println("OldMKSP is", MKSP)
			fmt.Println("NewMKSP is", Value2)
			MKSPGain2 := mt.SUBxc(Value2, MKSP)
			SuperToSuperLPConvertor(AddySuper, SuperReduction, ScannedSuperPrices)
			fmt.Println("")
			fmt.Println("If you use ", SuperReduction, "SUPER to buy LP")
			fmt.Println("You will increase MKSP by", MKSPGain2)
			fmt.Println("")
		} else {
			fmt.Println("Your Super and LP amounts are optimal for Maximum MKSP")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

//======================================================================================================================
//======================================================================================================================
