package coordination_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	coordination "github.com/royvandewater/upgraded-fiesta/coordination"
)

var interval = 10 * time.Millisecond

var _ = Describe("Coordination", func() {
	Describe("When a node is instantiated", func() {
		var channel chan string
		var node coordination.Node

		BeforeEach(func() {
			channel = make(chan string)
			node = coordination.NewNode(channel, interval)
			go node.Run()
		})

		AfterEach(func() {
			node.Stop()
		})

		Describe("And twice the maximum interval period has passed", func() {
			BeforeEach(func() {
				<-time.After(2 * interval)
			})

			It("should broadcast", func() {
				Eventually(channel).Should(Receive(Equal("Hello world!")))
			})
		})
	})
})
