package evmStructs

import (
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"regexp"
	"solity/utils/logger"
	"strings"
)

type evmSignature struct {
	Signature    string
	Hash         string
	Types        []string
	IndexedTypes []string
}

type SignatureKeeper struct {
	hashList map[string]evmSignature
}

func standardizeSignature(input string) (stdSignature string, types []string, indexedTypes []string, err error) {
	// Check if the signature has the necessary symbols in it
	rgx, cmplErr := regexp.Compile(`(^[^\s()]+)\(([^()]+)\)$`)

	if cmplErr != nil {
		err = cmplErr
		return
	}

	matches := rgx.FindStringSubmatch(input)

	if matches == nil || len(matches) <= 2 {
		err = errors.New("incorrect signature: " + input)
		return
	}

	params := strings.Split(matches[2], ",")
	stdSignature = matches[1] + "("

	for _, param := range params {
		trimmedParam := strings.Trim(param, " ")

		splittedParam := strings.Split(trimmedParam, " ")

		switch len(splittedParam) {
		case 1:
			// Case 1: "uint256"
			types = append(types, splittedParam[0])
			stdSignature = stdSignature + splittedParam[0] + ","
			continue
		case 2:
			// Case 2.1: "uint256 baseFeeL1"
			if !(strings.ToLower(splittedParam[0]) == "indexed") && !(strings.ToLower(splittedParam[1]) == "indexed") {
				types = append(types, splittedParam[0])
				stdSignature = stdSignature + splittedParam[0] + ","
				continue
			}

			// Case 2.2: "indexed uint256"
			if strings.ToLower(splittedParam[0]) == "indexed" {
				indexedTypes = append(indexedTypes, splittedParam[1])
				stdSignature = stdSignature + splittedParam[1] + ","
				continue
			}

			// Case 2.3: "uint256 indexed"
			if strings.ToLower(splittedParam[1]) == "indexed" {
				indexedTypes = append(indexedTypes, splittedParam[0])
				stdSignature = stdSignature + splittedParam[0] + ","
				continue
			}

			err = errors.New("incorrect function params in: " + trimmedParam)
			return

		case 3:
			// Case 3.1: "indexed uint256 messageIndex"
			if strings.ToLower(splittedParam[0]) == "indexed" {
				indexedTypes = append(indexedTypes, splittedParam[1])
				stdSignature = stdSignature + splittedParam[1] + ","
				continue
			}

			// Case 3.2: "uint256 indexed messageIndex"
			if strings.ToLower(splittedParam[1]) == "indexed" {
				indexedTypes = append(indexedTypes, splittedParam[0])
				stdSignature = stdSignature + splittedParam[0] + ","
				continue
			}

			err = errors.New("incorrect function params in: " + trimmedParam)
			return

		default:
			err = errors.New("incorrect function params in: " + trimmedParam)
			return
		}

	}

	stdSignature = stdSignature[:len(stdSignature)-1] + ")"

	return
}

func NewSignatureKeeper(inputs ...string) (keeper SignatureKeeper) {
	// Initialize
	keeper = SignatureKeeper{hashList: map[string]evmSignature{}}

	for _, signature := range inputs {
		keeper.AddSignature(signature)
	}

	return
}

func (sK *SignatureKeeper) AddSignature(signature string) {
	// Standardize
	standartSignature, inputTypes, indexedTypes, sErr := standardizeSignature(signature)

	if sErr != nil {
		logger.LogW(sErr)
		return
	}

	// Take the keccak hash
	signatureHash := strings.ToUpper(crypto.Keccak256Hash([]byte(standartSignature)).Hex())

	// Create the siganture object
	evmSig := evmSignature{
		Signature:    standartSignature,
		Hash:         signatureHash,
		Types:        inputTypes,
		IndexedTypes: indexedTypes,
	}

	sK.hashList[signatureHash] = evmSig
}

func (sK *SignatureKeeper) AddHash(hash string, types []string, indexedTypes []string) {

	signatureHash := strings.ToUpper(hash)

	// Create the siganture object
	evmSig := evmSignature{
		Signature:    "",
		Hash:         signatureHash,
		Types:        types,
		IndexedTypes: indexedTypes,
	}

	sK.hashList[signatureHash] = evmSig
}

func (sK *SignatureKeeper) GetHash(hash string) (signature string, theHash string, types []string, indexedTypes []string, err error) {
	// Standardize
	signatureHash := strings.ToUpper(hash)

	value := sK.hashList[signatureHash]

	if value.Hash != "" {
		signature = value.Signature
		theHash = value.Hash
		types = value.Types
		indexedTypes = value.IndexedTypes

		return
	}

	err = errors.New("no data found for the hash: " + hash)

	return
}

func (sK *SignatureKeeper) PrintAllSignatures() {
	for k, v := range sK.hashList {
		logger.LogD("Hash: ", k)
		logger.LogD("Corresponding Data: ", v)
	}
}
