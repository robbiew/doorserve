package connection

import (
	"log"
	"strings"
)

// DebugModule handles debug connections
type DebugModule struct{}

// NewDebugModule creates a new DebugModule instance
func NewDebugModule() *DebugModule {
	return &DebugModule{}
}

// HandleInput processes input from the debug client
func (m *DebugModule) HandleInput(wrapper *Wrapper, input []byte) {
	cmd := strings.TrimSpace(strings.ToLower(string(input))) // Normalize command

	switch cmd {
	case "m":
		m.monitorConnections(wrapper)
	case "r":
		m.runDoor(wrapper)
	case "s":
		m.setUsername(wrapper)
	case "d":
		m.disconnect(wrapper)
	case "h":
		m.showHelp(wrapper)
	default:
		m.invalidCommand(wrapper)
	}
}

func (m *DebugModule) monitorConnections(wrapper *Wrapper) {
	log.Println("Monitor connections")
	wrapper.conn.Write([]byte("Monitoring connections is not yet implemented.\r\n"))
}

func (m *DebugModule) runDoor(wrapper *Wrapper) {
	log.Println("Run a door")
	wrapper.conn.Write([]byte("Running a door is not yet implemented.\r\n"))
}

func (m *DebugModule) setUsername(wrapper *Wrapper) {
	log.Println("Set username")
	wrapper.conn.Write([]byte("Enter new username:\r\n"))
	// Assuming input is captured in a subsequent read
}

func (m *DebugModule) disconnect(wrapper *Wrapper) {
	log.Println("Disconnecting debug connection")
	wrapper.conn.Write([]byte("Goodbye!\r\n"))
	wrapper.conn.Close() // Close the connection gracefully
}

func (m *DebugModule) showHelp(wrapper *Wrapper) {
	helpMessage := `Available commands:
  m - Monitor connections
  r - Run a door
  s - Set username
  d - Disconnect
  h - Show this help message
`
	wrapper.conn.Write([]byte(strings.ReplaceAll(helpMessage, "\n", "\r\n")))
}

func (m *DebugModule) invalidCommand(wrapper *Wrapper) {
	wrapper.conn.Write([]byte("Invalid command. Type 'h' for help.\r\n"))
}
