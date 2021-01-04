package entity

import (
	"fmt"
	"math/big"
	"strings"
)

//WeiToETH convert Wei to ETH (10^-18)
func WeiToETH(wei *big.Int) float64 {
	var div float64 = 1000000000000000000 //10^18
	bigF := new(big.Float).SetInt(wei)
	f, _ := bigF.Float64()
	return f / div
}

//HexToBigInt convert Hexadeciaml string to UInt64
func HexToBigInt(hex string) (*big.Int, error) {
	numrStr := strings.Replace(hex, "0x", "", 1)
	num, ok := new(big.Int).SetString(numrStr, 16)
	if !ok {
		return num, ErrIncorrectValue
	}
	return num, nil
}

//BigIntToHex convert Int to Hexadeciaml string
func BigIntToHex(num *big.Int) string {
	return fmt.Sprintf("0x%x", num)
}
