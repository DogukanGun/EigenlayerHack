package structs

import "github.com/ethereum/go-ethereum/common"

type EigenlayerPayload struct {
	AvsName         string
	OperatorName    string
	AvsAddress      common.Address
	OperatorAddress common.Address
}
