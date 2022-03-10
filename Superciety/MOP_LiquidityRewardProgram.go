package Superciety

import (
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"log"
	"os"
	"time"
)

var (
	N01 = "Snapshot_SUPER.txt"
	N02 = "Snapshot_SUPER-EGLD-LP.txt"
	N03 = "Snapshot_SUPER-Camel.txt"
	N04 = "Computed_SUPER-VLP.txt"
	N05 = "Computed_SUPER-VLP-Rewards.txt"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/..MetaOperations../MOP_LiquidityRewardProgram.go
//		Functions that are used for the Liquidity-Reward-Program Meta-Operation.
//
//
//[1]           SnapshooterPrinterSuper                 Snapshots and prints a Super Chain.
//[2]           SnapshooterPrinterSuperLP               Snapshots and prints a SuperLP Chain.
//[3]           SnapshooterPrinterSuperCamel            Snapshots and prints a Super Camel Chain.
//[4]           SnapshooterPrinterSuperVirtualLP        Computes and prints a Super VirtualLP Chain.
//[5]           SnapshooterPrinterSuperVirtualLPRewards Computes and prints a Super Liquidity-Reward-Program Chain.
//[6]           LiquidityRewardProgram                  The Main Function for this Meta-Operation.
//
//======================================================================================================================
//======================================================================================================================
//
//
//[1]           SnapshooterPrinterSuper
//              Snapshots and prints a Super Chain.
//
//
func SnapshooterPrinterSuper() []BalanceSuper {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER Amounts ...")
	Start := time.Now()
	SuperChain := CreateSuperChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER Amounts to", N01)
	OutputFile, err := os.Create(N01)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, SuperChain)
	fmt.Println("DONE Outputting SUPER Amounts to", N01)
	fmt.Println("")

	return SuperChain
}

//
//
//======================================================================================================================
//
//
//[2]           SnapshooterPrinterSuperLP
//              Snapshots and prints a SuperLP Chain.
//
//
func SnapshooterPrinterSuperLP() []BalanceSuperLP {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER-LP Amounts ...")
	Start := time.Now()
	LPChain := CreateSuperLPChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER-LP Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-LP Amounts to", N02)
	OutputFile, err := os.Create(N02)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, LPChain)
	fmt.Println("DONE Outputting SUPER Amounts to", N02)
	fmt.Println("")

	return LPChain
}

//
//
//======================================================================================================================
//
//
//[3]           SnapshooterPrinterSuperCamel
//              Snapshots and prints a Super Camel Chain.
//
//
func SnapshooterPrinterSuperCamel() []BalanceCamel {
	fmt.Println("")
	fmt.Println("Snapshotting SUPER-Camel Amounts ...")
	Start := time.Now()
	CamelChain := CreateCamelChain()
	Elapsed := time.Since(Start)
	fmt.Println("Done snapshotting SUPER-Camel Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-Camel Amounts to", N03)
	OutputFile, err := os.Create(N03)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, CamelChain)
	fmt.Println("DONE Outputting SUPER Amounts to", N03)
	fmt.Println("")

	return CamelChain
}

//
//
//======================================================================================================================
//
//
//[4]           SnapshooterPrinterSuperVirtualLP
//              Computes and prints a Super VirtualLP Chain.
//
//
func SnapshooterPrinterSuperVirtualLP(Chain1 []BalanceSuperLP, Chain2 []BalanceCamel) []BalanceVLP {
	fmt.Println("")
	fmt.Println("Computing SUPER-VLP Amounts ...")
	Start := time.Now()
	VLPChain := CreateVLPChain(Chain1, Chain2)
	Elapsed := time.Since(Start)
	fmt.Println("Done computing SUPER-VLP Amounts, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-VLP Amounts to", N04)
	OutputFile, err := os.Create(N04)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, VLPChain)
	fmt.Println("DONE Outputting SUPER-VLP Amounts to", N04)
	fmt.Println("")

	return VLPChain
}

//
//
//======================================================================================================================
//
//
//[5]           SnapshooterPrinterSuperVirtualLPRewards
//              Computes and prints a Super Liquidity-Reward-Program Chain.
//
//
func SnapshooterPrinterSuperVirtualLPRewards(Chain1 []BalanceVLP, Reward *p.Decimal) []BalanceSFR {
	fmt.Println("")
	fmt.Println("Computing SUPER-VLP Rewards considering ", Reward, " per day")
	Start := time.Now()
	RewardChain := SuperVLPRewardComputerChain(Chain1, Reward)
	Elapsed := time.Since(Start)
	fmt.Println("Done computing SUPER-VLP Rewards, time required:", Elapsed)

	//Printing Snapshot
	fmt.Println("Outputting SUPER-VLP-Rewards to", N05)
	OutputFile, err := os.Create(N05)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, RewardChain)
	fmt.Println("DONE Outputting SUPER-VLP Amounts to", N05)
	fmt.Println("")

	return RewardChain
}

//
//
//======================================================================================================================
//
//
//[6]           LiquidityRewardProgram
//              The Main Function for this Meta-Operation.
//
//
func LiquidityRewardProgram(Reward *p.Decimal) {
	//SuperChain := SnapshooterPrinterSuper()
	SuperLPChain := SnapshooterPrinterSuperLP()
	CamelChain := SnapshooterPrinterSuperCamel()
	VLPChain := SnapshooterPrinterSuperVirtualLP(SuperLPChain, CamelChain)
	SnapshooterPrinterSuperVirtualLPRewards(VLPChain, Reward)

	fmt.Println("")
	fmt.Println("======RESULTS======")
	fmt.Println("There are ", len(SuperLPChain), "addresses that have LP")
	fmt.Println("There are ", len(CamelChain), "addresses that have Camels")
	fmt.Println("There are only ", len(VLPChain), "addresses that are eligible for LP Rewards")
}
