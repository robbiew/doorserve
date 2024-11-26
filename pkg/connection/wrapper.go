package connection

import (
	"log"
	"net"
)

type Wrapper struct {
	conn     net.Conn
	node     int
	isDebug  bool
	module   Module
	remoteIP string
}

// NewWrapper initializes a new Wrapper.
func NewWrapper(conn net.Conn, node int, isDebug bool) *Wrapper {
	return &Wrapper{
		conn:     conn,
		node:     node,
		isDebug:  isDebug,
		remoteIP: conn.RemoteAddr().String(),
	}
}

// SetModule assigns a module to the wrapper.
func (w *Wrapper) SetModule(module Module) {
	w.module = module
}

// HandleConnection processes incoming data for the connection.
func (w *Wrapper) HandleConnection() {
	defer w.conn.Close()
	log.Printf("Handling connection for node %d (Debug: %v)", w.node, w.isDebug)
	for {
		buf := make([]byte, 1024)
		n, err := w.conn.Read(buf)
		if err != nil {
			log.Printf("Connection closed for node %d", w.node)
			return
		}
		w.module.HandleInput(w, buf[:n])
	}
}
