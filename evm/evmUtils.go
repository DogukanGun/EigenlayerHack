package evmUtils

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-redis/redis/v8"
	"regexp"
	"solity/schemas"
	"solity/utils/database"
	"solity/utils/ethereum/node"
	"solity/utils/evm/evmInterfaces"
	"solity/utils/evm/evmStructs"
	"solity/utils/logger"
	"strings"
	"sync"
)

/*
DecodeInput decodes the supplied byte array into the given types. Currently, supports integer, bytes, address, string,
bool types (and any variations of them)
*/
func DecodeInput(input []byte, types []string) (ret []evmStructs.DecodeOutput, err error) {
	// Initialize return
	ret = []evmStructs.DecodeOutput{}

	for i := 0; i < len(types); i++ {
		ret = append(ret, evmStructs.DecodeOutput{})
	}

	// Create signal for inter thread communication
	txThreadStatusSignal := new(sync.WaitGroup)
	txThreadStatusSignal.Add(len(types))

	// Initialize the Chunker
	chunkedData, chunkErr := evmStructs.NewChunker(input, 32)

	if chunkErr != nil {
		err = chunkErr
		return
	}

	if chunkedData.GetNumberOfChunks() < len(types) {
		err = errors.New("supplied data and number of specified types missmatch")
		return
	}

	for i, typ := range types {
		go func(idx int, t string) {
			defer txThreadStatusSignal.Done()

			// For ease of use switch to the lowercase
			lowerCaseType := strings.ToLower(t)

			// Assume that everything is array and initialize the array variables
			startingChunkIndex := idx
			arrayLengthINT := 0
			// Check if data is tuple
			if strings.Contains(lowerCaseType, "tuple") {
				tupleOutput := []evmStructs.DecodeOutput{}

				tupleTypes, tupleData, tupleErr := tupleHandler(lowerCaseType, &chunkedData, startingChunkIndex)

				if tupleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: tupleErr}
					return
				}

				// Recursion
				for _, tupleBytes := range tupleData {
					tupleDecoded, tupleErr := DecodeInput(tupleBytes, tupleTypes)

					if tupleErr != nil {
						ret[idx] = evmStructs.DecodeOutput{DecodeErr: tupleErr}
						return
					}

					tupleOutput = append(tupleOutput, tupleDecoded...)
				}

				ret[idx] = evmStructs.DecodeOutput{DecodedData: tupleOutput, DataType: 10}
				return
			}

			// Check if the type is an array (Dynamic Data Structure)
			if strings.Contains(lowerCaseType, "[") {
				// Get the starting chunk
				startingByteIndex, startingByteIndexDecodeErr := decodeBytesToInteger(chunkedData.GetChunk(idx),
					256, false)

				if startingByteIndexDecodeErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: startingByteIndexDecodeErr}
					return
				}

				// Get Starting Location
				startingByteIdxINT := int(startingByteIndex.Int64())
				startingChunkIndex, _ = chunkedData.ByteIndexToChunk(startingByteIdxINT)

				// Get the array length
				arrayLength, arrLengthErr := decodeBytesToInteger(chunkedData.GetChunk(startingChunkIndex),
					256, false)

				if arrLengthErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: arrLengthErr}
					return
				}

				arrayLengthINT = int(arrayLength.Int64())
			}

			switch {
			case strings.Contains(lowerCaseType, "int"):
				// Handle integer types
				singleVar, arrayVar, isArray, handleErr := handleIntegerTypes(arrayLengthINT, startingChunkIndex, t, &chunkedData)

				// Error handling
				if handleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: handleErr}
					return
				}

				// Check if the returned value is a single variable or an array
				if isArray {
					// Add the array to the return array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: arrayVar, DataType: 1}
				} else {
					// Add the single int variable to the array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: singleVar, DataType: 0}
				}

			case strings.Contains(lowerCaseType, "address"):
				// Handle address types
				singleVar, arrayVar, isArray, handleErr := handleAddressTypes(arrayLengthINT, startingChunkIndex, &chunkedData)

				// Error handling
				if handleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: handleErr}
					return
				}

				// Check if the returned value is a single variable or an array
				if isArray {
					// Add the array to the return array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: arrayVar, DataType: 9}
				} else {
					// Add the single int variable to the array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: singleVar, DataType: 8}
				}

			case strings.Contains(lowerCaseType, "bool"):

				// Handle bool types
				singleVar, arrayVar, isArray, handleErr := handleBoolTypes(arrayLengthINT, startingChunkIndex, &chunkedData)

				// Error handling
				if handleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: handleErr}
					return
				}

				// Check if the returned value is a single variable or an array
				if isArray {
					// Add the array to the return array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: arrayVar, DataType: 7}
				} else {
					// Add the single int variable to the array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: singleVar, DataType: 6}
				}

			case strings.Contains(lowerCaseType, "bytes"):
				// Handle byte types
				singleVar, arrayVar, isArray, handleErr := handleByteTypes(arrayLengthINT, startingChunkIndex, t, &chunkedData)

				// Error handling
				if handleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: handleErr}
					return
				}

				// Check if the returned value is a single variable or an array
				if isArray {
					// Add the array to the return array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: arrayVar, DataType: 5}
				} else {
					// Add the single int variable to the array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: singleVar, DataType: 4}
				}

			case strings.Contains(lowerCaseType, "string"):
				// Handle STRING type
				// Handle integer types
				singleVar, arrayVar, isArray, handleErr := handleStringTypes(arrayLengthINT, startingChunkIndex, &chunkedData)

				// Error handling
				if handleErr != nil {
					ret[idx] = evmStructs.DecodeOutput{DecodeErr: handleErr}
					return
				}

				// Check if the returned value is a single variable or an array
				if isArray {
					// Add the array to the return array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: arrayVar, DataType: 3}
				} else {
					// Add the single int variable to the array
					ret[idx] = evmStructs.DecodeOutput{DecodedData: singleVar, DataType: 2}
				}

			default:
				return
			}
		}(i, typ)
	}

	// Wait for all threads to finish
	txThreadStatusSignal.Wait()

	return
}

