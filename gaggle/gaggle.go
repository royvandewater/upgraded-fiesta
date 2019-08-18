package gaggle

// Gaggle describes the constraints and provides
// the means with which Nodes can coordinate. The
// intention is that each Node is given their own
// node channel. That channel is intended to be used
// for both sending and receiving messages. Simplex
// communication is enforced by the Gaggle.
type Gaggle interface {
	// NewConnection returns a connection intended to be consumed
	// by a single node. This connection should be a node's only
	// way of communicating with other members of the gaggle.
	NewConnection() *Connection
}

// Connection contains an Input and an Output. Nodes will receive
// message through their Input channel, and emit messages using
// their Output channel.
type Connection struct {
	Input  chan string
	Output chan string
}

// New returns a new Gaggle instance
func New() Gaggle {
	return &_Gaggle{
		connections: make([]*Connection, 0),
	}
}

type _Gaggle struct {
	connections []*Connection
}

func (g *_Gaggle) NewConnection() *Connection {
	c := &Connection{
		Input:  make(chan string),
		Output: make(chan string),
	}
	g.connections = append(g.connections, c)

	go g.proxy(c)
	return c
}

func (g *_Gaggle) broadcast(sender *Connection, message string) {
	for _, c := range g.connections {
		if c == sender {
			continue
		}
		c.Input <- message
	}
}

func (g *_Gaggle) proxy(connection *Connection) {
	for message := range connection.Output {
		g.broadcast(connection, message)
	}
}
