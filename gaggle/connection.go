package gaggle

type _Connection struct {
	input  chan string
	output chan string
	close  chan bool
}

func (c *_Connection) Input() <-chan string {
	return c.input
}

func (c *_Connection) Output() chan<- string {
	return c.output
}

func (c *_Connection) Close() {
	close(c.output)
	c.close <- true
}