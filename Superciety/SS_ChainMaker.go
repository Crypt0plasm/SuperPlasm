package Superciety

import (
	"encoding/json"
)

//======================================================================================================================
//======================================================================================================================
//
//
// 	        SuperScanner/SS_ChainMaker.go
//	        Creates Chains with Blockchain Scanned Values.
//
//
//              [A]Blockchain Scanner Chain Makers.
//[A]01         CreateSuperChain            Super Snapshooter Function; Creates a Chain of Super Values.
//[A]02         CreateSuperLPChain          Super-EGLD-LP Snapshooter Function; Creates a Chain of Super-EGLD-Values.
//[A]03         CreateCamelChain            Camel Snapshooter Function; Creates a Chain of Camel Values.
//
//======================================================================================================================
//======================================================================================================================
//
//              [A]Blockchain Scanner Chain Makers.
//
//
//[A]01         CreateSuperChain
//              Super Snapshooter Function; Creates a Chain of Super Values.
//
//
func CreateSuperChain() []BalanceSuper {
	var OutputChain []BalanceSuper
	SS := OnPage(SUPER)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//
//
//======================================================================================================================
//
//[A]02         CreateSuperLPChain
//              Super-EGLD-LP Snapshooter Function; Creates a Chain of Super-EGLD-Values.
//
func CreateSuperLPChain() []BalanceSuperLP {
	var OutputChain []BalanceSuperLP
	SS := OnPage(SuperEgldLP)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//
//
//======================================================================================================================
//
//
//
//[A]03         CreateCamelChain
//              Camel Snapshooter Function; Creates a Chain of Camel Values.
//
//
func CreateCamelChain() []BalanceCamel {
	var OutputChain []BalanceCamel
	SS := OnPage(SuperCamel)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

//
//
//======================================================================================================================
//======================================================================================================================
