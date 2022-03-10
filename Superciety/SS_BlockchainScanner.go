package Superciety

import (
	mt "SuperPlasm/SuperMath"
	"encoding/json"
	p "github.com/Crypt0plasm/Firefly-APD"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/SuperScanner/SS_BlockchainScanner.go
//		Functions that scan the Elrond Blockchain for Data.
//
//
//		[A]Pure Blockchain Scanner Functions
//[A]01         GetSuperLpCirculatingSupply         Scans the Blockchain for the SuperLP circulating Amount.
//[A]02         SuperLPTotalSupply                  Computes the correct Super-LP circulating Amount.
//[A]03         GetAllSuperPrices                   Scans the Blockchain for all relevant Super price information.
//
//              [B]Elrond Address based Blockchain Scanning Functions
//[B]01         ReadSuperBalance                    Returns the Super Balance of a given Elrond Address.
//[B]02         ReadSuperLPBalance                  Returns the Super-EGLD-LP Balance of a given Elrond Address.
//[B]03         ReadWEGLDBalance                    Returns the wrappedEGLD Balance of a given Elrond Address.
//[B]04         ReadUSDCBalance                     Returns the USDC Balance of a given Elrond Address.
//[B]05         GetAddySuperValues                  Returns Super and SuperLP Value of a given Elrond Address.
//[B]06         AddySpecs                           Returns Super, SuperLP and MKSP of a given Elrond Address.
//
//              [C]LookUp Functions, look up an address inside a given chain.
//[C]01         GetSuperAmount                      Looks up the Super Amount for an Elrond Address in a Super Chain.
//[C]02         GetSuperLPAmount                    Looks up the SuperLP Amount for an Elrond Address in a SuperLP Chain.
//[C]03         GetCamelAmount                      Looks up the Camel Amount for an Elrond Address in a Camel Chain.
//[C]04         IzMeta                              Checks if an address is Superciety Meta.
//
//======================================================================================================================
//======================================================================================================================
//
//
//              [A]Pure Blockchain Scanner Functions
//
//
//[A]01         GetSuperLpCirculatingSupply
//              Scans the Blockchain for the SuperLP circulating Amount.
//              Deprecated function, as it was observed that the Elrond-Explorer listed amount
//              Doest correspond to the Sum of all Super-Egld-LP Values
//              For the actual value, use SuperLPTotalSupply
//
//
func GetSuperLpCirculatingSupply() string {
	var (
		String1     = "https://api.elrond.com/tokens/"
		ScannedJSON SuperLPSpecifications
	)
	ScanURL := String1 + string(IdentifierSuperLP)
	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	CS := ScannedJSON.CirculatingSupply
	return CS
}

//
//
//======================================================================================================================
//
//
//[A]02         SuperLPTotalSupply
//              Scans all the Super-Egld-LP amounts and computes their Sum
//              Effectively returning the Super-Egld-LP circulating Supply
//
//
func SuperLPTotalSupply() *p.Decimal {
	LPs := CreateSuperLPChain()
	var Sum = new(p.Decimal)
	for i := 0; i < len(LPs); i++ {
		Balance := ConvertAU18(LPs[i].SuperLpB)
		Sum = mt.ADDxc(Sum, Balance)
	}
	return Sum
}

//
//
//======================================================================================================================
//
//
//[A]03         GetAllSuperPrices
//              Scans the Blockchain for the current EGLD-USDC and EGLD-Super Pool ESDT quantities
//              Effectively returning all relevant Super price information
//
//
func GetAllSuperPrices() (Result MetaSuperPrice) {
	var (
		DollarPool, SuperPool ESDTPoolAmounts
		SV                    SuperValue
		LPVC                  SuperLPValueComposition
	)
	//fmt.Println("=============SUPER==================")
	USDCPoolUSDCAmount := ConvertAU06(ReadUSDCBalance(EgldUSDCPool))
	USDCPoolEGLDAmount := ConvertAU18(ReadWEGLDBalance(EgldUSDCPool))
	USDperEGLD := mt.TruncateCustom(mt.DIVxc(USDCPoolUSDCAmount, USDCPoolEGLDAmount), 18)
	DollarPool = ESDTPoolAmounts{EsdtAmount: USDCPoolUSDCAmount, EgldAmount: USDCPoolEGLDAmount}

	SuperPoolSuperAmount := ConvertAU18(ReadSuperBalance(SuperEgldPool))
	SuperPoolEGLDAmount := ConvertAU18(ReadWEGLDBalance(SuperEgldPool))
	SuperperEGLD := mt.TruncateCustom(mt.DIVxc(SuperPoolSuperAmount, SuperPoolEGLDAmount), 18)
	SuperPool = ESDTPoolAmounts{EsdtAmount: SuperPoolSuperAmount, EgldAmount: SuperPoolEGLDAmount}

	USDperSUPER := mt.TruncateCustom(mt.DIVxc(USDperEGLD, SuperperEGLD), 18)
	SV = SuperValue{USDperEGLD: USDperEGLD, SUPERperEGLD: SuperperEGLD, USDperSUPER: USDperSUPER}

	//fmt.Println("===========SUPER-LP==================")
	Liquidity := SuperLPTotalSupply()
	SuperLP1 := mt.TruncateCustom(mt.DIVxc(SuperPoolSuperAmount, Liquidity), 18)
	SuperLP2 := mt.TruncateCustom(mt.DIVxc(SuperPoolEGLDAmount, Liquidity), 18)
	SuperLPValue := mt.TruncateCustom(mt.ADDxc(mt.MULxc(SuperLP1, USDperSUPER), mt.MULxc(SuperLP2, USDperEGLD)), 18)
	LPVC = SuperLPValueComposition{SuperHalf: SuperLP1, ElrondHalf: SuperLP2, TotalUSD: SuperLPValue}

	Intermediary := SuperPrices{DollarPool: DollarPool, SuperPool: SuperPool, SV: SV, LPVC: LPVC}
	Result = MetaSuperPrice{Liquidity: Liquidity, SP: Intermediary}
	return Result
}

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//              [B]Elrond Address based Blockchain Scanning Functions
//
//
//[B]01         ReadSuperBalance
//              Returns the Super Balance of a given Elrond Address.
//
//
func ReadSuperBalance(Addy ElrondAddress) string {
	var (
		String1     = "https://api.elrond.com/accounts/"
		String2     = "/tokens?identifier="
		String3     = "&identifiers="
		ScannedJSON SuperESDT
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(IdentifierSuper) + String3
	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON[0].Balance
	}

	return Balance
}

//======================================================================================================================
//
//
//[B]02         ReadSuperLPBalance
//              Returns the Super-EGLD-LP Balance of a given Elrond Address.
//
//
func ReadSuperLPBalance(Addy ElrondAddress) string {
	var (
		String1     = "https://api.elrond.com/accounts/"
		String2     = "/tokens?identifier="
		String3     = "&identifiers="
		ScannedJSON SuperLpESDT
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(IdentifierSuperLP) + String3
	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON[0].Balance
	}
	return Balance
}

//======================================================================================================================
//
//
//[B]03         ReadWEGLDBalance
//              Returns the wrappedEGLD Balance of a given Elrond Address.
//
//
func ReadWEGLDBalance(Addy ElrondAddress) string {
	var (
		String1     = "https://api.elrond.com/accounts/"
		String2     = "/tokens?identifier="
		String3     = "&identifiers="
		ScannedJSON WEGLDESDT
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(IdentifierWEGLD) + String3
	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON[0].Balance
	}
	return Balance
}

//======================================================================================================================
//
//
//[B]04         ReadUSDCBalance
//              Returns the USDC Balance of a given Elrond Address.
//
//
func ReadUSDCBalance(Addy ElrondAddress) string {
	var (
		String1     = "https://api.elrond.com/accounts/"
		String2     = "/tokens?identifier="
		String3     = "&identifiers="
		ScannedJSON USDCESDT
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(IdentifierUSDC) + String3
	Snapshot := OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)

	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON[0].Balance
	}
	return Balance
}

