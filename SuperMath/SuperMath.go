package SuperMath

import (
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"os"
	"strconv"
)

const (
	SuperPlasmMaxMathPrecision  = uint32(150) //Total and Decimal Precision
	SuperPlasmStdMathPrecision  = uint32(50)  //Total and Decimal Precision
	SuperPlasmXPPrecision       = uint32(8)
	SuperPlasmCurrencyPrecision = uint32(18)
	SuperPlasmPromillePrecision = uint32(6)
	AuPerUnit                   = "1000000000000000000" // Atomic Units per Cryptoplasm
)

var (
	LOCPrecisionContext = p.Context{
		Precision:   SuperPlasmStdMathPrecision,
		MaxExponent: p.MaxExponent,
		MinExponent: p.MinExponent,
		Rounding:    p.RoundDown,
		// Default error conditions.
		Traps: p.InvalidOperation,
	}

	c   = LOCPrecisionContext
	AUs = p.NFS(AuPerUnit)
)

//
//	SuperMath.go					Precision Math Specific Functions
//		Copied after:
//		BlockChain_F.Firefly.go			Cryptoplasm(Koson) Precision Math Specific Functions
//================================================================================================
//************************************************************************************************
//================================================================================================
// 	Function List:
//
//	01 Comparison Functions operating on decimal type
//		00  - SummedMaxLengthPlusOne		SummedMaxLength returns the sum of the maximums length of digits (b4 and after coma)
//		00a - MaxInt32				Returns the maximum between two int32 numbers
//		00b - MaxInt64				Returns the maximum between two int64 numbers
//		00c - MaxDecimal			Returns the maximum between two decimals
//		01  - DecimalEqual			x == y
//		02  - DecimalNotEqual			x != y
//		03  - DecimalLessThan			x < y
//		04  - DecimalLessThanOrEqual		x <= y
//		05  - DecimalGreaterThan		x > y
//		06  - DecimalGreaterThanOrEqual		x >= y
//	02 Addition Functions
//		01  - ADDx				Adds 2 numbers with custom total precision
//		02  - ADDs				Adds 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - ADD				Adds 2 numbers with custom decimal precision and elastic integer precision
//		03a - ADDxs				Adds 2 numbers with 70 decimal precision and elastic integer precision
//		03b - ADDxc				Adds 2 numbers with 100 decimal precision and elastic integer precision
//		04  - SUMx				Adds multiple numbers with custom total precision
//		05  - SUMs				Adds multiple numbers within CryptoplasmPrecisionContext (70 total precision)
//		06  - SUM				Adds multiple numbers with custom decimal precision and elastic integer precision
//		06a - SUMxs				Adds multiple numbers with 70 decimal precision and elastic integer precision
//		06b - SUMxc				Adds multiple numbers with 100 decimal precision and elastic integer precision
//	03 Subtraction Functions
//		01  - SUBx				Subtracts 2 numbers with custom total precision
//		02  - SUBs				Subtracts 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - SUB				Subtracts 2 numbers with custom decimal precision and elastic integer precision
//		03a - SUBxs				Subtracts 2 numbers with 70 decimal precision and elastic integer precision
//		03b - SUBxc				Subtracts 2 numbers with 100 decimal precision and elastic integer precision
//		04  - DIFx				Subtracts multiple numbers with custom total precision
//		05  - DIFs				Subtracts multiple numbers within CryptoplasmPrecisionContext (70 total precision)
//		06  - DIF				Subtracts multiple numbers with custom decimal precision and elastic integer precision
//		06a - DIFxs				Subtracts multiple numbers with 70 decimal precision and elastic integer precision
//		06b - DIFxc				Subtracts multiple numbers with 100 decimal precision and elastic integer precision
//	04 Multiplication Functions
//		01  - MULx				Multiplies 2 numbers with custom total precision
//		02  - MULs				Multiplies 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - MULxc				Multiplies 2 numbers with elastic integer precision and 100 max decimal precision
//		04  - PRDx				Multiplies multiple numbers within a specific precision context
//		05  - PRDs				Multiplies multiple numbers within CryptoplasmPrecisionContext
//		06  - PRDxc				Multiplies multiple numbers with elastic integer precision and 100 max decimal precision
//		07  - POWx				Computes x ** y within a specific precision context
//		08  - POWs				Computes x ** y within CryptoplasmPrecisionContext
//		09  - POWxcs				Computes x ** y with elastic integer precision and custom max decimal precision
//		10  - POWxc				Computes x ** y with elastic integer precision and 150 max decimal precision
//		11  - Logarithm				Computes the logarithm from "number" in base "base".
//	05 Division Functions
//		01  - DIVx				Divides 2 numbers within a specific precision context
//		02  - DIVs				Divides 2 numbers within CryptoplasmPrecisionContext
//		03  - DIVxc				Divides 2 numbers with elastic integer precision and 100/101 max decimal precision
//		04  - DivInt				Returns x // y, uses elastic Precision (result is "integer")
//		05  - DivMod				Returns x % y, uses elastic Precision (result is the rest)
//  05a Mean Functions
//		01  - TwoMean				Returns the mean of two decimals
//	06 Truncate Functions
//		01  - TruncateCustom			Truncates using custom Precision (it must be know beforehand)
//		02  - TruncSeed				Truncates elastically to CryptoplasmSeedPrecision
//		03  - TruncToCurrency			Truncates elastically to CryptoplasmCurrencyPrecision
//		04  - TruncPercent			Truncates elastically to CryptoplasmPercentPrecision
//	07 List Functions
//		01  - SumDL				Adds all the decimals in a slice of decimals
//		02  - LastDE				Returns the last element in a slice
//		03  - AppDec				Unites 2 slices made of decimals
//		04  - Reverse				Reverses a slice of decimals
//		05  - PrintDL				Prints the "decimals" from a slice of strings
//		06  - WriteList				Writes strings from a slice to an external file
//	08 Digit Manipulation Functions
//		01  - RemoveDecimals			Removes the decimals of a number, uses floor function
//		02  - Count4Coma			Counts the number of digits before precision
//	09 Koson Amount String Manipulation Function
//		01  - Convert2AU			Converts Koson Amount to AtomicUnits (AttoPlasms)
//		02  - AttoPlasm2String			Converts AttoPlasms into a slice of strings
//		03  - KosonicDecimalConversion		Converts a Koson Amount into a string that can be better used for display purposes
//
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 01.00 - SummedMaxLength
//
// SummedMaxLength returns the sum of the maximums length of digits.
// before and after the coma for two decimals.
// Used in comparison operation, and to elastically increase precision based on integer part size of the decimals
// Even thought it on itself is enough to secure total operation precision, it is used as extra buffer when computing
// total operation precision for ADD SUB MUL and DIV functions. (because an additional DecimalPrecision is added on top of it)
func SummedMaxLengthPlusOne(x, y *p.Decimal) uint32 {
	var SML uint32
	IntegerDigitsMember1 := Count4Coma(x)  //int64
	IntegerDigitsMember2 := Count4Coma(y)  //int64
	DecimalDigitsMember1 := 0 - x.Exponent //int32
	DecimalDigitsMember2 := 0 - y.Exponent //int32

	MaxIntegerDigitsInt64 := MaxInt64(IntegerDigitsMember1, IntegerDigitsMember2) //int64
	MaxDecimalDigitsInt32 := MaxInt32(DecimalDigitsMember1, DecimalDigitsMember2) //int32

	// Observation1
	// Converting Int64 to Int32, going down from 9.223.372.036.854.775.807 to 2.147.483.647
	// As 2.147.483.647 are already a huge number, no check is implemented here.
	MaxIntegerDigitsInt32 := int32(MaxIntegerDigitsInt64)

	// Observation 2
	// SML is of uint32 type, this means this function works reliably for numbers with up to
	// 2.147.483.647 digits each.
	// Two times this is 4.294.967.294. Adding another 1 equals 4.294.967.295.
	// This is the maximum number representable by uint32.
	//
	// So the maximum length x and y can have is 2.147.483.647 digits before and after the coma.
	SML = uint32(MaxDecimalDigitsInt32) + uint32(MaxIntegerDigitsInt32) + 1
	return SML
}

