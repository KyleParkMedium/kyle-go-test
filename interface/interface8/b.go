package interface8

import "math/big"

type STO interface {
	init()

	isIssuable() bool
	issue(_tokenHolder string, _value *big.Int, _data []byte) bool
}