//======================================================================================================================
//
//
//[B]05         GetAddySuperValues
//              Returns Super and SuperLP Value of a given Elrond Address.
//
//
func GetAddySuperValues(Addy ElrondAddress) (AddressSuperAmount, AddressSuperLpAmount *p.Decimal) {
	AddressSuperAmount = ConvertAU18(ReadSuperBalance(Addy))
	AddressSuperLpAmount = ConvertAU18(ReadSuperLPBalance(Addy))
	return AddressSuperAmount, AddressSuperLpAmount
}

//======================================================================================================================
//
//
//[B]06         AddySpecs
//              Returns Super, SuperLP and MKSP of a given Elrond Address.
//
//
func AddySpecs(Addy ElrondAddress) (Super, SuperLP, SP *p.Decimal) {
	Super, SuperLP = GetAddySuperValues(Addy)
	SP = GetAddyMKSP(Addy)

	return
}

//======================================================================================================================
//======================================================================================================================
//
//
//              [C]LookUp Functions, look up an address inside a given chain.
//
//
//[C]01         GetSuperAmount
//              Looks up the Super Amount for an Elrond Address in a Super Chain.
//
//
func GetSuperAmount(Address ElrondAddress, Chain []BalanceSuper) *p.Decimal {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].SuperB
			break
		} else {
			Result = "0"
		}
	}

	//Converting ReadString to Decimal
	EndResult := ConvertAU18(Result)
	return EndResult
}

//
//
//======================================================================================================================
//
//
//[C]02         GetSuperLPAmount
//              Looks up the SUPER-EGLD-LP Amount for an Elrond Address in a SuperLP Chain.
//
//
func GetSuperLPAmount(Address ElrondAddress, Chain []BalanceSuperLP) *p.Decimal {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].SuperLpB
			break
		} else {
			Result = "0"
		}
	}
	//Converting ReadString to Decimal
	EndResult := ConvertAU18(Result)
	return EndResult
}

//
//
//======================================================================================================================
//
//
//[C]03         GetCamelAmount
//              Looks up the Camel Amount for an Elrond Address in a Camel Chain.
//
//
func GetCamelAmount(Address ElrondAddress, Chain []BalanceCamel) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Address {
			Result = Chain[i].CamelB
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

//
//
//======================================================================================================================
//
//
//[C]04         IzMeta
//              Checks if an address is Superciety Meta.
//              Checking is done by checking 1 or multiple SFTs/NFTs
//
//
func IzMeta(Addy ElrondAddress, Chain1 []BalanceCamel) bool {
	var (
		MetaResult bool //Total boolean Value
		IzCamel    bool //1st boolean value to check
	)

	//1st SFT Check
	CamelValue := GetCamelAmount(Addy, Chain1)
	if mt.DecimalGreaterThanOrEqual(p.NFS(CamelValue), p.NFS("1")) == true {
		IzCamel = true
	} else {
		IzCamel = false
	}

	//If all SFT Checks are true, IzMeta is true.
	if IzCamel == true {
		MetaResult = true
	}
	return MetaResult
}

//
//
//======================================================================================================================
//======================================================================================================================