func DecodeReceipt(rcpt *types.Receipt, sk evmStructs.SignatureKeeper) (dLogs []evmStructs.DecodedLog, err error) {

	for _, log := range rcpt.Logs {
		// Initialize the log object
		decodedLog := evmStructs.DecodedLog{}

		// Check the size of the topics array in order to prevent pointer errors
		if len(log.Topics) <= 0 {
			logger.LogW("given log does not have topics in it")
			continue
		}

		// Check if the emmited event is relevant
		eventSignature, eventHash, eventDataTypes, indexedDataTypes, gErr := sk.GetHash(log.Topics[0].Hex())

		if gErr != nil {
			continue
		}

		// Update the decoded data object
		decodedLog.CalledAddress = log.Address
		decodedLog.FunctionSignature = eventSignature
		decodedLog.SignatureHash = eventHash
		decodedLog.Types = eventDataTypes
		decodedLog.IndexedTypes = indexedDataTypes
		decodedLog.LogIndex = log.Index

		//Check if there are any indexed data
		if len(indexedDataTypes) > 0 {
			indexedData, dErr := decodeIndexedTopics(log.Topics[1:], indexedDataTypes)

			if dErr != nil {
				logger.LogW("skipping decode of indexed variables: " + dErr.Error())
				continue
			}

			// Add it to the decoded log object
			decodedLog.DecodedIndexedData = indexedData
		}

		// Decode the data
		decodedData, dErr := DecodeInput(log.Data, eventDataTypes)

		if dErr != nil {
			err = dErr
			return
		}

		// Add the decoded data to the decoded log object
		decodedLog.DecodedData = decodedData

		// Append the decoded data object to the returned variable
		dLogs = append(dLogs, decodedLog)
	}

	return
}

func DecodeLog(log *types.Log, sk evmStructs.SignatureKeeper) (decodedLog evmStructs.DecodedLog, err error) {
	// Initialize the log object
	decodedLog = evmStructs.DecodedLog{}

	// Check the size of the topics array in order to prevent pointer errors
	if len(log.Topics) <= 0 {
		err = errors.New("given log does not have topics in it")
		return
	}

	// Check if the emmited event is relevant
	eventSignature, eventHash, eventDataTypes, indexedDataTypes, gErr := sk.GetHash(log.Topics[0].Hex())

	if gErr != nil {
		err = gErr
		return
	}

	// Update the decoded data object
	decodedLog.CalledAddress = log.Address
	decodedLog.FunctionSignature = eventSignature
	decodedLog.SignatureHash = eventHash
	decodedLog.Types = eventDataTypes
	decodedLog.IndexedTypes = indexedDataTypes
	decodedLog.LogIndex = log.Index

	//Check if there are any indexed data
	if len(indexedDataTypes) > 0 {
		indexedData, dErr := decodeIndexedTopics(log.Topics[1:], indexedDataTypes)

		if dErr != nil {
			err = errors.New("skipping decode of indexed variables: " + dErr.Error())
			return
		}

		// Add it to the decoded log object
		decodedLog.DecodedIndexedData = indexedData
	}

	// Decode the data
	decodedData, dErr := DecodeInput(log.Data, eventDataTypes)

	if dErr != nil {
		err = dErr
		return
	}

	// Add the decoded data to the decoded log object
	decodedLog.DecodedData = decodedData

	return
}

