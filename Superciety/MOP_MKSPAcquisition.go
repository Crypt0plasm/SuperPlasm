package Superciety

import (
	mt "SuperPlasm/SuperMath"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

func Acquisition(Addy ElrondAddress, Money *p.Decimal) {
	FirstPrice := GetAllSuperPrices()
	//BuyAmountDecimal := p.NFS(BuyAmount)

	SuperBought, VirtualPrice := BuySuper(Money, FirstPrice)
	MySuper, MySuperLp := GetAddySuperValues(Addy)
	SFT1Chain := CreateCamelChain()
	GetMeta := IzMeta(Addy, SFT1Chain)

	TotalVirtualSuper := mt.ADDxc(SuperBought, MySuper)

	fmt.Println("")
	fmt.Println("Scanning ADDRESS:", Addy)
	AddLpWith := MKSPCoreOptimizer(TotalVirtualSuper, MySuperLp, GetMeta, false, VirtualPrice)
	BuySuperWith := mt.SUBxc(Money, AddLpWith)
	fmt.Println("")
	if mt.DecimalEqual(AddLpWith, p.NFI(0)) == true {
		fmt.Println("==========================Results=======================================================")
		fmt.Println("STEP1,           use all:", KO(Money), "EGLD to buy $Uper")
		fmt.Println("STEP2,     run optimizer: with your address !")
		fmt.Println("======================END-Results=======================================================")
	} else {
		fmt.Println("==========================Results=======================================================")
		fmt.Println("STEP1,               use:", KO(BuySuperWith), "EGLD to buy $Uper")
		fmt.Println("STEP2,               use:", KO(AddLpWith), "EGLD to add $UPER-EGLD-LP")
		fmt.Println("======================END-Results=======================================================")
	}
}