//================================================================================================
//	01 Comparison Functions between integers:
//================================================
//
// Function 01.00a - MaxInt32
//
// MaxInt32 returns the maximum between two int32 numbers
func MaxInt32(x, y int32) int32 {
	var max int32
	digdiff := x - y
	if digdiff <= 0 {
		max = y
	} else if digdiff > 0 {
		max = x
	}
	return max
}

//
//================================================
//
// Function 01.00b - MaxInt64
//
// MaxInt64 returns the maximum between two int64 numbers
func MaxInt64(x, y int64) int64 {
	var max int64
	digdiff := x - y
	if digdiff <= 0 {
		max = y
	} else if digdiff > 0 {
		max = x
	}
	return max
}

//
//================================================
//
// Function 01.00c - MaxDecimal
//
// MaxDecimal returns the maximum between two Decimals
func MaxDecimal(x, y *p.Decimal) (max *p.Decimal) {
	Difference := SUBxc(x, y)
	if DecimalLessThanOrEqual(Difference, p.NFI(0)) == true {
		max = y
	} else if DecimalGreaterThan(Difference, p.NFI(0)) == true {
		max = x
	}
	return max
}

//
//================================================================================================
//	01 Comparison Functions between decimals:
//	The functions use the SummedMaxLengthPlusOne function to set the ComparisonContextPrecision
//================================================================================================
//
// Function 01.01 - DecimalEqual
//
// DecimalEqual returns true if decimal x is equal to decimal y.
func DecimalEqual(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, x, y)
	IsThreshold := Difference.IsZero()

	if IsThreshold == true {
		Result = true
	} else {
		Result = false
	}

	return Result
}

//
//================================================
//
// Function 01.02 - DecimalNotEqual
//
// DecimalNotEqual returns true if decimal x is not equal to decimal y.
// Only works with valid Decimal type numbers.
func DecimalNotEqual(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, x, y)
	IsThreshold := Difference.IsZero()

	if IsThreshold == true {
		Result = false
	} else {
		Result = true
	}

	return Result
}

//
//================================================
//
// Function 01.03 - DecimalLessThan
//
// DecimalLessThan returns true if decimal x is less than decimal y.
// Only works with valid Decimal type numbers.
// x equals y would return false as in this case x isnt less than y
func DecimalLessThan(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, x, y)
	//IsThreshold := Difference.IsZero()

	if Difference.Negative == true {
		Result = true
	} else {
		Result = false
	}

	return Result
}

