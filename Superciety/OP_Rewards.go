package Superciety

import (
	mt "SuperPlasm/SuperMath"
	p "github.com/Crypt0plasm/Firefly-APD"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/Operations/OP_Rewards.go
//		Functions for Super Reward Definitions.
//
//
//		[A]Liquidity-Reward Program Functions
//[A]01		LpTierProcent			Returns SUPER-EGLD-LP Bonus % given by its Amount based on 7 Tiers.
//[A]02		CamelProcent			Returns SUPER-EGLD-LP Bonus % given by the Camel.
//[A]03		WeightBonus			Returns the combined Bonus % (Lp-Amount Bonus and Camel Bonus)
//[A]04		VirtualLP			Computers the Virtual LP
//[A]05		SuperVLPRewardComputer		Computes a single Super-Reward from a given VLP
//
//		[B]Super Power Computation Functions
//[B]01		MetaKosonicSuperPowerPromille	Computes the Promille needed to compute the meta-Super
//[B]02		SuperPowerComputerCore		Core Function that computes all 3 variants of Super-Power.
//[B]02a	SuperPowerComputer		Computes the 1st variant of Super-Power.
//[B]02b	KosonicSuperPowerComputer	Computes the 2nd variant of Super-Power.
//[B]02c	MetaKosonicSuperPowerComputer	Computes the 3rd variant of Super-Power.
//[B]03		ConvertSupersToMKSP		Outputs MKSP of a given Super, SuperLP, and Meta Combination.
//[B]04		GetAddyMKSP			Outputs the MKSP of a given Elrond Address.
//[B]05		ComputeMetaSuper		Computes meta-Super Amount.
//
//======================================================================================================================
//======================================================================================================================
//
//
//		[A]Liquidity-Reward Program Functions
//
//
//[A]01		LpTierProcent
//		Returns SUPER-EGLD-LP Bonus % given by its Amount based on 7 Tiers
//		This is the Definition of the SUPER-EGLD-LP Amount based Bonus %
//
//
func LpTierProcent(LP *p.Decimal) *p.Decimal {
	var (
		TierBonus = new(p.Decimal)
		T1        = p.NFI(1)
		T2        = p.NFI(2)
		T3        = p.NFI(5)
		T4        = p.NFI(10)
		T5        = p.NFI(20)
		T6        = p.NFI(50)
		T7        = p.NFI(100)

		NoBonus = p.NFI(0)
		BonusT1 = p.NFI(1)
		BonusT2 = p.NFI(2)
		BonusT3 = p.NFI(3)
		BonusT4 = p.NFI(4)
		BonusT5 = p.NFI(5)
		BonusT6 = p.NFI(6)
		BonusT7 = p.NFI(7)
	)

	if mt.DecimalLessThan(LP, T1) == true {
		TierBonus = NoBonus
	} else if mt.DecimalGreaterThanOrEqual(LP, T1) == true && mt.DecimalLessThan(LP, T2) == true {
		TierBonus = BonusT1
	} else if mt.DecimalGreaterThanOrEqual(LP, T2) == true && mt.DecimalLessThan(LP, T3) == true {
		TierBonus = BonusT2
	} else if mt.DecimalGreaterThanOrEqual(LP, T3) == true && mt.DecimalLessThan(LP, T4) == true {
		TierBonus = BonusT3
	} else if mt.DecimalGreaterThanOrEqual(LP, T4) == true && mt.DecimalLessThan(LP, T5) == true {
		TierBonus = BonusT4
	} else if mt.DecimalGreaterThanOrEqual(LP, T5) == true && mt.DecimalLessThan(LP, T6) == true {
		TierBonus = BonusT5
	} else if mt.DecimalGreaterThanOrEqual(LP, T6) == true && mt.DecimalLessThan(LP, T7) == true {
		TierBonus = BonusT6
	} else {
		TierBonus = BonusT7
	}

	return TierBonus
}

//
//
//======================================================================================================================
//
//
//[A]02		CamelProcent
//		Returns SUPER-EGLD-LP Bonus % given by the Camel.
//		If Camel is more than 0, Bonus is 10%
//		This is the Definition of the Camel based Bonus %
//
//
func CamelProcent(Camel *p.Decimal) *p.Decimal {
	var (
		CamelTotalBonus = new(p.Decimal)
		Bonus           = p.NFI(10)
	)

	if mt.DecimalGreaterThan(Camel, p.NFI(0)) == true {
		CamelTotalBonus = Bonus
	} else {
		CamelTotalBonus = p.NFI(1)
	}
	//CamelTotalBonus = mt.MULxc(Camel, Bonus)
	return CamelTotalBonus
}

//
//
//======================================================================================================================
//
//
//[A]03		WeightBonus
//		Returns the Bonus given by the combined:
//			SUPER-EGLD-LP Bonus % given by its Amount and
//			SUPER-EGLD-LP Bonus % given by the Camel
//		as multiplier (not as %)
//
//
func WeightBonus(TB, CB *p.Decimal) *p.Decimal {
	P1 := mt.ADDxc(p.NFI(1), mt.DIVxc(TB, p.NFI(100)))
	P2 := mt.ADDxc(p.NFI(1), mt.DIVxc(CB, p.NFI(100)))
	TP := mt.MULxc(P1, P2)

	return TP
}

//
//
//======================================================================================================================
//
//
//[A]04		VirtualLP
//		Returns the Virtual SUPER-EGLD-LP
//		Virtual SUPER-EGLD-LP is the SUPER-EGLD-LP multiplied by the Weight Bonus
//
//
func VirtualLP(LpAmount, CamelAmount *p.Decimal) *p.Decimal {
	TierBonus := LpTierProcent(LpAmount)
	CamelBonus := CamelProcent(CamelAmount)
	Weight := WeightBonus(TierBonus, CamelBonus)

	Result := mt.MULxc(LpAmount, Weight)
	return mt.TruncateCustom(Result, 18)
}

//
//
//======================================================================================================================
//
//
//[A]05		SuperVLPRewardComputer
//		Computes the Super reward earned by a given VLP
//			considering a VLP Chain and a given RewardAmount per Day.
//
//
func SuperVLPRewardComputer(VirtualLP *p.Decimal, VirtualLPChain []BalanceVLP, RewardAmount *p.Decimal) (Reward *p.Decimal) {
	var Sum = new(p.Decimal)
	for i := 0; i < len(VirtualLPChain); i++ {
		Balance := VirtualLPChain[i].VLPB
		Sum = mt.ADDxc(Sum, Balance)
	}
	Reward = mt.TruncateCustom(mt.DIVxc(mt.MULxc(VirtualLP, RewardAmount), Sum), 18)
	return Reward
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//		[B]Super Power Computation Functions
//
//
//[B]01		MetaKosonicSuperPowerPromille
//		Returns the Meta-Kosonic Super-Power Promille used to compute the meta-Super
//		Meta-Super is the "virtual" Super Amount used in calculating the Meta-Kosonic Super-Power
//
//
func MetaKosonicSuperPowerPromille(SuperAmount *p.Decimal) *p.Decimal {
	var (
		ProMille = new(p.Decimal)

		T00 = p.NFI(0)
		T01 = p.NFI(10)
		T02 = p.NFI(25)
		T03 = p.NFI(50)
		T04 = p.NFI(75)
		T05 = p.NFI(100)
		T06 = p.NFI(250)
		T07 = p.NFI(500)
		T08 = p.NFI(750)
		T09 = p.NFI(1000)
		T10 = p.NFI(2500)
		T11 = p.NFI(5000)
		T12 = p.NFI(7500)
		T13 = p.NFI(10000)
		T14 = p.NFI(25000)
		T15 = p.NFI(50000)
		T16 = p.NFI(75000)
		T17 = p.NFI(100000)
		T18 = p.NFI(250000)
		T19 = p.NFI(500000)
		T20 = p.NFI(750000)
		T21 = p.NFI(1000000)
		T22 = p.NFI(2500000)
		T23 = p.NFI(5000000)
		T24 = p.NFI(7500000)
		STS = p.NFI(10000000) //SuperTotalSupply.

		TP01 = p.NFI(10000)
		TP02 = p.NFI(8000)
		TP03 = p.NFI(7000)
		TP04 = p.NFI(6000)
		TP05 = p.NFI(5000)
		TP06 = p.NFI(4000)
		TP07 = p.NFI(3000)
		TP08 = p.NFI(2000)
		TP09 = p.NFI(1000)

		TP10 = mt.DIVxc(TP02, p.NFI(10)) //800
		TP11 = mt.DIVxc(TP03, p.NFI(10)) //700
		TP12 = mt.DIVxc(TP04, p.NFI(10)) //600
		TP13 = mt.DIVxc(TP05, p.NFI(10)) //500
		TP14 = mt.DIVxc(TP06, p.NFI(10)) //400
		TP15 = mt.DIVxc(TP07, p.NFI(10)) //300
		TP16 = mt.DIVxc(TP08, p.NFI(10)) //200
		TP17 = mt.DIVxc(TP09, p.NFI(10)) //100

		TP18 = mt.DIVxc(TP10, p.NFI(10)) //80
		TP19 = mt.DIVxc(TP11, p.NFI(10)) //70
		TP20 = mt.DIVxc(TP12, p.NFI(10)) //60
		TP21 = mt.DIVxc(TP13, p.NFI(10)) //50
		TP22 = mt.DIVxc(TP14, p.NFI(10)) //40
		TP23 = mt.DIVxc(TP15, p.NFI(10)) //30
		TP24 = mt.DIVxc(TP16, p.NFI(10)) //20
		TP25 = mt.DIVxc(TP17, p.NFI(10)) //10

		TPW0 = p.NFI(2000)
		TPW1 = p.NFI(1000)
		TPW2 = mt.DIVxc(TPW0, p.NFI(10)) //200
		TPW3 = mt.DIVxc(TPW1, p.NFI(10)) //100
		TPW4 = mt.DIVxc(TPW2, p.NFI(10)) //20
		TPW5 = mt.DIVxc(TPW3, p.NFI(10)) //10
	)

	//Rounding SuperAmount to IntegerValue
	SA := SuperAmount

	ProMilleFunc := func(StartThreshold, EndThreshold, TopPromille, WidthPromille, Value *p.Decimal) *p.Decimal {
		var (
			V1, V2, V3, V4, PM = new(p.Decimal), new(p.Decimal), new(p.Decimal), new(p.Decimal), new(p.Decimal)
		)
		V1 = mt.SUBxc(Value, StartThreshold)
		V2 = mt.SUBxc(EndThreshold, StartThreshold)
		V3 = mt.MULxc(V1, WidthPromille)
		V4 = mt.DIVxc(V3, V2)
		PM = mt.SUBxc(TopPromille, V4)
		return PM
	}

	if mt.DecimalGreaterThan(SA, T00) == true && mt.DecimalLessThanOrEqual(SA, T01) == true {
		ProMille = ProMilleFunc(T00, T01, TP01, TPW0, SA)
	} else if mt.DecimalGreaterThan(SA, T01) == true && mt.DecimalLessThanOrEqual(SA, T02) == true {
		ProMille = ProMilleFunc(T01, T02, TP02, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T02) == true && mt.DecimalLessThanOrEqual(SA, T03) == true {
		ProMille = ProMilleFunc(T02, T03, TP03, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T03) == true && mt.DecimalLessThanOrEqual(SA, T04) == true {
		ProMille = ProMilleFunc(T03, T04, TP04, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T04) == true && mt.DecimalLessThanOrEqual(SA, T05) == true {
		ProMille = ProMilleFunc(T04, T05, TP05, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T05) == true && mt.DecimalLessThanOrEqual(SA, T06) == true {
		ProMille = ProMilleFunc(T05, T06, TP06, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T06) == true && mt.DecimalLessThanOrEqual(SA, T07) == true {
		ProMille = ProMilleFunc(T06, T07, TP07, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T07) == true && mt.DecimalLessThanOrEqual(SA, T08) == true {
		ProMille = ProMilleFunc(T07, T08, TP08, TPW1, SA)
	} else if mt.DecimalGreaterThan(SA, T08) == true && mt.DecimalLessThanOrEqual(SA, T09) == true {
		ProMille = ProMilleFunc(T08, T09, TP09, TPW2, SA)
	} else if mt.DecimalGreaterThan(SA, T09) == true && mt.DecimalLessThanOrEqual(SA, T10) == true {
		ProMille = ProMilleFunc(T09, T10, TP10, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T10) == true && mt.DecimalLessThanOrEqual(SA, T11) == true {
		ProMille = ProMilleFunc(T10, T11, TP11, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T11) == true && mt.DecimalLessThanOrEqual(SA, T12) == true {
		ProMille = ProMilleFunc(T11, T12, TP12, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T12) == true && mt.DecimalLessThanOrEqual(SA, T13) == true {
		ProMille = ProMilleFunc(T12, T13, TP13, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T13) == true && mt.DecimalLessThanOrEqual(SA, T14) == true {
		ProMille = ProMilleFunc(T13, T14, TP14, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T14) == true && mt.DecimalLessThanOrEqual(SA, T15) == true {
		ProMille = ProMilleFunc(T14, T15, TP15, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T15) == true && mt.DecimalLessThanOrEqual(SA, T16) == true {
		ProMille = ProMilleFunc(T15, T16, TP16, TPW3, SA)
	} else if mt.DecimalGreaterThan(SA, T16) == true && mt.DecimalLessThanOrEqual(SA, T17) == true {
		ProMille = ProMilleFunc(T16, T17, TP17, TPW4, SA)
	} else if mt.DecimalGreaterThan(SA, T17) == true && mt.DecimalLessThanOrEqual(SA, T18) == true {
		ProMille = ProMilleFunc(T17, T18, TP18, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T18) == true && mt.DecimalLessThanOrEqual(SA, T19) == true {
		ProMille = ProMilleFunc(T18, T19, TP19, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T19) == true && mt.DecimalLessThanOrEqual(SA, T20) == true {
		ProMille = ProMilleFunc(T19, T20, TP20, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T20) == true && mt.DecimalLessThanOrEqual(SA, T21) == true {
		ProMille = ProMilleFunc(T20, T21, TP21, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T21) == true && mt.DecimalLessThanOrEqual(SA, T22) == true {
		ProMille = ProMilleFunc(T21, T22, TP22, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T22) == true && mt.DecimalLessThanOrEqual(SA, T23) == true {
		ProMille = ProMilleFunc(T22, T23, TP23, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T23) == true && mt.DecimalLessThanOrEqual(SA, T24) == true {
		ProMille = ProMilleFunc(T23, T24, TP24, TPW5, SA)
	} else if mt.DecimalGreaterThan(SA, T24) == true && mt.DecimalLessThanOrEqual(SA, STS) == true {
		ProMille = ProMilleFunc(T24, STS, TP25, TPW5, SA)
	}

	return ProMille
}

//
//
//======================================================================================================================
//
//
//[B]02		SuperPowerComputerCore
//		Is the Core Function that computes all 3 variants of Super-Power.
//		The Super-Power Variants are:
//			1)		Super-Power =      Super.Amount * (Super-EGLD-LP).Amount
//			2)Kosonic 	Super-Power = 	   Super.Amount * [LOG(base 2) from (Super-EGLD-LP).Amount]
//			3)Meta-Kosonic 	Super-Power = meta-Super.Amount * [LOG(base 2) from (Super-EGLD-LP).Amount]
//
//
func SuperPowerComputerCore(SuperAmount, LPAmount, LPThreshold *p.Decimal) *p.Decimal {
	var (
		SP                           = new(p.Decimal)
		FinalSP                      = new(p.Decimal)
		LiquidityBonus               = new(p.Decimal)
		SuperPowerExistenceThreshold = p.NFI(10)
	)

	//Checks if LP is less than the Threshold
	if mt.DecimalLessThanOrEqual(LPAmount, LPThreshold) == true {
		//Checks is Address has less than 1 Super. A minimum of 1 Super is required for Super-Power
		if mt.DecimalLessThan(SuperAmount, p.NFS("1")) == true {
			//In case address has less than 1 Super, SuperPower is zero
			SP = p.NFS("0")
		} else if mt.DecimalGreaterThanOrEqual(SuperAmount, p.NFS("1")) == true {
			//In case address has more than 1 Super or equal, SuperPower is non Zero
			//Because LP is less than Threshold, Liquidity Bonus is 1, thus
			//SuperPower is equal to SuperAmount
			SP = mt.TruncateCustom(SuperAmount, 0)
		}
		//Checks if LP is greater than or equal to the Threshold
	} else if mt.DecimalGreaterThan(LPAmount, LPThreshold) == true {
		//In case the address has a greater than Threshold LP Amount
		//Which Translates to a Liquidity Bonus greater than 1.
		//
		//If LPThreshold is 1 (Normal Super-Power is to be calculated)
		//"Liquidity Bonus" equals normal LP
		//If LPThreshold is 2 (Kosonic/Meta-Kosonic Super-Power is to be calculated)
		//"Liquidity Bonus" equals Log Base 2 of LP
		if mt.DecimalEqual(LPThreshold, p.NFI(1)) == true {
			LiquidityBonus = LPAmount
		} else if mt.DecimalEqual(LPThreshold, p.NFI(2)) == true {
			LiquidityBonus = mt.TruncateCustom(mt.Logarithm(LPThreshold, LPAmount), 18)
		}

		//And Thus SuperPower is computed by multiplying SuperAmount with "Liquidity Bonus"
		SP = mt.MULxc(SuperAmount, LiquidityBonus)
	}

	//Resulted SuperPower must be greater than "SuperPowerExistenceThreshold" to exist.
	//If it is lower than "SuperPowerExistenceThreshold", it is set automatically to 0.
	if mt.DecimalGreaterThanOrEqual(SP, SuperPowerExistenceThreshold) == true {
		FinalSP = SP
	} else {
		FinalSP = p.NFS("0")
	}

	return mt.TruncateCustom(FinalSP, 0)
}

//
//
//======================================================================================================================
//
//
//[B]02a	SuperPowerComputer
//		Computes the 1st variant of Super-Power:
//			1)		Super-Power =      Super.Amount * (Super-EGLD-LP).Amount
//
//
func SuperPowerComputer(SuperValue, LPValue *p.Decimal) *p.Decimal {
	//Threshold 1 means LP is left in its native state
	LPThresholdValue := p.NFS("1")
	Result := SuperPowerComputerCore(SuperValue, LPValue, LPThresholdValue)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]02b	KosonicSuperPowerComputer
//		Computes the 2nd variant of Super-Power:
//			2)Kosonic 	Super-Power = 	   Super.Amount * [LOG(base 2) from (Super-EGLD-LP).Amount]
func KosonicSuperPowerComputer(SuperValue, LPValue *p.Decimal) *p.Decimal {
	LPThresholdValue := p.NFS("2")
	Result := SuperPowerComputerCore(SuperValue, LPValue, LPThresholdValue)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]02c	MetaKosonicSuperPowerComputer
//		Computes the 3rd variant of Super-Power:
//			3)Meta-Kosonic 	Super-Power = meta-Super.Amount * [LOG(base 2) from (Super-EGLD-LP).Amount]
//
//
func MetaKosonicSuperPowerComputer(SuperValue, LPValue *p.Decimal) *p.Decimal {
	MetaSuper := ComputeMetaSuper(SuperValue)
	Result := KosonicSuperPowerComputer(MetaSuper, LPValue)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]03		ConvertSupersToMKSP
//		Outputs MKSP of a given Super, SuperLP, and Meta Combination.
//
//
func ConvertSupersToMKSP(SuperValue, LPValue *p.Decimal, Meta bool) (MKSP *p.Decimal) {
	if Meta == true {
		MKSP = MetaKosonicSuperPowerComputer(SuperValue, LPValue)
	} else {
		MKSP = KosonicSuperPowerComputer(SuperValue, LPValue)
	}
	return MKSP
}

//
//
//======================================================================================================================
//
//
//[B]04		GetAddyMKSP
//		Outputs the MKSP of a given Elrond Address.
//
//
func GetAddyMKSP(Addy ElrondAddress) (MKSP *p.Decimal) {
	MySuper, MySuperLp := GetAddySuperValues(Addy)
	SFT1Chain := CreateCamelChain()
	GetMeta := IzMeta(Addy, SFT1Chain)
	MKSP = ConvertSupersToMKSP(MySuper, MySuperLp, GetMeta)
	return MKSP
}

//
//
//======================================================================================================================
//
//
//[B]05		ComputeMetaSuper
//		Computes the meta-Super using the MetaKosonicSuperPowerPromille
//
//
func ComputeMetaSuper(Super *p.Decimal) *p.Decimal {
	MetaKosonicPromille := MetaKosonicSuperPowerPromille(Super)
	PurePromille := mt.DIVxc(MetaKosonicPromille, p.NFI(1000))
	PromilleMultiplier := mt.ADDxc(p.NFI(1), PurePromille)
	MetaSuper := mt.MULxc(Super, PromilleMultiplier)
	return MetaSuper
}

//
//
//======================================================================================================================
//======================================================================================================================
