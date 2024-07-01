package evmUtils

import (
	"eigenlayer_hack/evm/evmStructs"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

/*
handleIntegerTypes helper function for handling XintXXX and []XintXXX types. Use the bool (isArray) to determine which returned
value to be used as this function returns both single variable and a variable array. If the flag is true use the array variable
else use the single variable
*/
func handleIntegerTypes(arrLength int, startingIndex int, specifiedType string, chunked *evmStructs.Chunker) (*big.Int, []*big.Int, bool, error) {
	// Parse the given type string
	bS, sF, aF, parseErr := parseIntType(specifiedType)

	if parseErr != nil {
		return nil, nil, false, parseErr
	}

	// Check if the requested type is array or a single value
	if aF {
		// Handle the INTEGER ARRAY (Dynamic Type)
		retArr := []*big.Int{}

		for arrayIdx := 1; arrayIdx < arrLength+1; arrayIdx++ {
			// Decode the array elements one by one
			arrElem, arrDecodeErr := decodeBytesToInteger(chunked.GetChunk(startingIndex+arrayIdx), bS,
				sF)

			if arrDecodeErr != nil {
				return nil, nil, false, arrDecodeErr
			}

			retArr = append(retArr, arrElem)
		}

		// Return the array
		return nil, retArr, true, nil

	} else if !aF {
		// Handle the INTEGER VALUE (Static Type)
		retInt, decodeErr := decodeBytesToInteger(chunked.GetChunk(startingIndex), bS, sF)

		if decodeErr != nil {
			return nil, nil, false, decodeErr
		}

		// Return the variable
		return retInt, nil, aF, nil

	}

	// In theory this line will never be executed
	return nil, nil, aF, errors.New("error while handling the data")
}

/*
handleAddressTypes helper function for handling address and []address types. Use the bool (isArray) to determine which returned
value to be used as this function returns both single variable and a variable array. If the flag is true use the array variable
else use the single variable
*/
func handleAddressTypes(arrLength int, startingIndex int, chunked *evmStructs.Chunker) (common.Address, []common.Address, bool, error) {
	if arrLength > 0 {
		// Handle the dynamic ADDRESS Type
		dynamicAddressArr := []common.Address{}

		for arrayIdx := 1; arrayIdx < arrLength+1; arrayIdx++ {
			// Decode the array elements one by one
			tAddrs := common.BytesToAddress(chunked.GetChunk(startingIndex + arrayIdx))

			dynamicAddressArr = append(dynamicAddressArr, tAddrs)
		}

		return common.HexToAddress("0x0"), dynamicAddressArr, true, nil

	} else if arrLength == 0 {
		// Handle the static ADDRESS Type
		tAddress := common.BytesToAddress(chunked.GetChunk(startingIndex))

		return tAddress, []common.Address{}, false, nil
	}

	// In theory this line will never be executed
	return common.HexToAddress("0x0"), nil, false, errors.New("error while handling the data")
}

/*
handleBoolTypes helper function for handling bool and []bool types. Use the bool (isArray) to determine which returned
value to be used as this function returns both single variable and a variable array. If the flag is true use the array variable
else use the single variable
*/
func handleBoolTypes(arrLength int, startingIndex int, chunked *evmStructs.Chunker) (bool, []bool, bool, error) {
	if arrLength > 0 {
		// Handle the dynamic ADDRESS Type
		dynamicBoolArr := []bool{}

		for arrayIdx := 1; arrayIdx < arrLength+1; arrayIdx++ {
			// Decode the array elements one by one
			tBool, decodeErr := decodeBytesToBoolean(chunked.GetChunk(startingIndex + arrayIdx))

			if decodeErr != nil {
				return false, nil, false, decodeErr
			}

			dynamicBoolArr = append(dynamicBoolArr, tBool)
		}

		// Return array variable
		return false, dynamicBoolArr, true, nil

	} else if arrLength == 0 {

		retBool, decodeErr := decodeBytesToBoolean(chunked.GetChunk(startingIndex))

		if decodeErr != nil {
			return false, nil, false, decodeErr
		}

		// Return single variable
		return retBool, nil, false, nil
	}

	// In theory this line will never be executed
	return false, nil, false, errors.New("error while handling the data")
}

/*
handleByteTypes helper function for handling bytesXXX and []bytesXXX types. Use the bool (isArray) to determine which returned
value to be used as this function returns both single variable and a variable array. If the flag is true use the array variable
else use the single variable
*/
func handleByteTypes(arrLength int, startingIndex int, specifiedType string, chunked *evmStructs.Chunker) ([]byte, [][]byte, bool, error) {
	// "bytes" without bit size specified is a dynamic type, acts just like "string"
	if strings.ToLower(specifiedType) == "bytes" {
		retVal, _, decodeErr := bytesHandler(chunked, startingIndex)

		if decodeErr != nil {
			return nil, nil, false, decodeErr
		}

		return retVal, nil, false, nil
	}

	// Parse the given type string
	bS, aF, parseErr := parseByteType(specifiedType)

	if parseErr != nil {
		return nil, nil, false, parseErr
	}

	// Check if the requested type is array or a single value
	if aF {
		// Handle the INTEGER ARRAY (Dynamic Type)
		retArr := [][]byte{}

		for arrayIdx := 1; arrayIdx < arrLength+1; arrayIdx++ {
			// Decode the array elements one by one
			arrElem, _ := chunked.GetByteArrayByByteIndex(chunked.GetIndexOfStartingByte(startingIndex+arrayIdx), bS)

			// Append the elements
			retArr = append(retArr, arrElem)
		}

		// Return the array
		return nil, retArr, true, nil

	} else if !aF {
		// Handle the INTEGER VALUE (Static Type)
		retVal, _ := chunked.GetByteArrayByByteIndex(chunked.GetIndexOfStartingByte(startingIndex), bS)

		// Return the variable
		return retVal, nil, false, nil

	}

	// In theory this line will never be executed
	return nil, nil, aF, errors.New("error while handling the data")
}

/*
handleStringTypes helper function for handling string and []string types. Use the bool (isArray) to determine which returned
value to be used as this function returns both single variable and a variable array. If the flag is true use the array variable
else use the single variable
*/
func handleStringTypes(arrLength int, startingIndex int, chunked *evmStructs.Chunker) (string, []string, bool, error) {
	if arrLength > 0 {
		// Handle the dynamic STRING Type
		retArr := []string{}
		lastChunkIndex := startingIndex - 1

		for arrayIdx := 1; arrayIdx < arrLength+1; arrayIdx++ {
			// Decode the array elements one by one
			arrElem, newLastChunkIndex, handleErr := stringHandler(chunked, lastChunkIndex+1)

			if handleErr != nil {
				return "", nil, true, handleErr
			}

			// Append the elements
			retArr = append(retArr, arrElem)

			// Update the last chunk index
			lastChunkIndex = newLastChunkIndex
		}

		return "", retArr, true, nil

	} else if arrLength == 0 {
		// Handle the STRING Type
		retVal, _, handleErr := stringHandler(chunked, startingIndex)
		return retVal, nil, false, handleErr
	}

	return "", nil, false, errors.New("error while handling the data")
}

/*
stringHandler takes a *evmStructs.Chunker and starting chunk index (int) as arguments and decodes the string data
starting from the given chunk index. Returns the decoded string, chunk index of the last byte of the given string
and error
*/
func stringHandler(chunked *evmStructs.Chunker, startingChunk int) (string, int, error) {
	// Get the starting point
	strStartByteIndex, srtStartIndexErr := decodeBytesToInteger(chunked.GetChunk(startingChunk),
		256, false)

	if srtStartIndexErr != nil {
		return "", -1, srtStartIndexErr
	}

	strStartIndexINT := int(strStartByteIndex.Int64())
	strStartChunkIndex, _ := chunked.ByteIndexToChunk(strStartIndexINT)

	// Get the string length
	strLength, strLengthErr := decodeBytesToInteger(chunked.GetChunk(strStartChunkIndex),
		256, false)

	if strLengthErr != nil {
		return "", -1, strLengthErr
	}

	// Convert it to int
	strLengthINT := int(strLength.Int64())

	targetStringAsByte, lastChunkIndex := chunked.GetByteArrayByByteIndex(chunked.GetIndexOfStartingByte(strStartChunkIndex+1),
		strLengthINT)

	return string(targetStringAsByte), lastChunkIndex, nil
}

/*
bytesHandler takes a *evmStructs.Chunker and starting chunk index (int) as arguments and decodes the bytes data
starting from the given chunk index. Returns the decoded byte array, chunk index of the last byte of the given bytes
and error
*/
func bytesHandler(chunked *evmStructs.Chunker, startingChunk int) ([]byte, int, error) {
	// Get the starting point
	strStartByteIndex, srtStartIndexErr := decodeBytesToInteger(chunked.GetChunk(startingChunk),
		256, false)

	if srtStartIndexErr != nil {
		return []byte{}, -1, srtStartIndexErr
	}

	strStartIndexINT := int(strStartByteIndex.Int64())
	strStartChunkIndex, _ := chunked.ByteIndexToChunk(strStartIndexINT)

	// Get the string length
	strLength, strLengthErr := decodeBytesToInteger(chunked.GetChunk(strStartChunkIndex),
		256, false)

	if strLengthErr != nil {
		return []byte{}, -1, strLengthErr
	}

	// Convert it to int
	bytesLengthINT := int(strLength.Int64())

	targetStringAsByte, lastChunkIndex := chunked.GetByteArrayByByteIndex(chunked.GetIndexOfStartingByte(strStartChunkIndex+1),
		bytesLengthINT)

	return targetStringAsByte, lastChunkIndex, nil
}

/*
tupleHandler takes a *evmStructs.Chunker and starting chunk index (int) as arguments and extracts the data related to the
tuple from the whole data
*/
func tupleHandler(typeString string, chunked *evmStructs.Chunker, startingChunk int) (tupleTypes []string, data [][]byte, err error) {
	// Determine if the tuple contains dynamic type, as otherwise it is handled just like normally
	tupleTypes, hasDynamicType, isArray, err := parseTuple(typeString)

	data = [][]byte{}
	var chunkerData []byte

	if err != nil {
		return
	}

	if !hasDynamicType && !isArray {
		chunkerData, err = chunked.GetSlice(startingChunk, -1)
		data = append(data, chunkerData)
		return
	}

	// Get the starting point
	tupleStartByteIndex, err := decodeBytesToInteger(chunked.GetChunk(startingChunk),
		256, false)

	if err != nil {
		return
	}

	tupleStartIndexINT := int(tupleStartByteIndex.Int64())
	tupleStartChunkIndex, _ := chunked.ByteIndexToChunk(tupleStartIndexINT)

	// Handle the array case
	if isArray {
		// If it is an array we need to find its length
		arrayLength, arrLengthErr := decodeBytesToInteger(chunked.GetChunk(tupleStartChunkIndex),
			256, false)

		if arrLengthErr != nil {
			err = arrLengthErr
			return
		}

		arrayLengthINT := int(arrayLength.Int64())

		elemString := "tuple("

		for _, theType := range tupleTypes {
			elemString = elemString + theType + ", "
		}

		elemString = elemString[:len(elemString)-2] + ")"

		// Reset the tuple types
		tupleTypes = []string{}

		// Now we need to get the starting positions for each tuple
		for i := 1; i < arrayLengthINT+1; i++ {

			tupleTypes = append(tupleTypes, elemString)

		}

		chunkerData, err = chunked.GetSlice(tupleStartChunkIndex+1, -1)
		data = append(data, chunkerData)
		return
	}

	chunkerData, err = chunked.GetSlice(tupleStartChunkIndex, -1)
	data = append(data, chunkerData)
	return
}