//
//================================================
//
// Function 01.04 - DecimalLessThanOrEqual
//
// DecimalLessThanOrEqual returns true if decimal either
// x is less than decimal y, or if they are equal.
// Only works with valid Decimal type numbers.
func DecimalLessThanOrEqual(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, x, y)
	IsThreshold := Difference.IsZero()

	if Difference.Negative == true || IsThreshold == true {
		Result = true
	} else {
		Result = false
	}

	return Result
}

//
//================================================
//
// Function 01.05 - DecimalGreaterThan
//
// DecimalGreaterThan returns true if decimal x is greater than decimal y.
// Only works with valid Decimal type numbers.
// x equals y would return false as in this case x isn't less than y
func DecimalGreaterThan(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, y, x)
	//IsThreshold := Difference.IsZero()

	if Difference.Negative == true {
		Result = true
	} else {
		Result = false
	}

	return Result
}

//
//================================================
//
// Function 01.06 - DecimalGreaterThanOrEqual
//
// DecimalGreaterThanOrEqual returns true if decimal either
// x is greater than decimal y, or if they are equal.
// Only works with valid Decimal type numbers.
func DecimalGreaterThanOrEqual(x, y *p.Decimal) bool {
	var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x, y)

	Difference := SUBx(ComparisonContextPrecision, y, x)
	IsThreshold := Difference.IsZero()

	if Difference.Negative == true || IsThreshold == true {
		Result = true
	} else {
		Result = false
	}

	return Result
}

//
//================================================================================================
//	02,03,04,05 Mathematical operator Functions:
// 		Mathematical operating functions used for computing
//		Addition, Subtraction, Div, Multiplication, etc
//		Basic Operations done under CryptoplasmPrecisionContext without
//		Condition and error reporting as supported by p
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 02.01 - ADDx
//
// ADDx adds two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func ADDx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}

//
//================================================
//
// Function 02.02 - ADDs
//
// ADDs adds two decimals within CryptoplasmPrecisionContext Context
func ADDs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Add(result, member1, member2)
	return result
}

//
//================================================
//
// Function 02.03 - ADD
//
// ADD adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func ADD(DecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DNBDP := SummedMaxLengthPlusOne(member1, member2) //DigitNumberBasedDecimalPrecision
	//Observation
	// As "SummedMaxLengthPlusOne" returns a uint32 variable (maximum of 4.294.967.295)
	// TotalDecimalPrecision will overflow uint32 if adding the "DecimalPrecision" on top of DNBDP because
	// it (TotalDecimalPrecision) would get bigger than 4.294.967.295.
	// However, this isn't expected to happen, which is why no check or error detection is implemented.
	TotalDecimalPrecision := DNBDP + DecimalPrecision

	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}

//
//================================================
//
// Function 02.03a - ADDxs
//
// ADDxs adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 70 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func ADDxs(member1, member2 *p.Decimal) *p.Decimal {
	return ADD(SuperPlasmStdMathPrecision, member1, member2)
}

//
//================================================
//
// Function 02.03b - ADDxc
//
// ADDxc adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func ADDxc(member1, member2 *p.Decimal) *p.Decimal {
	return ADD(SuperPlasmMaxMathPrecision, member1, member2)
}

//
//================================================
//
// Function 02.04 - SUMx
//
// SUMx adds multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUMx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Add(restsum, restsum, item)
	}
	_, _ = cc.Add(sum, first, restsum)
	return sum
}

//
//================================================
//
// Function 02.05 - SUMs
//
// SUMs adds multiple decimals within CryptoplasmPrecisionContext Context
func SUMs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)

	for _, item := range rest {
		_, _ = c.Add(restsum, restsum, item)
	}
	_, _ = c.Add(sum, first, restsum)
	return sum
}

//
//================================================
//
// Function 02.06 - SUM
//
// SUM sums multiple decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func SUM(DecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	for _, item := range rest {
		restsum = ADD(DecimalPrecision, restsum, item)
	}
	sum = ADD(DecimalPrecision, first, restsum)
	return sum
}

//
//================================================
//
// Function 02.06a - SUMxs
//
// SUMxs adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 50 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func SUMxs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return SUM(SuperPlasmStdMathPrecision, first, rest...)
}

//
//================================================
//
// Function 02.06b - SUMxc
//
// SUMxc adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func SUMxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return SUM(SuperPlasmMaxMathPrecision, first, rest...)
}

//
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 03.01 - SUBx
//
// SUBx subtract two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUBx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}

//================================================
//
// Function 03.02 - SUBs
//
// SUBs subtract two decimals within CryptoplasmPrecisionContext Context
func SUBs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Sub(result, member1, member2)
	return result
}

//
//================================================
//
// Function 03.03 - SUB
//
// SUB subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.

func SUB(DecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DNBDP := SummedMaxLengthPlusOne(member1, member2) //DigitNumberBasedDecimalPrecision
	//Observation
	// As "SummedMaxLengthPlusOne" returns a uint32 variable (maximum of 4.294.967.295)
	// TotalDecimalPrecision will overflow uint32 if adding the "DecimalPrecision" on top of DNBDP because
	// it (TotalDecimalPrecision) would get bigger than 4.294.967.295.
	// However, this isn't expected to happen, which is why no check or error detection is implemented.
	TotalDecimalPrecision := DNBDP + DecimalPrecision

	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}

