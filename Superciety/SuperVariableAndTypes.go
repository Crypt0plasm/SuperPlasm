package Superciety

import p "github.com/Crypt0plasm/Firefly-APD"

//======================================================================================================================
//======================================================================================================================
//
//
// 		Superciety/SuperVariableAndTypes.go
//		Contains all the necessary Variables and Types to run the SuperPlasm code.
//
//
//              [A]Variable Definitions
//[A]1          TokenAccount Links
//[A]2          Exception Addresses
//[A]3          Token Identifiers
//[A]4          Pool Addresses
//
//              [B]Type Definitions
//[B]01         Simple types
//
//              [C]Complex Type Definitions
//[C]01         String Balance Type
//[C]02         Decimal Balance Type
//[C]03         Combined Type (percent containing types)
//[C]04         Super Price Structures
//[C]05         Custom JSON types, used for scanning Blockchain Json Data
//
//
//======================================================================================================================
//======================================================================================================================
//
//
//              [A]Variable Definitions
//
//
var (
	//
	//
	//[A]1      TokenAccount Links
	SUPER       = "https://api.elrond.com/tokens/SUPER-507aa6/accounts?size=10000"
	SuperEgldLP = "https://api.elrond.com/tokens/SUPEREGLD-a793b9/accounts?size=10000"
	SuperCamel  = "https://api.elrond.com/nfts/SCYMETA-3104d5-01/owners?size=10000"
	//
	//
	//[A]2      Exception Addresses
	ExA1 = ElrondAddress("erd1jd7gxdrv7qkghmm4afzk9hy6pw4qa5cfwt0nl7tmyhqujktc27rskzqmke") //Community Funds
	ExA2 = ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp") //Maiar Super-EGLD-LP Pool
	ExA3 = ElrondAddress("erd1qqqqqqqqqqqqqpgqawkm2tlyyz6vtg02fcr5w02dyejp8yrw0y8qlucnj2") //Jexchange Smart Contract
	//
	//
	//[A]3      Token Identifiers
	IdentifierSuper   = ESDTIdentifier("SUPER-507aa6")
	IdentifierWEGLD   = ESDTIdentifier("WEGLD-bd4d79")
	IdentifierUSDC    = ESDTIdentifier("USDC-c76f1f")
	IdentifierSuperLP = ESDTIdentifier("SUPEREGLD-a793b9")
	//
	//
	//[A]4      Pool Addresses
	SuperEgldPool = ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	EgldUSDCPool  = ElrondAddress("erd1qqqqqqqqqqqqqpgqeel2kumf0r8ffyhth7pqdujjat9nx0862jpsg2pqaq")
)

//======================================================================================================================
//======================================================================================================================
//
//
//              [B]Type Definitions
//
//
//[B]01         Simple types:
type ElrondAddress string
type ESDTIdentifier string

//
//
//======================================================================================================================
//======================================================================================================================
//
//
//              [C]Complex Type Definitions
//
//
//[C]01         String Balance Type:
//
type BalanceSuper struct {
	Address ElrondAddress `json:"address"`
	SuperB  string        `json:"balance"`
}
type BalanceSuperLP struct {
	Address  ElrondAddress `json:"address"`
	SuperLpB string        `json:"balance"`
}
type BalanceCamel struct {
	Address ElrondAddress `json:"address"`
	CamelB  string        `json:"balance"`
}

//======================================================================================================================
//
//
//[C]02         Decimal Balance Type:
//
//
type BalanceVLP struct { //Virtual LP Balance Structure
	Address ElrondAddress
	VLPB    *p.Decimal
}
type BalanceSFR struct { //Super Farm Reward Balance Structure
	Address ElrondAddress
	SFRB    *p.Decimal
}
type SuperPower struct { //Super-Power Value Structure
	Address ElrondAddress
	SPV     *p.Decimal
}
type MetaKosonicSuperPower struct { //MKSuperPower Type (used for Meta-Kosonic Super-Power)
	Address    ElrondAddress
	Super      *p.Decimal
	MetaSuper  *p.Decimal
	SuperPower *p.Decimal
}

//
//
//======================================================================================================================
//
//
//[C]03         Combined Type (percent containing types):
//
//
type SPPercent struct { //SPPercent Type (used for Super-Power and Kosonic Super-Power to display Percentages)
	Main              SuperPower
	SuperPowerPercent *p.Decimal
}
type MKSuperPowerPercent struct { //MKSuperPowerPercent Type (used for Meta-Kosonic Super-Power to display Percentages)
	Main                         MetaKosonicSuperPower
	MetaKosonicSuperPowerPercent *p.Decimal
}

