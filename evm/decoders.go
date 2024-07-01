package evmUtils

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

/*
decodeBytesToInteger takes a byte array as input and converts that array to a bigInt based on the specified bitSize
Example: int256 -> bitSize = 256, uint160 -> bitSize = 160, int24 -> bitSize = 24
signed is flag to determine if the output should be a signed integer or not
Example: int256 -> signed = true, uint160 -> signed = false, int24 -> signed = true

Returns a big int and error
*/
func decodeBytesToInteger(input []byte, bitSize int, signed bool) (*big.Int, error) {

	// Remove padding
	if len(input) < bitSize/8 {
		log.Println("Supplied bytes are not enough for the specified bitSize")
		return big.NewInt(0), errors.New("len(input) < bitSize/8")
	}
	input = input[len(input)-(bitSize/8):]

	// Initialize a new big int
	ret := new(big.Int)

	// Set the value
	ret.SetBytes(input)

	// Check if it should be negative
	if ret.Bit(bitSize-1) == 1 && signed {
		unsignedBitsize := uint(bitSize)
		maxValue := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, unsignedBitsize), common.Big1)
		ret.Add(maxValue, new(big.Int).Neg(ret))
		ret.Add(ret, common.Big1)
		ret.Neg(ret)
	}

	return ret, nil
}

/*
decodeBytesToBoolean takes a byte array as input and converts that array to a boolean
Returns a boolean and error
*/
func decodeBytesToBoolean(input []byte) (bool, error) {
	// Convert the value to integer
	value, status := decodeBytesToInteger(input, 256, false)

	if status != nil {
		return false, errors.New("unsuccessful")
	}

	if value.Int64() > 0 {
		return true, nil
	}

	return false, nil
}