//
//================================================
//
// Function 03.03a - SUBxs
//
// SUBxs subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 50 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func SUBxs(member1, member2 *p.Decimal) *p.Decimal {
	return SUB(SuperPlasmStdMathPrecision, member1, member2)
}

//
//================================================
//
// Function 03.03b - SUBxc
//
// SUBxc subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 150 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func SUBxc(member1, member2 *p.Decimal) *p.Decimal {
	return SUB(SuperPlasmMaxMathPrecision, member1, member2)
}

//
//================================================
//
// Function 03.04 - DIFx
//
// DIFx subtracts multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIFx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Add(restsum, restsum, item)
	}
	_, _ = cc.Sub(sum, first, restsum)
	return sum
}

//
//================================================
//
// Function 03.05 - DIFs
//
// DIFs subtracts multiple decimals within CryptoplasmPrecisionContext Context
func DIFs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)

	for _, item := range rest {
		_, _ = c.Add(restsum, restsum, item)
	}
	_, _ = c.Sub(sum, first, restsum)
	return sum
}

//
//================================================
//
// Function 03.06 - DIF
//
// DIF subtracts multiple decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func DIF(DecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	for _, item := range rest {
		restsum = ADD(DecimalPrecision, restsum, item)
	}
	sum = SUB(DecimalPrecision, first, restsum)
	return sum
}

//
//================================================
//
// Function 03.06a - DIFxs
//
// DIFxs subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 50 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func DIFxs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return DIF(SuperPlasmStdMathPrecision, first, rest...)
}

//
//================================================
//
// Function 03.06b - DIFxc
//
// DIFxc subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 150 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func DIFxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return DIF(SuperPlasmMaxMathPrecision, first, rest...)
}

//
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 04.01 - MULx
//
// MULx multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
// Total number of digits is equal to the Precision specified in the TotalDecimalPrecision variable
func MULx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Mul(result, member1, member2)
	return result
}

//
//================================================
//
// Function 04.02 - MULs
//
// MULs multiplies two decimals within LOCPrecisionContext Context
// Total number of digits is equal to the Precision specified in LOCPrecisionContext
func MULs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Mul(result, member1, member2)
	return result
}

//
//================================================
//
// Function 04.03 - MULxc
//
// MULxc multiplies two decimals within an elastically modified Precision CryptoplasmPrecisionContext Context
// The elastic Precision's decimal limit is set to LOCMaxMathPrecision (it grows up to this value),
// while the integer precision scales without any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.195 digits, 100 units less than uint32.
// This is however expected never to be achieved.
func MULxc(member1, member2 *p.Decimal) *p.Decimal {
	var (
		result           = new(p.Decimal)
		DecimalPrecision uint32
	)

	IntegerDigitsMember1 := Count4Coma(member1)  //int64
	IntegerDigitsMember2 := Count4Coma(member2)  //int64
	DecimalDigitsMember1 := 0 - member1.Exponent //int32
	DecimalDigitsMember2 := 0 - member2.Exponent //int32

	IntegerSumInt64 := IntegerDigitsMember1 + IntegerDigitsMember2 //int64 9.223.372.036.854.775.807
	DecimalSumInt32 := DecimalDigitsMember1 + DecimalDigitsMember2 //int32 2.147.483.647

	IntegerSumUint32 := uint32(IntegerSumInt64) // from 9.223.372.036.854.775.807 to max 4.294.967.295
	DecimalSumUint32 := uint32(DecimalSumInt32) // from 2.147.483.647 to max 4.294.967.295

	//Max IntegerSum can be 4.294.967.295
	//Max DecimalSum is limited to 100.
	//As these are added to give the total precision, Max IntegerSum can be as high as 4.294.967.195

	if DecimalSumUint32 < SuperPlasmMaxMathPrecision {
		DecimalPrecision = DecimalSumUint32
	} else {
		DecimalPrecision = SuperPlasmMaxMathPrecision
	}
	MultiplicationPrecision := IntegerSumUint32 + DecimalPrecision

	cc := c.WithPrecision(MultiplicationPrecision)
	_, _ = cc.Mul(result, member1, member2)

	result = TruncateCustom(result, DecimalPrecision)
	return result
}

//
//================================================
//
// Function 04.04 - PRDx
//
// PRDx multiplies multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func PRDx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		product     = new(p.Decimal)
		restproduct = p.NFI(1)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Mul(restproduct, restproduct, item)
	}
	_, _ = cc.Mul(product, first, restproduct)

	return product
}

//
//================================================
//
// Function 04.05 - PRDs
//
// PRDs multiplies multiple decimals within CryptoplasmPrecisionContext Context
func PRDs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		product     = new(p.Decimal)
		restproduct = p.NFI(1)
	)

	for _, item := range rest {
		_, _ = c.Mul(restproduct, restproduct, item)
	}
	_, _ = c.Mul(product, first, restproduct)

	return product
}

//
//================================================
//
// Function 04.06 - PRDxc
//
// PRDxc multiplies two decimals within an elastically modified Precision CryptoplasmPrecisionContext Context
// The elastic Precision's decimal limit is set to 100, while the integer precision scales without any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.195 digits, 100 units less than uint32.
// This is however expected never to happen.
func PRDxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		product     = new(p.Decimal)
		restproduct = p.NFI(1)
	)

	for _, item := range rest {
		restproduct = MULxc(restproduct, item)
	}
	product = MULxc(first, restproduct)
	_, _ = c.Mul(product, first, restproduct)

	return product
}

