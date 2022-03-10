package Superciety

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	N06 = "Computed_SUPER-Power-variant1-PS.txt"
	N07 = "Computed_SUPER-Power-variant2-PS.txt"
	N08 = "Computed_SUPER-Power-variant3-PS.txt"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/..MetaOperations../MOP_SuperPowers.go
//		Functions that are used for the SuperPowers Meta-Operation.
//
//
//[1]           SnapshooterPrinterSuperPower                Computes and prints a Super-Power Chain.
//[2]           SnapshooterPrinterKosonicSuperPower         Computes and prints a Kosonic Super-Power Chain.
//[3]           SnapshooterPrinterMetaKosonicSuperPower     Computes and prints a Meta-Kosonic Super-Power Chain.
//[4]           SuperPowers                                 The Main Function for this Meta-Operation.
//
//======================================================================================================================
//======================================================================================================================
//
//
//[1]           SnapshooterPrinterSuperPower
//              Computes and prints a Super-Power Chain.
//
//
func SnapshooterPrinterSuperPower(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MKSuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing SUPER-Power Chain ...")
	Start1 := time.Now()
	SuperPowerChain := CreateSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing  SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	SuperPowerPercentChain := SuperPowerPercentComputer(SuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting  SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedSuperPowerPercentChain := SortSuperPowerPercentChain(SuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting  SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted SUPER-Power-Percent-Chain to", N06)
	OutputFile, err := os.Create(N06)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SortedSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted SUPER-Power-Percent-Chain to", N06)
	fmt.Println("")

	return SortedSuperPowerPercentChain

}

//
//
//======================================================================================================================
//
//
//[2]           SnapshooterPrinterKosonicSuperPower
//              Computes and prints a Kosonic Super-Power Chain.
//
//
func SnapshooterPrinterKosonicSuperPower(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MKSuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing Kosonic SUPER-Power Chain ...")
	Start1 := time.Now()
	KosonicSuperPowerChain := CreateKosonicSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  Kosonic SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing Kosonic SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	KosonicSuperPowerPercentChain := SuperPowerPercentComputer(KosonicSuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  Kosonic SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting Kosonic SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedKosonicSuperPowerPercentChain := SortSuperPowerPercentChain(KosonicSuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting Kosonic SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted Kosonic SUPER-Power-Percent-Chain to", N07)
	OutputFile, err := os.Create(N07)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SortedKosonicSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted Kosonic SUPER-Power-Percent-Chain to", N07)
	fmt.Println("")

	return SortedKosonicSuperPowerPercentChain

}

//
//
//======================================================================================================================
//
//
//[3]           SnapshooterPrinterMetaKosonicSuperPower
//              Computes and prints a Meta-Kosonic Super-Power Chain.
//
//
func SnapshooterPrinterMetaKosonicSuperPower(Chain1 []BalanceSuper, Chain2 []BalanceSuperLP) []MKSuperPowerPercent {
	fmt.Println("")
	fmt.Println("Computing Meta-Kosonic SUPER-Power Chain ...")
	Start1 := time.Now()
	MetaKosonicSuperPowerChain := CreateMetaKosonicSuperPowerChain(Chain1, Chain2)
	Elapsed1 := time.Since(Start1)
	fmt.Println("Done computing  Meta-Kosonic-SUPER-Power Chain, time required", Elapsed1)
	fmt.Println("===")

	fmt.Println("Computing Meta-Kosonic SUPER-Power-Percent Chain ...")
	Start2 := time.Now()
	MetaKosonicSuperPowerPercentChain := SuperPowerPercentComputer(MetaKosonicSuperPowerChain)
	Elapsed2 := time.Since(Start2)
	fmt.Println("Done computing  Meta-Kosonic SUPER-Power-Percent Chain, time required", Elapsed2)
	fmt.Println("===")

	fmt.Println("Sorting Meta-Kosonic SUPER-Power-Percent Chain ...")
	Start3 := time.Now()
	SortedMetaKosonicSuperPowerPercentChain := SortSuperPowerPercentChain(MetaKosonicSuperPowerPercentChain)
	Elapsed3 := time.Since(Start3)
	fmt.Println("Done sorting Meta-Kosonic SUPER-Power-Percent Chain, time required", Elapsed3)
	fmt.Println("===")

	fmt.Println("Outputting sorted Meta-Kosonic SUPER-Power-Percent-Chain to", N08)
	OutputFile3, err := os.Create(N08)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile3.Close()
	_, _ = fmt.Fprintln(OutputFile3, SortedMetaKosonicSuperPowerPercentChain)
	fmt.Println("Done Outputting sorted Meta-Kosonic SUPER-Power-Percent-Chain to", N08)
	fmt.Println("")

	return SortedMetaKosonicSuperPowerPercentChain
}

//
//
//======================================================================================================================
//
//
//[4]           SuperPowers
//              The Main Function for this Meta-Operation.
//
//
func SuperPowers() {
	SuperChain := SnapshooterPrinterSuper()
	SuperLPChain := SnapshooterPrinterSuperLP()

	SortedSuperPowerChain := SnapshooterPrinterSuperPower(SuperChain, SuperLPChain)
	SortedKosonicSuperPowerChain := SnapshooterPrinterKosonicSuperPower(SuperChain, SuperLPChain)
	SortedMetaKosonicSuperPowerChain := SnapshooterPrinterMetaKosonicSuperPower(SuperChain, SuperLPChain)

	fmt.Println("")
	fmt.Println("======RESULTS======")
	fmt.Println("There are only ", len(SortedSuperPowerChain), "addresses that have SuperPower")
	fmt.Println("There are only ", len(SortedKosonicSuperPowerChain), "addresses that have Kosonic SuperPower")
	fmt.Println("There are only ", len(SortedMetaKosonicSuperPowerChain), "addresses that have Meta-Kosonic SuperPower")
}
