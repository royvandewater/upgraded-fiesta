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
	NewConnection() Connection
}

// New returns a new Gaggle instance
func New() Gaggle {
	return &_Gaggle{
		connections: make([]*_Connection, 0),
	}
}

// Connection contains an Input and an Output. Nodes will receive
// message through their Input channel, and emit messages using
// their Output channel.
type Connection interface {
	// Output returns a read only channel that a node can
	// use to receive messages
	Input() <-chan string
	// Output returns a write only channel that a node can
	// use to send messages
	Output() chan<- string

	// Close marks this connection as no longer in use. A connection
	// cannot be re-opened.
	Close()
}

type _Gaggle struct {
	connections []*_Connection
}

func (g *_Gaggle) NewConnection() Connection {
	c := &_Connection{
		input:  make(chan string),
		output: make(chan string),
		close:  make(chan bool),
	}
	g.connections = append(g.connections, c)

	go g.proxy(c)
	return c
}

func (g *_Gaggle) broadcast(sender *_Connection, message string) {
	for _, c := range g.connections {
		if c == sender {
			continue
		}
		c.input <- message
	}
}

func (g *_Gaggle) close(connection *_Connection) {
	for i, c := range g.connections {
		if c != connection {
			continue
		}

		g.deleteConnectionByIndex(i)
		close(connection.input)
	}
}

func (g *_Gaggle) deleteConnectionByIndex(i int) {
	g.connections[i] = g.connections[len(g.connections)-1]
	g.connections[len(g.connections)-1] = nil
	g.connections = g.connections[:len(g.connections)-1]
}

func (g *_Gaggle) proxy(connection *_Connection) {
	for {
		select {
		case message := <-connection.output:
			g.broadcast(connection, message)
		case <-connection.close:
			g.close(connection)
			return
		}
	}
}
