package gaggle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/royvandewater/upgraded-fiesta/gaggle"
)

var _ = Describe("Gaggle", func() {

	Describe("New", func() {
		var sut gaggle.Gaggle

		BeforeEach(func() {
			sut = gaggle.New()
		})

		It("Should not be nil", func() {
			Expect(sut).NotTo(BeNil())
		})
	})

	Describe("Given a Gaggle", func() {
		var sut gaggle.Gaggle

		BeforeEach(func() {
			sut = gaggle.New()
		})

		Describe("When two connection have been created", func() {
			var conn1 *gaggle.Connection
			var conn2 *gaggle.Connection

			BeforeEach(func() {
				conn1 = sut.NewConnection()
				conn2 = sut.NewConnection()
			})

			Describe("When a message is emitted on one channel", func() {
				BeforeEach(func() {
					conn1.Output <- "Hello world!"
				})

				It("Should come out the other channel", func() {
					Eventually(conn2.Input).Should(Receive(Equal("Hello world!")))
				})
			})
		})
	})
})
