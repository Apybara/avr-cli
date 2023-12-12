package util

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// ToBinary converts a string to binary
// use 2 to represent space

func ToBinary(input string) string {
	var binaryStr string
	for _, char := range input {
		// Convert character to binary and format it to 8 bits
		binaryStr += fmt.Sprintf("%08b2", char)
	}
	return "3" + binaryStr
}

func BinaryToString(binaryStr string) (string, error) {
	var strBuilder strings.Builder

	// remove all 3
	binaryStr = strings.ReplaceAll(binaryStr, "3", "")
	binaryCodes := strings.Split(binaryStr, "2")

	for _, binaryCode := range binaryCodes {
		// ignore all 3
		if binaryCode == "" || binaryCode == "3" {
			continue
		}
		num, err := strconv.ParseInt(binaryCode, 2, 64)
		if err != nil {
			return "", err // handle invalid binary input
		}
		strBuilder.WriteByte(byte(num))
	}

	return strBuilder.String(), nil
}

// utf8StringToBigInt converts a UTF-8 string to a big.Int.
func Utf8StringToBigInt(s string) *big.Int {
	result := big.NewInt(0)
	temp := new(big.Int)

	for i := 0; i < len(s); i++ {
		result.Lsh(result, 8)
		temp.SetInt64(int64(s[i]))
		result.Add(result, temp)
	}

	return result
}

func BigIntStringToUtf8String(s string) string {
	bigInt := new(big.Int)
	bigInt.SetString(s, 10)
	return BigIntToUTF8String(bigInt)
}

// bigIntToUTF8String converts a big.Int to a UTF-8 string.
func BigIntToUTF8String(number *big.Int) string {
	bytes := number.Bytes()
	var result string
	for _, b := range bytes {
		result += string(b)
	}
	return result
}

// calculateBits returns the number of bits used to store a string in UTF-8.
func CalculateLength(s string) int {
	return len(s)
}