//
//================================================
//
// Function 04.07 - POWx
//
// POWx computes x ** y within a custom Precision modified CryptoplasmPrecisionContext Context
func POWx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Pow(result, member1, member2)
	return result
}

//
//================================================
//
// Function 04.08 - POWs
//
// POWs computes x ** y within CryptoplasmPrecisionContext Context
func POWs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)

	_, _ = c.Pow(result, member1, member2)
	return result
}

//
//================================================
//
// Function 04.09 - POWxc
//
// POWxcs computes x ** y within an elastically custom modified Precision LOCPrecisionContext Context
// The elastic Precision's decimal limit is chosen by the user, while the integer precision scales without
// any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.295 - "chosen decimal precision" digits.
// This is however expected never to happen.
// Number of digit of a^b is D=1+b*log(10,a)

func POWxcs(DecimalNumber uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	var Logarithm = new(p.Decimal)

	//Getting the number of Digits the power would have
	_, _ = c.Log10(Logarithm, member1)
	Digits := ADDxc(p.NFI(1), MULxc(member2, Logarithm))
	DigitsR := TruncateCustom(Digits, 0)
	DigitsRI := uint32(p.INT64(DigitsR))

	TotalPowerPrecision := DigitsRI + DecimalNumber

	cc := c.WithPrecision(TotalPowerPrecision)
	_, _ = cc.Pow(result, member1, member2)

	return result
}

//
//================================================
//
// Function 04.10 - POWxc
//
// POWxc computes x ** y within an elastically modified Precision LOCPrecisionContext Context
// Same as POWxcs, the custom Decimal limit is set to LOCMaxMathPrecision (150)
func POWxc(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	result = POWxcs(SuperPlasmMaxMathPrecision, member1, member2)

	return result
}

//
//================================================
//
// Function 04.11 - Logarithm
//
// POWxc computes x ** y within an elastically modified Precision LOCPrecisionContext Context
// Logarithm returns the logarithm from "number" in base "base".
func Logarithm(base, number *p.Decimal) *p.Decimal {
	var (
		LogBase   = new(p.Decimal)
		LogNumber = new(p.Decimal)
	)
	//For LogBase and LogNumber Context precision
	//2+24 Context precision is enough, for base and number below e^100
	//if such were the case, a 3+24 (CryptoplasmCurrencyPrecision)
	//precision would be required. However e^100 has an Integer of 44 digits, namely
	//26.881.171.418.161.354.484.126.255.515.800.135.873.611.118
	//So one would have to need to compute the OverSend for a CP amount
	//bigger than this number to have the need to use a 3+24 context precision,
	//for computing the first logarithm below. Therefore 2+24 context precision
	//for both logarithms should be enough
	//27 Context Precision would be enough to compute the needed logarithm
	//for many more coins that could ever be minted until the End of the Universe.
	//if the Cryptoplasm emission would be repeated for every subsequent 524.596.891 Blocks (1 ERA)
	//1(ERA ~ 107 to 110 Trillion CP).

	//As the resulted LNs have the same number of digits for their integer part
	//a context Precision of 1+24 would always be enough, as the division would always
	//look like 1,.....

	//+3 is used, so such a high amount of coins to compute the OverSend for will also
	//work, and, as has been tested, indeed the code allows it to work.

	NumberDigits := number.NumDigits()
	IP := 2*SuperPlasmCurrencyPrecision + uint32(NumberDigits)
	cc := c.WithPrecision(IP)
	_, _ = cc.Ln(LogBase, base)
	_, _ = cc.Ln(LogNumber, number)
	CustomLog := DIVx(IP, LogNumber, LogBase)
	return CustomLog
}

//
//================================================
//
// Function 05.01 - DIVx
//
// DIVx divides two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIVx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Quo(result, member1, member2)
	return result
}

//
//================================================
//
// Function 05.02 - DIVs
//
// DIVs divides two decimals within CryptoplasmPrecisionContext Context
func DIVs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Quo(result, member1, member2)
	return result
}

