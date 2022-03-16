package Superciety

import (
	mt "SuperPlasm/SuperMath"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"reflect"
	"runtime"
	"strings"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/Operations/OP_ArtificialChains.go
//		Functions that create artificial Chains from Blockchain Scanned Values.
//
//
//		[A]VirtualLP artificial Chain Functions:
//[A]01 	CreateVLPChain			 Creates a Chain of Virtual (Super-EGLD-LP).
//[A]02		SuperVLPRewardComputerChain	 VLP based Super rewards computer.
//
//		[B]Super-Power artificial Chain Functions:
//[B]01		CreateSuperPowerChainCore	 The Core Function that is used to create a Chain of Super-Power values.
//[B]01a	CreateSuperPowerChain		 Creates a chain of Super-Power values.
//[B]01b	CreateKosonicSuperPowerChain	 Creates a chain of Kosonic Super-Power values.
//[B]01c	CreateMetaKosonicSuperPowerChain Creates a chain of Meta-Kosonic SuperPower values
//[B]02a	SuperPowerPercentComputer	 Creates unsorted Super-Power Chains with % displays.
//[B]02b	SortSuperPowerPercentChain	 Sorts the Super-Power % Chain.
//
//======================================================================================================================
//======================================================================================================================
//
//
//		[A]VirtualLP artificial Chain Functions
//
//
//[A]01 	CreateVLPChain
//		Creates a Chain of Virtual (Super-EGLD-LP)
//		Virtual (Super-EGLD-LP) is a weighted Super-EGLD-LP by its Amount based on Tiers and Camel Bonus.
//
//
func CreateVLPChain(Chain1 []BalanceSuperLP, Chain2 []BalanceCamel) []BalanceVLP {
	var FinalChain []BalanceVLP
	for i := 0; i < len(Chain1); i++ {
		if Chain1[i].Address == ExA1 || mt.DecimalLessThan(ConvertAU18(Chain1[i].SuperLpB), p.NFS("0.5")) == true {
			//Unit := SuperVLP{Chain1[i].Address, p.NFS("0")}
			//FinalChain = append(FinalChain, Unit)
		} else {
			LPAmount := ConvertAU18(Chain1[i].SuperLpB)
			Camels := p.NFS(GetCamelAmount(Chain1[i].Address, Chain2))
			VLP := VirtualLP(LPAmount, Camels)
			Unit := BalanceVLP{Address: Chain1[i].Address, VLPB: VLP}
			FinalChain = append(FinalChain, Unit)
		}
	}
	return FinalChain
}

//
//
//======================================================================================================================
//
//
//[A]02		SuperVLPRewardComputerChain
//		Computes the Super rewards earned by a all VLPs in a VLP Chain given RewardAmount per Day
//  		and creates a Chain with the Results
//
//
func SuperVLPRewardComputerChain(Chain1 []BalanceVLP, RewardAmount *p.Decimal) []BalanceSFR {
	var (
		VLPSum     = new(p.Decimal)
		FinalChain []BalanceSFR
	)
	for i := 0; i < len(Chain1); i++ {
		VLPSum = mt.ADDxc(VLPSum, Chain1[i].VLPB)
	}
	for i := 0; i < len(Chain1); i++ {
		Reward := mt.TruncateCustom(mt.DIVxc(mt.MULxc(Chain1[i].VLPB, RewardAmount), VLPSum), 18)
		Unit := BalanceSFR{Chain1[i].Address, Reward}
		FinalChain = append(FinalChain, Unit)
	}
	return FinalChain
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//		[B]Super-Power artificial Chain Functions
//
//
//[B]01		CreateSuperPowerChainCore Function
//		The Core Function that is used to create a Chain of Super-Power values.
//		Super-Power values are:
//			1)Super-Power
//			2)Kosonic Super-Power
//			3)Meta-Kosonic Super-Power
//		Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
//
func CreateSuperPowerChainCore(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP, SuperPowerFunction func(*p.Decimal, *p.Decimal) *p.Decimal) []MetaKosonicSuperPower {
	var (
		FinalChain      []MetaKosonicSuperPower
		GetMeta         = false
		SuperPowerValue = new(p.Decimal)
	)

	//MetaCheck Snapshots - used for Meta-Kosonic Super-Power
	//Multiple Chains can be added if multiple SFTs must be checked
	//Remember to add Checks in the IzMeta Function as well
	SFT1Chain := CreateCamelChain()

	//Getting a Function name.
	GetFunctionName := func(temp interface{}) string {
		Value := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
		return Value[len(Value)-1]
	}

	for i := 0; i < len(Chain1); i++ {
		var MetaSuperAmount = new(p.Decimal)

		BaseStringPoint := "Iteration"
		//StringPoint := strings.Repeat(".",i)
		//StringToPrint := BaseStringPoint + StringPoint
		fmt.Print("\r", BaseStringPoint, " ", i)

		//Work only if Address is non Exception
		if ComputeExceptionAddress(Chain1[i].Address) == false {
			//0)Address is Chain1[i].Address

			//1)Getting the Super Value
			SuperAmount := ConvertAU18(Chain1[i].SuperB)
			TruncatedSuperAmount := mt.TruncateCustom(SuperAmount, 0)
			//Integers the non-integer Super

			//2)Getting the LP Amount
			LPAmount := GetSuperLPAmount(Chain1[i].Address, Chain2)

			//3)Computing MetaSuper and SuperPower
			if GetFunctionName(SuperPowerFunction) == "MetaKosonicSuperPowerComputer" {
				GetMeta = IzMeta(Chain1[i].Address, SFT1Chain)
				if GetMeta == true {
					//if meta is true, SuperPower applies the input Function,
					//which is in this case "MetaKosonicSuperPowerComputer"
					//It has built in Super to meta-Super conversion
					//That is why it is used with SuperAmount
					MetaSuperAmount = ComputeMetaSuper(SuperAmount)
					SuperPowerValue = SuperPowerFunction(SuperAmount, LPAmount)
				} else {
					//if meta is false, SuperPower applies doesnt apply the input Function
					//but applies the KosonicSuperPowerComputer function
					//because this doesnt use meta-Super
					MetaSuperAmount = SuperAmount
					SuperPowerValue = KosonicSuperPowerComputer(SuperAmount, LPAmount)
				}
			} else {
				//case where non Meta-Kosonic Super-Power has to be calculated
				//namely the Super-Power and Kosonic Super-Power
				MetaSuperAmount = SuperAmount
				SuperPowerValue = SuperPowerFunction(SuperAmount, LPAmount)
			}

			//Truncating the meta-Super since it must be integer
			TruncatedMetaSuperAmount := mt.TruncateCustom(MetaSuperAmount, 0)

			//Creating the Chain element. Only SuperPower values greater than 0 are added to the chain.
			//Since the SuperPower computing Function sets the SuperPower Result to 0
			//if it is below the "SuperPowerExistenceThreshold" this code here
			//Incorporates in Chain only non zero-values.

			if mt.DecimalGreaterThan(SuperPowerValue, p.NFS("0")) == true {
				Unit := MetaKosonicSuperPower{Address: Chain1[i].Address, Super: TruncatedSuperAmount, MetaSuper: TruncatedMetaSuperAmount, SuperPower: SuperPowerValue}
				FinalChain = append(FinalChain, Unit)
			}
		}
	}
	fmt.Println("")
	return FinalChain
}
func ComputeExceptionAddress(Addy ElrondAddress) bool {
	Exceptions := []ElrondAddress{ExA1, ExA2, ExA3, ExA4}
	var Result = false
	for i := 0; i < len(Exceptions); i++ {
		if Addy == Exceptions[i] {
			Result = true
		}
	}
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]01a	CreateSuperPowerChain Function
//		Creates a chain of Super-Power values
//		Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
//
func CreateSuperPowerChain(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MetaKosonicSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, SuperPowerComputer)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]01b	CreateKosonicSuperPowerChain Function
//		Creates a chain of Kosonic Super-Power values
//		Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
//
func CreateKosonicSuperPowerChain(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MetaKosonicSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, KosonicSuperPowerComputer)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]01c	CreateMetaKosonicSuperPowerChain Function
//		Creates a chain of Meta-Kosonic SuperPower values
//		Chain1 is a Super Chain, Chain2 is SUPER-EGLD-LP Chain
//
//
func CreateMetaKosonicSuperPowerChain(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MetaKosonicSuperPower {
	Result := CreateSuperPowerChainCore(Chain1, Chain2, MetaKosonicSuperPowerComputer)
	return Result
}

//
//
//======================================================================================================================
//
//
//[B]02a	SuperPowerPercentComputer
//		Creates an unsorted Chain with the % Values of each address SUPER-Power
//		Works for all Super-Power variants.
//
//
func SuperPowerPercentComputer(Chain []MetaKosonicSuperPower) []MKSuperPowerPercent {
	var (
		SPSum      = new(p.Decimal)
		FinalChain []MKSuperPowerPercent
	)
	for i := 0; i < len(Chain); i++ {
		SPSum = mt.ADDxc(SPSum, Chain[i].SuperPower)
	}
	for i := 0; i < len(Chain); i++ {
		Percent := mt.TruncateCustom(mt.DIVxc(mt.MULxc(Chain[i].SuperPower, p.NFS("100")), SPSum), 18)
		Unit := MKSuperPowerPercent{Main: Chain[i], MetaKosonicSuperPowerPercent: Percent}
		FinalChain = append(FinalChain, Unit)
	}
	return FinalChain
}

//
//
//======================================================================================================================
//
//
//[B]02b	SortSuperPowerPercentChain
//		Sorts the Super-Power % Chain from highest % to lowest %
//		Works for all Super-Power variants.
//
//
func SortSuperPowerPercentChain(Chain []MKSuperPowerPercent) []MKSuperPowerPercent {
	var (
		SortedChain []MKSuperPowerPercent
	)
	GetMaxElement := func(Chain []MKSuperPowerPercent) int {
		Max := 0
		for i := 0; i < len(Chain); i++ {
			if mt.DecimalGreaterThanOrEqual(Chain[i].MetaKosonicSuperPowerPercent, Chain[Max].MetaKosonicSuperPowerPercent) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := MKSuperPowerPercent{Main: Chain2Sort[Biggest].Main, MetaKosonicSuperPowerPercent: Chain2Sort[Biggest].MetaKosonicSuperPowerPercent}
		SortedChain = append(SortedChain, Unit)

		//Removing biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)
	}
	return SortedChain
}

//
//
//======================================================================================================================
//======================================================================================================================