//
//
//======================================================================================================================
//
//
//[C]04         Super Price Structures:
//
//
type MetaSuperPrice struct {
	Liquidity *p.Decimal
	SP        SuperPrices
}
type SuperPrices struct {
	DollarPool ESDTPoolAmounts
	SuperPool  ESDTPoolAmounts
	SV         SuperValue
	LPVC       SuperLPValueComposition
}
type ESDTPoolAmounts struct {
	EsdtAmount *p.Decimal
	EgldAmount *p.Decimal
}
type SuperValue struct {
	USDperEGLD   *p.Decimal
	SUPERperEGLD *p.Decimal
	USDperSUPER  *p.Decimal
}
type SuperLPValueComposition struct {
	SuperHalf  *p.Decimal
	ElrondHalf *p.Decimal
	TotalUSD   *p.Decimal
}

//
//
//======================================================================================================================
//
//
//[C]05         Custom JSON types, used for scanning Blockchain Json Data
//              Type is needed for unmarshalling. And getting the required field as data.
//
//
type USDCESDT []struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Ticker     string `json:"ticker"`
	Owner      string `json:"owner"`
	Minted     string `json:"minted"`
	Burnt      string `json:"burnt"`
	Decimals   int    `json:"decimals"`
	IsPaused   bool   `json:"isPaused"`
	Assets     struct {
		Description     string `json:"description"`
		LedgerSignature string `json:"ledgerSignature"`
		Status          string `json:"status"`
		PngURL          string `json:"pngUrl"`
		SvgURL          string `json:"svgUrl"`
	} `json:"assets"`
	Accounts       int    `json:"accounts"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

//
//
type WEGLDESDT []struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Ticker     string `json:"ticker"`
	Owner      string `json:"owner"`
	Minted     string `json:"minted"`
	Burnt      string `json:"burnt"`
	Decimals   int    `json:"decimals"`
	IsPaused   bool   `json:"isPaused"`
	Assets     struct {
		Website         string `json:"website"`
		Description     string `json:"description"`
		LedgerSignature string `json:"ledgerSignature"`
		Status          string `json:"status"`
		PngURL          string `json:"pngUrl"`
		SvgURL          string `json:"svgUrl"`
	} `json:"assets"`
	Transactions   int    `json:"transactions"`
	Accounts       int    `json:"accounts"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

//
//
type SuperESDT []struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Ticker     string `json:"ticker"`
	Owner      string `json:"owner"`
	Minted     string `json:"minted"`
	Burnt      string `json:"burnt"`
	Decimals   int    `json:"decimals"`
	IsPaused   bool   `json:"isPaused"`
	Assets     struct {
		Website     string `json:"website"`
		Description string `json:"description"`
		Social      struct {
			Email      string `json:"email"`
			Twitter    string `json:"twitter"`
			Whitepaper string `json:"whitepaper"`
			Coingecko  string `json:"coingecko"`
			Discord    string `json:"discord"`
			Telegram   string `json:"telegram"`
		} `json:"social"`
		Status string `json:"status"`
		PngURL string `json:"pngUrl"`
		SvgURL string `json:"svgUrl"`
	} `json:"assets"`
	Accounts       int    `json:"accounts"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

//
//
type SuperLpESDT []struct {
	Identifier     string `json:"identifier"`
	Name           string `json:"name"`
	Ticker         string `json:"ticker"`
	Owner          string `json:"owner"`
	Minted         string `json:"minted"`
	Burnt          string `json:"burnt"`
	Decimals       int    `json:"decimals"`
	IsPaused       bool   `json:"isPaused"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

//
//
type SuperLPSpecifications struct {
	Identifier        string `json:"identifier"`
	Name              string `json:"name"`
	Ticker            string `json:"ticker"`
	Owner             string `json:"owner"`
	Minted            string `json:"minted"`
	Burnt             string `json:"burnt"`
	Decimals          int    `json:"decimals"`
	IsPaused          bool   `json:"isPaused"`
	Transactions      int    `json:"transactions"`
	Accounts          int    `json:"accounts"`
	CanUpgrade        bool   `json:"canUpgrade"`
	CanMint           bool   `json:"canMint"`
	CanBurn           bool   `json:"canBurn"`
	CanChangeOwner    bool   `json:"canChangeOwner"`
	CanPause          bool   `json:"canPause"`
	CanFreeze         bool   `json:"canFreeze"`
	CanWipe           bool   `json:"canWipe"`
	Supply            string `json:"supply"`
	CirculatingSupply string `json:"circulatingSupply"`
}

//
//
//======================================================================================================================
//======================================================================================================================