//
//================================================
//
// Function 05.02 - DIVxc
//
// DIVxc divides 2 numbers with elastic integer precision and 100 max decimal precision
func DIVxc(member1, member2 *p.Decimal) *p.Decimal {
	var (
		result           = new(p.Decimal)
		IntegerPrecision uint32
	)

	IntegerDigitsMember1 := Count4Coma(member1) //int64		//Number of integer digits
	IntegerDigitsMember2 := Count4Coma(member2) //int64		//Number of integer digits

	DecimalDigitsMember1 := 0 - member1.Exponent //int32		//Number of decimals digits
	DecimalDigitsMember2 := 0 - member2.Exponent //int32		//Number of decimals digits

	NumberDigitsMember1 := member1.NumDigits() //Total Number of digits
	NumberDigitsMember2 := member2.NumDigits() //Total Number of digits

	IntegerMember1 := RemoveDecimals(member1) //Integer Value without decimals
	IntegerMember2 := RemoveDecimals(member2) //Integer Value without decimals

	if DecimalGreaterThan(IntegerMember1, p.NFI(0)) == true && DecimalGreaterThan(IntegerMember2, p.NFI(0)) {
		//Case 1 Integer Part is similar
		//fmt.Println("Case1")
		if DecimalEqual(IntegerMember1, IntegerMember2) == true {
			//fmt.Println("Case1.1")
			if DecimalGreaterThanOrEqual(member1, member2) == true {
				//fmt.Println("Case1.1.1")
				IntegerPrecision = 1
			} else {
				//fmt.Println("Case1.1.2")
				IntegerPrecision = 0
			}
		} else if DecimalGreaterThan(IntegerMember1, IntegerMember2) == true {
			//fmt.Println("Case1.2")
			if IntegerDigitsMember1 == IntegerDigitsMember2 {
				//fmt.Println("Case1.2.1")
				IntegerPrecision = 1
			} else if IntegerDigitsMember1 > IntegerDigitsMember2 {
				//fmt.Println("Case1.2.2")
				IntegerPrecision = uint32(IntegerDigitsMember1) - uint32(IntegerDigitsMember2) + 1
				//fmt.Println("IntegerPrecision is",IntegerPrecision)
			}
		} else {
			//fmt.Println("Case1.3")
			IntegerPrecision = 0
		}
	} else if DecimalGreaterThan(IntegerMember1, p.NFI(0)) == true && DecimalEqual(IntegerMember2, p.NFI(0)) {
		//Case 2 Integer Part of member2 is zero
		fmt.Println("Case2")
		if int32(NumberDigitsMember2) == DecimalDigitsMember2 {
			//fmt.Println("Case2.1")
			IntegerPrecision = uint32(IntegerDigitsMember1) + 1
		} else {
			//fmt.Println("Case2.2")
			Zeros := DecimalDigitsMember2 - int32(NumberDigitsMember2)
			IntegerPrecision = uint32(IntegerDigitsMember1) + 1 + uint32(Zeros)
		}
	} else if DecimalGreaterThan(IntegerMember2, p.NFI(0)) == true && DecimalEqual(IntegerMember1, p.NFI(0)) {
		//Case 3 Integer Part of member1 is zero
		//fmt.Println("Case3")
		IntegerPrecision = 0
	} else if DecimalEqual(IntegerMember1, p.NFI(0)) && DecimalEqual(IntegerMember2, p.NFI(0)) {
		//Case 4 both Integer Parts are zero
		//fmt.Println("Case4")
		Zeros1 := DecimalDigitsMember1 - int32(NumberDigitsMember1)
		Zeros2 := DecimalDigitsMember2 - int32(NumberDigitsMember2)
		if Zeros1 < Zeros2 {
			//fmt.Println("Case4.1")
			IntegerPrecision = uint32(Zeros2-Zeros1) + 1
		} else if Zeros1 > Zeros2 {
			//fmt.Println("Case4.2")
			IntegerPrecision = 0
		} else if Zeros1 == Zeros2 {
			//fmt.Println("Case4.3")
			if DecimalLessThan(member1, member2) == true {
				//fmt.Println("Case4.3.1")
				IntegerPrecision = 0
			} else {
				//fmt.Println("Case4.3.2")
				IntegerPrecision = 1
			}
		}
	}

	TotalDivisionPrecision := IntegerPrecision + SuperPlasmMaxMathPrecision
	result = DIVx(TotalDivisionPrecision, member1, member2)
	return result
}

//
//================================================
//
// Function 05.04 - DivInt
//
// DivInt returns the integer part of x divided by y
// It is equal to x // y
// Returned Value is also of decimal Type
func DivInt(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1, member2) //DivisionContextPrecision
	cc := c.WithPrecision(DCP)
	_, _ = cc.QuoInteger(result, member1, member2)
	return result
}

//
//================================================
//
// Function 05.05 - DivMod
//
// DivMod returns the remainder from the division of x to y
// It is equal to x % y
// Returned Value is also of decimal Type
func DivMod(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1, member2) //DivisionContextPrecision
	divresult := TruncateCustom(DivInt(member1, member2), 0)
	result = SUBx(DCP, member1, MULx(DCP, member2, divresult))
	return result
}

//
//================================================
//	05a Mean Functions:
// 		Different types of means used for computing purposes
//		In specific ways
//================================================
//
// Function 05a.01 - TwoMean
//
// TwoMean returns the mean of two decimals
func TwoMean(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1, member2) //DivisionContextPrecision
	result = DIVx(DCP, ADDxc(member2, member2), p.NFI(2))
	return result
}

//
//================================================
//	06 Truncate Functions:
// 		Functions used to Truncate Decimals to specific precision
//		In specific ways
//================================================
//
// Function 06.01 - TruncateCustom
//
// TruncateCustom truncates the decimal to the specified precision number
func TruncateCustom(Number *p.Decimal, DecimalPrecision uint32) *p.Decimal {
	var result = new(p.Decimal)

	NumberDigits := Count4Coma(Number)
	TruncatingContextPrecision := uint32(NumberDigits) + DecimalPrecision
	cc := c.WithPrecision(TruncatingContextPrecision)

	CSP := 0 - int32(DecimalPrecision)
	_, _ = cc.Quantize(result, Number, CSP)
	return result
}

//
//================================================
//
// Function 06.02 - TruncSeed
//
// TruncSeed truncates the decimal to SuperPlasmXPPrecision
// XP has a decimal precision of 8 Decimals
func TruncSeed(SeedNumber *p.Decimal) *p.Decimal {
	return TruncateCustom(SeedNumber, SuperPlasmXPPrecision)
}

