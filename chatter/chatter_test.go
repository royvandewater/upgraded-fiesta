package chatter_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	chatter "github.com/royvandewater/upgraded-fiesta/chatter"
)

var interval = 10 * time.Millisecond

var _ = Describe("Chatter", func() {
	Describe("NewNode", func() {
		var node chatter.Node
		BeforeEach(func() {
			c := make(chan string)
			node = chatter.NewNode(c, 1*time.Second)
		})

		It("Should exist", func() {
			Expect(node).NotTo(BeNil())
		})
	})

	Describe("When a node is instantiated", func() {
		var channel chan string
		var node chatter.Node

		BeforeEach(func() {
			channel = make(chan string)
			node = chatter.NewNode(channel, interval)
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
