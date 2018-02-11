package token_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/relnod/alp/token"
)

func TestToken(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token Suite")
}

var _ = Describe("token.Type", func() {
	Describe("String()", func() {
		It("converts to a string if type is valid", func() {
			t := token.Int
			Expect(t.String()).To(Equal("Int"))
		})

		It("doesn't convert correctly for internal type", func() {
			t := token.Type(3)
			Expect(t.String()).To(Equal("Unkown Token: 3"))
		})

		It("doesn't convert correctly for undefined type", func() {
			t := token.Type(255)
			Expect(t.String()).To(Equal("Unkown Token: 255"))
		})
	})
})
