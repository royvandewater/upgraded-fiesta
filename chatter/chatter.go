package chatter

import "time"

// Node is a single member of the fiesta. In order to have
// a real fiesta, you must have more than one node.
type Node interface {
	Run()
	Stop()
}

// NewNode constructs a new instance that implements Node
func NewNode(channel chan string, interval time.Duration) Node {
	return &_Node{channel: channel, interval: interval}
}

type _Node struct {
	channel  chan string
	interval time.Duration
}

func (n *_Node) Run() {
	n.channel <- "Hello world!"
}
func (n *_Node) Stop() {}
