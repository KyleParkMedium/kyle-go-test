package main

import "math/big"

type InitTransaction struct {
	from string
	to   string
	// amount uint64
	amount *big.Int
}

func (it *InitTransaction) Key() string {
	return "init"
}

// func Key() string {
// 	return "init"
// }

func one() int {
	return 1
}
func two() int {
	return 2
}
func name() string {
	return "name"
}

type TransferTransaction struct {
	from string
	// toAmount map[string]uint64
	toAmount map[string]*big.Int
}

type CoinTransactionResult struct {
	AddAmountResult map[string]*big.Int
	SetAmountResult map[string]*big.Int
	FromAmount      *big.Int
}

// func (tt *TransferTransaction) transfer(result *cointransaction.CoinTransactionResult, from string) uint64 {
func (tt *TransferTransaction) transfer(result *CoinTransactionResult, from string) *big.Int {

	// var totalAmount uint64
	totalAmount := new(big.Int)
	for to, amount := range tt.toAmount {
		// result.AddAmountResult[to] += amount
		// totalAmount += amount
		result.AddAmountResult[to].Add(result.AddAmountResult[to], amount)
		totalAmount.Add(totalAmount, amount)
	}

	return totalAmount
}
