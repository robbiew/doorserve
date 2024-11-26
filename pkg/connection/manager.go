package connection

import "sync"

var (
	connections []*Wrapper
	mutex       sync.Mutex
)

// AddConnection adds a wrapper to the connection list.
func AddConnection(wrapper *Wrapper) {
	mutex.Lock()
	defer mutex.Unlock()
	connections = append(connections, wrapper)
}

// RemoveConnection removes a wrapper from the connection list.
func RemoveConnection(wrapper *Wrapper) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, conn := range connections {
		if conn == wrapper {
			connections = append(connections[:i], connections[i+1:]...)
			return
		}
	}
}
