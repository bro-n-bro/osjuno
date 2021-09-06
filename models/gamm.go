package models

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"unicode"
)

type ChainedSwap struct {
	Sender     string
	PoolsSwaps []PoolSwap
}

type PoolSwap struct {
	PoolId  uint64
	InCoin  sdk.Coin
	OutCoin sdk.Coin
}

func NewPoolSwap(poolId uint64, inCoin sdk.Coin, outCoin sdk.Coin) PoolSwap {
	return PoolSwap{
		PoolId:  poolId,
		InCoin:  inCoin,
		OutCoin: outCoin,
	}
}

func NewCoinFromStr(fromStr string) (*sdk.Coin, error) {
	for charIndex, char := range fromStr {
		if !unicode.IsNumber(char) {
			denom := fromStr[charIndex:]
			amount, ok := sdk.NewIntFromString(fromStr[:charIndex])
			if !ok {
				return nil, errors.New("can't parse coin amount")
			}
			coint := sdk.NewCoin(denom, amount)
			return &coint, nil
		}
	}
	return nil, errors.New("cant parse denom-amount string")
}
