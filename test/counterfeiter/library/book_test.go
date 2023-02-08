package library_test

import (
	"counterfeiter/library"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var fake = &library.BookService{}

var _ = Describe("Test functions",
	func() {
		Context("sample", func() {

			It("example", func() {

			})
		})
	},
)
