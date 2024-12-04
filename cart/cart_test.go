package cart_test

import (
	"go_sandbox/cart"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {
	Context("initially", func() {
		It("has 0 items", func() {
			cart := cart.Item{ID: "1", Name: "Bread", Price: 1.00, Qty: 1}
		})
		It("has 0 units", func() {})
		Specify("the total amount is 0.00", func() {})
	})
})
