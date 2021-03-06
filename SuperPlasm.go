package main

import (
	ss "SuperPlasm/Superciety"
	"flag"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

func main() {
	var (
		CmpMKSPH = `--cmp-mksp <ElrondAddress>;
Computes and Prints the Meta-Kosonic Super-Power of a given Elrond Address
The MKSP is a whole number
`
		OptimzeH = `--optimize <ElrondAddress>;
Computes if an address is optimized for maximum MKSP; If not it display what
must be done to optimize its holding
`
		SuperHldH = `--get-super <ElrondAddress>;
Displays the Super Holdings, LP and MKSP of a given Elrond Address
`
		SuperPriceH = `--get-super-price;
Display all relevant Super Prices.
`
		LiqRewPrgmH = `--print-lrp <Daily-Reward-Amount>;
Prints Liquidity Reward Program rewards Chain.
`
		SuperPowerH = `--print-sp;
Prints Super-Power variant Chains.
`
		BuySuperPowerH = `--buy-sp <ElrondAddress>;
Computes on what you must spend EGLD, given a certain address, to maximize MKSP.
`
		BuySuperH = `--buy-super <EGLD-Amount>;
Computes new Prices if Super is bought with the given EGLD Amount.
`
	)

	const (
		CmpMKSP       = "cmp-mksp"        //String
		Optimize      = "optimize"        //String
		SuperHoldings = "get-super"       //String
		SuperPrice    = "get-super-price" //String
		LiqRewPrg     = "print-lrp"       //String
		SuperPower    = "print-sp"        //String
		BuySuperPower = "buy-sp"          //String
		BuySuper      = "buy-super"       //String
	)

	FlagCmpMKSP := flag.String(CmpMKSP, "0", CmpMKSPH)
	FlagOptimize := flag.String(Optimize, "0", OptimzeH)
	FlagShowSuper := flag.String(SuperHoldings, "0", SuperHldH)
	FlagSuperPrice := flag.Bool(SuperPrice, false, SuperPriceH)
	FlagLiqRewPrg := flag.String(LiqRewPrg, "0", LiqRewPrgmH)
	FlagSuperPower := flag.Bool(SuperPower, false, SuperPowerH)
	FlagBuySuperPower := flag.String(BuySuperPower, "0", BuySuperPowerH)
	FlagBuySuper := flag.String(BuySuper, "0", BuySuperH)
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

	//7)Flag Prints Liquidity Reward Program Chains
	if *FlagBuySuperPower != "0" {
		var EGLD string
		Addy := ss.ElrondAddress(*FlagBuySuperPower)

		//Getting EGLD Amount
		fmt.Println("How Much EGLD you want to spend?")
		_, _ = fmt.Scan(&EGLD)
		Money := p.NFS(EGLD)

		ss.Acquisition(Addy, Money)
	}

	//8)Flag Prints Liquidity Reward Program Chains
	if *FlagBuySuper != "0" {
		Money := p.NFS(*FlagBuySuper)

		ss.TraderBuyer(Money)
	}
}
