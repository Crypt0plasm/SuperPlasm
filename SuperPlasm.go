package main

import (
	ss "SuperPlasm/Superciety"
	"flag"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

func main() {
	var (
		CmpMKSPH = `--cmp-mksp=<ElrondAddress>;
Computes and Prints the Meta-Kosonic Super-Power of a given Elrond Address
The MKSP is a whole number
`
		OptimzeH = `--optimize=<ElrondAddress>;
Computes if an address is optimized for maximum MKSP; If not it display what
must be done to optimize its holding
`
		SuperHldH = `--get-super=<ElrondAddress>;
Displays the Super Holdings, LP and MKSP of a given Elrond Address
`
		SuperPriceH = `--get-super-price;
Display all relevant Super Prices.
`
		LiqRewPrgmH = `--print-lrp=<Daily-Reward-Amount>;
Prints Liquidity Reward Program rewards Chain.
`
		SuperPowerH = `--print-sp;
Prints Super-Power variant Chains.
`
	)

	const (
		CmpMKSP       = "cmp-mksp"        //String
		Optimize      = "optimize"        //String
		SuperHoldings = "get-super"       //String
		SuperPrice    = "get-super-price" //String
		LiqRewPrg     = "print-lrp"       //String
		SuperPower    = "print-sp"        //String
	)

	FlagCmpMKSP := flag.String(CmpMKSP, "0", CmpMKSPH)
	FlagOptimize := flag.String(Optimize, "0", OptimzeH)
	FlagShowSuper := flag.String(SuperHoldings, "0", SuperHldH)
	FlagSuperPrice := flag.Bool(SuperPrice, false, SuperPriceH)
	FlagLiqRewPrg := flag.String(LiqRewPrg, "0", LiqRewPrgmH)
	FlagSuperPower := flag.Bool(SuperPower, false, SuperPowerH)
	//
	flag.Parse()

	//1)First Flag outputs the Meta-Kosonic Super-Power of a given address
	if *FlagCmpMKSP != "0" {
		Addy := ss.ElrondAddress(*FlagCmpMKSP)
		MKSP := ss.GetAddyMKSP(Addy)
		fmt.Println(MKSP)
	}

	//2)Second Flag runs the MKSP Optimizer
	if *FlagOptimize != "0" {
		Addy := ss.ElrondAddress(*FlagOptimize)
		ss.Optimizer(Addy)
	}

	//3)Flag Displays Super Holdings
	if *FlagShowSuper != "0" {
		Addy := ss.ElrondAddress(*FlagShowSuper)
		ss.AddySpecsPrinter(Addy)
	}

	//4)Flag Displays Super Price
	if *FlagSuperPrice == true {
		ScannedSuperPrices := ss.GetAllSuperPrices()
		ss.PricePrinter(ScannedSuperPrices)
	}

	//5)Flag Prints Liquidity Reward Program Chains
	if *FlagLiqRewPrg != "0" {
		Reward := p.NFS(*FlagLiqRewPrg)
		ss.LiquidityRewardProgram(Reward)
	}

	//6)Flag Prints Super Power Chains
	if *FlagSuperPower == true {
		ss.SuperPowers()
	}

}