func DecodeTxData(tx *types.Transaction, sk evmStructs.SignatureKeeper) (dTx evmStructs.DecodedTx, err error) {
	// Initialize the return variable
	dTx = evmStructs.DecodedTx{}

	// Get the Tx sender
	sender, err := node.GetTxSender(tx)

	if err != nil {
		return
	}

	// Fill the fields
	dTx.FromAddress = sender
	dTx.ToAddress = common.HexToAddress(tx.To().Hex())

	// Check if the tx has valid data
	if len(tx.Data()) < 4 {
		err = errors.New("tx data is not valid for parsing")
		return
	}

	// Fill the field
	dTx.CalledFunctionBytes = tx.Data()[:4]

	functionSignatureHash := common.BytesToHash(tx.Data()[:4])

	// Check if the function signature is relevant
	signature, _, dataTypes, _, err := sk.GetHash(functionSignatureHash.Hex())

	if err != nil {
		return
	}

	// Fill the field
	dTx.CalledFunctionSignature = signature
	dTx.Types = dataTypes

	decodedData, err := DecodeInput(tx.Data()[4:], dataTypes)

	if err != nil {
		return
	}

	// Fill the field
	dTx.DecodedData = decodedData

	return
}

func decodeIndexedTopics(topics []common.Hash, types []string) (decodedData []evmStructs.DecodeOutput, err error) {
	// Check if their lengths are matching
	if len(topics) != len(types) {
		err = errors.New("topics length and supplied input types length miss match")
		return
	}
	// Initialize buffer
	buffer := []byte{}

	for _, topic := range topics {
		buffer = append(buffer, topic.Bytes()...)
	}

	// Decode the indexed data buffer
	decodedData, err = DecodeInput(buffer, types)

	return
}

/*
GetFunctionSignatures extracts the function signatures from the bytecode stored in the given address
*/
func GetFunctionSignatures(client evmInterfaces.ByteCodeRequestor, targetAddress *common.Address) (functionSigs *[]string,
	err error) {
	// Request the byte code
	code, err := client.CodeAt(context.Background(), *targetAddress, nil)

	if err != nil {
		return
	}

	if len(code) == 0 {
		err = errors.New("no code data found at the given address")
		return
	}

	// Convert byte code to hex string
	hexCode := common.Bytes2Hex(code)

	// Initialize the regex selector

	/*
		Selector looks for the following byte code pattern:

			DUP1 (80)
			PUSH4 (63)
			ANY-8-CHARACTERS -> function signature
			EQ (14)
			PUSH2 (61)
			ANY-4-CHARACTERS -> jump location
			JUMPI (57)

		As functions are encoded like this in the hex byte codes
	*/

	selector, err := regexp.Compile("8063(.{8})1461.{4}57")

	if err != nil {
		return
	}

	matches := selector.FindAllStringSubmatch(hexCode, -1)
	functionSigsData := []string{}

	for _, match := range matches {
		// Check the length of the-each match
		if len(match) != 2 {
			logger.LogW("found an irregular match: ", match, " skipping...")
			continue
		}

		functionSigsData = append(functionSigsData, match[1])
	}

	functionSigs = &functionSigsData

	return
}

/*
DecodeInputAutomate gets tx data as parameter(+ redis variables) and automate the decoding procedure if we have the key before and returns the decoded values
*/
func DecodeInputAutomate(txData []byte, redisURL string, redisPassword string) ([]evmStructs.DecodeOutput, error) {

	//Create the redis client
	dbHandler, redisErr := database.NewDatabaseHandler("REDIS", redisURL)
	if redisErr != nil {
		return nil, redisErr
	}
	dbHandler.SetPassword(redisPassword)

	//Set function name
	fN := "0X" + common.Bytes2Hex(txData[0:4])
	dbHandler.SetTable("FN_" + fN)

	var signatureInfo []schemas.SignatureInformation

	getDataErr := dbHandler.GetData(&signatureInfo, -1, "", "")
	if getDataErr == redis.Nil {
		return nil, errors.New("function name is not available in redis")
	} else if getDataErr != nil {
		return nil, getDataErr
	}

	if len(signatureInfo) > 1 {
		return nil, errors.New("hash collision exists for the given function name")
	}

	decodedData, decodeErr := DecodeInput(txData, signatureInfo[1].SignatureParams)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return decodedData, nil

}
