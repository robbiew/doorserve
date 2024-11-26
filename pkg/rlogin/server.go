package rlogin

import (
	"log"
	"net"
	"strconv"

	"github.com/robbiew/go-doorserver/pkg/connection"
)

// StartServer starts an RLOGIN server on the specified port.
func StartServer(port int, isDebug bool) {
	address := ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server on port %d: %v", port, err)
	}
	log.Printf("RLOGIN server listening on port %d (Debug: %v)", port, isDebug)

	nodeCounter := 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn, nodeCounter, isDebug)
		nodeCounter++
	}
}

func handleConnection(conn net.Conn, node int, isDebug bool) {
	wrapper := connection.NewWrapper(conn, node, isDebug)

	if isDebug {
		wrapper.SetModule(connection.NewDebugModule())
	} else {
		wrapper.SetModule(connection.NewMenuModule())
	}

	wrapper.HandleConnection()
}