//
//================================================
//
// Function 06.03 - TruncToCurrency
//
// TruncToCurrency truncates the decimal to LOCCurrencyPrecision
// LOC Currency Precision is currently set to 24 Decimals
// It is Context Precision Independent
func TruncToCurrency(Amount2BecomeCurrency *p.Decimal) *p.Decimal {
	return TruncateCustom(Amount2BecomeCurrency, SuperPlasmCurrencyPrecision)
}

//
//================================================
//
// Function 06.03 - TruncPercent
//
// TruncPercent truncates the decimal to LOCPromillePrecision
// Promille has 3 Decimals precision
// It is Context Precision Independent
func TruncPercent(Amount2BeTruncated *p.Decimal) *p.Decimal {
	return TruncateCustom(Amount2BeTruncated, SuperPlasmPromillePrecision)
}

//
//================================================
//	07 List Function:
// 		Functions that operate on slices of decimals
//		for different Purposes,
//		basically "emulating" Python List capability.
//================================================
//
// Function 07.01 - SumDL
//
// SumDL short for SumDecimalList, return the sum of
// the decimals within the list/slice
func SumDL(a []*p.Decimal) *p.Decimal {
	var sum = new(p.Decimal)

	for i := 0; i < len(a); i++ {
		sum = ADDs(sum, a[i])
	}
	return sum
}

//
//================================================
//
// Function 07.02 - LastDE
//
// LastDE short for LastDecimalElement, returns the last element
// in the slice. Equivalent to pythons -1 index
func LastDE(a []*p.Decimal) *p.Decimal {
	Length := len(a)
	LastElementIndex := Length - 1
	LastElement := a[LastElementIndex]
	return LastElement
}

//
//================================================
//
// Function 07.03 - AppDec
//
// AppDec creates a new bigger slice from the 2 slices used as input
// slices must be made of decimals
func AppDec(w1, w2 []*p.Decimal) []*p.Decimal {
	w3 := append(w1, w2...)
	return w3
}

//
//================================================
//
// Function 07.04 - Reverse
//
// Returns the Reverse of the Slice/Lists
func Reverse(a []*p.Decimal) []*p.Decimal {
	var Reversed = make([]*p.Decimal, 0)
	Length := len(a)
	LastElementIndex := Length - 1
	for i := LastElementIndex; i >= 0; i-- {
		Reversed = append(Reversed, a[i])
	}
	return Reversed
}

//
//================================================
//
// Function 07.05 - PrintDL
//
// PrintDL short for PrintDecimalList, prints the decimals
// within the given list/slice
func PrintDL(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Println("Element is,", a[i])
	}
}

//
//================================================
//
// Function 07.06 - WriteList
//
// WriteList writes the strings from the slice to an external file
// as Name can be used "File.txt" as the output file.
func WriteList(Name string, List []string) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}

//
//================================================
//	08 Digit Manipulations Function:
// 		Operations done on the number of Digits of a decimal
//		Used for setting elastic precision on different other functions
//================================================
//
// Function 08.01 - RemoveDecimals
//
// RemoveDecimals removes the decimals and leaves the resulted number
// without them
func RemoveDecimals(Number *p.Decimal) *p.Decimal {
	var Whole = new(p.Decimal)
	NumberDigits := Number.NumDigits()
	cc := c.WithPrecision(uint32(NumberDigits))
	_, _ = cc.Floor(Whole, Number)
	return Whole
}

//
//================================================
//
// Function 08.02 - Count4Coma
//
// Count4Coma returns the number of digits before precision
func Count4Coma(Number *p.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Int64Digits := Whole.NumDigits() //int64, up to 9223372036854775807
	return Int64Digits
}

//
//================================================
//
//================================================
//	09 LOC Amount String Manipulation Function:
// 		Functions that manipulate CP Amounts
//		formatting them for displaying purposes.
//================================================
//
// Function 09.01 - CPConvert2AU
//
// Convert2AU converts a CryptoPlasm amount into Atomic Units
func Convert2AU(cpAmount *p.Decimal) *p.Decimal {
	tcpAmount := TruncToCurrency(cpAmount)
	NumberDigits := Count4Coma(cpAmount)
	IP := uint32(NumberDigits) + SuperPlasmCurrencyPrecision
	AU := MULx(IP, tcpAmount, AUs)

	return AU
}

//
//================================================
//
// Function 09.02 - YoctoPlasm2String
//
// AttoPlasm2String converts a CryptoPlasm AUs (AttoPlasms)
// into a slice of strings
func AttoPlasm2String(Number *p.Decimal) []string {
	var SliceStr []string
	Ten := p.NFI(10)
	AuDigits := Number.NumDigits()
	Exp := AuDigits - 1
	IP := uint32(AuDigits)
	//Exp := p.NFI(NumberDigitsAU - 1)
	ToSequence := Number
	for i := Exp; i >= 0; i-- {
		idec := p.NFI(i)
		Power := POWx(IP, Ten, idec)
		Division := DIVx(IP, ToSequence, Power)
		DigitIs := TruncateCustom(Division, 0)
		DI := p.INT64(DigitIs)
		DigitIsString := strconv.Itoa(int(DI))
		SliceStr = append(SliceStr, DigitIsString)

		Rest := SUBx(IP, Division, DigitIs)
		SmallAU := MULx(IP, Rest, Power)
		ToSequence = SmallAU
	}
	return SliceStr
}

