package connection

import "log"

type DebugModule struct{}

func NewDebugModule() *DebugModule {
	return &DebugModule{}
}

func (m *DebugModule) HandleInput(wrapper *Wrapper, input []byte) {
	cmd := string(input)
	switch cmd {
	case "m":
		log.Println("Monitor connections")
	case "r":
		log.Println("Run a door")
	case "s":
		log.Println("Set username")
	case "d":
		log.Println("Disconnect")
	default:
		wrapper.conn.Write([]byte("Invalid command\n"))
	}
}
