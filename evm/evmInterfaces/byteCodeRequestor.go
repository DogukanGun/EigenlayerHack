package evmInterfaces

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ByteCodeRequestor interface {
	/*
		CodeAt is a function that returns the byte code stored in the given address
	*/
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
}