//
//================================================
//
// Function 09.03 - CPAmountConv2Print
//
// CPAmountConv2Print converts CryptoPlasm amount into a string
// to be used for printing purposes.
// For now only a "." as decimal character is implemented.
// Different Schemas can be added (for instance using coma<,> as decimal separator
// instead of point<.>; Or using points for thousand separator,
// or even separating at 2 position for Lakhs and Crores.
func KosonicDecimalConversion(cpAmount *p.Decimal) string {
	var (
		StringResult  string
		ComaPosition  int64
		PointPosition int64
		DigitTier     int64
	)

	//String Variable
	DecimalSeparator := ","
	ThousandSeparator := "."
	InsertFront := "["
	InsertMiddle := "|"
	InsertEnd := "]"

	if DecimalEqual(cpAmount, p.NFI(0)) == true {
		StringResult = "0,[000|000|000][000|000|000]"
	} else {
		Prec := int64(SuperPlasmCurrencyPrecision)
		AU := Convert2AU(cpAmount)
		SliceStr := AttoPlasm2String(AU)
		NumberDigits := Count4Coma(AU)

		InsertString := func(a []string, index int64, value string) []string {
			if int64(len(a)) == index { // nil or empty slice or after last element
				return append(a, value)
			}
			a = append(a[:index+1], a[index:]...) // index < len(a)
			a[index] = value
			return a
		}

		//Computing the Decimal Separator position
		if NumberDigits <= (Prec + 1) {
			ComaPosition = 1
		} else {
			ComaPosition = NumberDigits - Prec
		}
		//Inserting the Decimal Separator
		SliceStr = InsertString(SliceStr, ComaPosition, DecimalSeparator)

		//Computing the 1000 Separator positions
		Difference := NumberDigits - (Prec + 1)
		if Difference%3 == 0 {
			DigitTier = 1
		} else if Difference%3 == 1 {
			DigitTier = 2
		} else if Difference%3 == 2 {
			DigitTier = 3
		}
		TSNumber := (NumberDigits - (Prec + 1)) / 3

		//Adding the 1000 Separator as points
		for i := int64(1); i <= TSNumber; i++ {
			PointPosition = (i-1)*4 + DigitTier
			SliceStr = InsertString(SliceStr, PointPosition, ThousandSeparator)
		}

		//fmt.Println("new slice is", SliceStr)
		//fmt.Println("Slice Str cu virgula si 1000 separator este", len(SliceStr))

		//Adding Decimal Separators
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)), InsertEnd)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-4), InsertMiddle)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-8), InsertMiddle)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-12), InsertFront)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-13), InsertEnd)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-17), InsertMiddle)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-21), InsertMiddle)
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)-25), InsertFront)

		//Removing "0," from the SliceString, displaying only Decimals, in case os subunitary values.
		if len(SliceStr) == 28 && SliceStr[0] == "0" {
			SliceStr = SliceStr[2:]
		}

		//Converting Slice to string
		for i := 0; i < len(SliceStr); i++ {
			StringResult = StringResult + SliceStr[i]
		}
	}

	return StringResult
}

//
//================================================
//
// Function 09.03 - XPAmountConv2Print
//
// XPAmountConv2Print converts the BlockHeight decimal into a string
// to be used for printing purposes. A "." is inserted ever 1000.
// For now only a "." as decimal character is implemented.
// Different Schemas can be added (for instance using coma<,> as decimal separator
// instead of point<.>; Or using points for thousand separator,
// or even separating at 2 position for Lakhs and Crores.
func MKSP2Print(MKSP *p.Decimal) string {
	var (
		StringResult  string
		DigitTier     int64
		PointPosition int64
	)

	//String Variable
	//DecimalSeparator := ","
	ThousandSeparator := "."
	InsertFront := "["
	//InsertMiddle := "|"
	InsertEnd := "]"

	if DecimalEqual(MKSP, p.NFI(0)) == true {
		StringResult = InsertFront + "ZERO" + InsertEnd
	} else {
		//InsertString Function
		InsertString := func(a []string, index int64, value string) []string {
			if int64(len(a)) == index { // nil or empty slice or after last element
				return append(a, value)
			}
			a = append(a[:index+1], a[index:]...) // index < len(a)
			a[index] = value
			return a
		}

		NumberDigits := Count4Coma(MKSP)
		SliceStr := AttoPlasm2String(MKSP)

		//Computing the 1000 Separator positions
		Difference := NumberDigits - 1
		if Difference%3 == 0 {
			DigitTier = 1
		} else if Difference%3 == 1 {
			DigitTier = 2
		} else if Difference%3 == 2 {
			DigitTier = 3
		}
		TSNumber := (NumberDigits - 1) / 3

		//Adding the 1000 Separator as points
		for i := int64(1); i <= TSNumber; i++ {
			PointPosition = (i-1)*4 + DigitTier
			SliceStr = InsertString(SliceStr, PointPosition, ThousandSeparator)
		}

		//Inserting Starting and Ending Brackets
		SliceStr = InsertString(SliceStr, int64(len(SliceStr)), InsertEnd)
		SliceStr = InsertString(SliceStr, int64(0), InsertFront)

		//Converting Slice to string
		for i := 0; i < len(SliceStr); i++ {
			StringResult = StringResult + SliceStr[i]
		}
	}

	return StringResult + ".[Meta-Kosonic Super-Power]"
}
