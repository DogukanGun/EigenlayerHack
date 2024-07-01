package evmStructs

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/*
DataType Table

	ID	| 	Type
-------------------
	0	|	big.Int
	1	| 	[]big.Int
	2	|	string
	3	| 	[]string
	4	|	[]byte
	5	| 	[][]byte
	6	|	bool
	7	| 	[]bool
	8	|	address
	9	| 	[]address
	10  |   tuple
*/

type DecodeOutput struct {
	DecodedData interface{}
	DataType    int8
	DecodeErr   error
}

type DecodedLog struct {
	CalledAddress      common.Address
	FunctionSignature  string
	SignatureHash      string
	DecodedIndexedData []DecodeOutput
	DecodedData        []DecodeOutput
	Types              []string
	IndexedTypes       []string
	DecodeErr          error
	LogIndex           uint
}

type DecodedTx struct {
	ToAddress               common.Address
	FromAddress             common.Address
	CalledFunctionBytes     []byte
	CalledFunctionSignature string
	DecodedData             []DecodeOutput
	Types                   []string
	DecodeErr               error
}

func (dO *DecodeOutput) AsInt() (value *big.Int, err error) {
	// Check if the value is type int (ID == 0)
	if dO.DataType != 0 {
		err = errors.New("contained value is not integer")
		return
	}

	value, isOk := dO.DecodedData.(*big.Int)

	if !isOk {
		err = errors.New("error while converting value to *big.Int")
		return
	}

	return
}

func (dO *DecodeOutput) AsIntArray() (value []*big.Int, err error) {
	// Check if the value is type int (ID == 1)
	if dO.DataType != 1 {
		err = errors.New("contained value is not integer array")
		return
	}

	value, isOk := dO.DecodedData.([]*big.Int)

	if !isOk {
		err = errors.New("error while converting value to []*big.Int")
		return
	}

	return
}

func (dO *DecodeOutput) AsString() (value string, err error) {
	// Check if the value is type int (ID == 2)
	if dO.DataType != 2 {
		err = errors.New("contained value is not string")
		return
	}

	value, isOk := dO.DecodedData.(string)

	if !isOk {
		err = errors.New("error while converting value to string")
		return
	}

	return
}

func (dO *DecodeOutput) AsStringArray() (value []string, err error) {
	// Check if the value is type int (ID == 3)
	if dO.DataType != 3 {
		err = errors.New("contained value is not string array")
		return
	}

	value, isOk := dO.DecodedData.([]string)

	if !isOk {
		err = errors.New("error while converting value to []string")
		return
	}

	return
}

func (dO *DecodeOutput) AsBytes() (value []byte, err error) {
	// Check if the value is type int (ID == 4)
	if dO.DataType != 4 {
		err = errors.New("contained value is not bytes")
		return
	}

	value, isOk := dO.DecodedData.([]byte)

	if !isOk {
		err = errors.New("error while converting value to []byte")
		return
	}

	return
}

func (dO *DecodeOutput) AsBytesArray() (value [][]byte, err error) {
	// Check if the value is type int (ID == 5)
	if dO.DataType != 5 {
		err = errors.New("contained value is not bytes array")
		return
	}

	value, isOk := dO.DecodedData.([][]byte)

	if !isOk {
		err = errors.New("error while converting value to [][]byte")
		return
	}

	return
}

func (dO *DecodeOutput) AsBool() (value bool, err error) {
	// Check if the value is type int (ID == 6)
	if dO.DataType != 6 {
		err = errors.New("contained value is not bool")
		return
	}

	value, isOk := dO.DecodedData.(bool)

	if !isOk {
		err = errors.New("error while converting value to bool")
		return
	}

	return
}

func (dO *DecodeOutput) AsBoolArray() (value []bool, err error) {
	// Check if the value is type int (ID == 7)
	if dO.DataType != 7 {
		err = errors.New("contained value is not bool array")
		return
	}

	value, isOk := dO.DecodedData.([]bool)

	if !isOk {
		err = errors.New("error while converting value to []bool")
		return
	}

	return
}

func (dO *DecodeOutput) AsAddress() (value common.Address, err error) {
	// Check if the value is type int (ID == 8)
	if dO.DataType != 8 {
		err = errors.New("contained value is not address")
		return
	}

	value, isOk := dO.DecodedData.(common.Address)

	if !isOk {
		err = errors.New("error while converting value to common.Address")
		return
	}

	return
}

func (dO *DecodeOutput) AsAddressArray() (value []common.Address, err error) {
	// Check if the value is type int (ID == 9)
	if dO.DataType != 9 {
		err = errors.New("contained value is not address array")
		return
	}

	value, isOk := dO.DecodedData.([]common.Address)

	if !isOk {
		err = errors.New("error while converting value to []common.Address")
		return
	}

	return
}

func (dO *DecodeOutput) AsDecodeOutput() (value []DecodeOutput, err error) {
	// Check if the value is type int (ID == 10)
	if dO.DataType != 10 {
		err = errors.New("contained value is not DecodeOutput")
		return
	}

	value, isOk := dO.DecodedData.([]DecodeOutput)

	if !isOk {
		err = errors.New("error while converting value to DecodeOutput")
		return
	}

	return
}
