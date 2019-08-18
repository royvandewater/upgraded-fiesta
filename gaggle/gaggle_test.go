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
			var conn1 gaggle.Connection
			var conn2 gaggle.Connection

			BeforeEach(func() {
				conn1 = sut.NewConnection()
				conn2 = sut.NewConnection()
			})

			Describe("When a message is emitted on one channel", func() {
				BeforeEach(func() {
					conn1.Output() <- "Hello world!"
				})

				It("Should come out the other channel", func() {
					Eventually(conn2.Input).Should(Receive(Equal("Hello world!")))
				})
			})
		})

		Describe("When three connection have been created", func() {
			var conn1 gaggle.Connection
			var conn2 gaggle.Connection
			var conn3 gaggle.Connection

			BeforeEach(func() {
				conn1 = sut.NewConnection()
				conn2 = sut.NewConnection()
				conn3 = sut.NewConnection()
			})

			Describe("When the third connection is closed", func() {
				BeforeEach(func() {
					conn3.Close()
				})

				Describe("When a message is emitted on connection 1", func() {
					BeforeEach(func() {
						conn1.Output() <- "Hello world!"
					})

					It("Should be emitted from connection 2", func() {
						Eventually(conn2.Input()).Should(Receive(Equal("Hello world!")))
					})
				})
			})
		})
	})
})
