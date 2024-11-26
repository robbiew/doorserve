package connection

import (
	"log"
	"net"
)

// Module defines the interface for handling input and managing connection logic
type Module interface {
	HandleInput(wrapper *Wrapper, input []byte)
}

// Wrapper manages a single connection and its associated module
type Wrapper struct {
	conn    net.Conn
	node    int
	isDebug bool
	Module  Module
	user    *User // For tracking user-specific information
}

// User represents a connected user's details
type User struct {
	Name string
}

// NewWrapper creates a new Wrapper instance
func NewWrapper(conn net.Conn, node int, isDebug bool) *Wrapper {
	return &Wrapper{
		conn:    conn,
		node:    node,
		isDebug: isDebug,
		user:    &User{Name: "Guest"},
	}
}

// HandleConnection reads input from the connection and passes it to the active module
func (w *Wrapper) HandleConnection() {
	buffer := make([]byte, 1024)
	for {
		n, err := w.conn.Read(buffer)
		if err != nil {
			log.Printf("Connection error on node %d: %v", w.node, err)
			return
		}

		if w.Module != nil {
			w.Module.HandleInput(w, buffer[:n]) // Pass input to the active module
		}
	}
}

// SetModule assigns a module to the wrapper
func (w *Wrapper) SetModule(module Module) {
	w.Module = module
}
