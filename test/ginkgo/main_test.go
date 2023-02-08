package main

import (
	"math/big"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var (
	it *InitTransaction
	tt *TransferTransaction
)

var _ = Describe("Test functions",
	func() {

		value := big.NewInt(100)
		it = &InitTransaction{from: "A", to: "B", amount: value}

		to := make(map[string]*big.Int)
		to["A"] = value

		tt = &TransferTransaction{from: "A", toAmount: to}

		from := "A"

		result := &CoinTransactionResult{AddAmountResult: to, SetAmountResult: to, FromAmount: value}

		Context("sample", func() {
			It("one should return 1", func() {
				Expect(big.NewInt(100)).To(Equal(tt.transfer(result, from)))
			})

			It("one should return 1", func() {
				Expect("init").To(Equal(it.Key()))
			})

			It("one should return 1", func() {
				Expect(1).To(Equal(one()))
			})
			It("name should return name", func() {
				Expect("name").To(Equal(name()))
			})
		})
	},
)
