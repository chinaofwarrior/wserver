package connection

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// StartTCPServer start a TCP server that listens for incoming connections
func StartTCPServer(address string, commandChan chan<- Command) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error listening on %s: %v", address, err)
	}
	defer listener.Close()
	log.Printf("TCP server listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleTCPConn(conn, commandChan)

	}

}

// handleTCPConn handles individual TCP connections
func handleTCPConn(conn net.Conn, commandChan chan<- Command) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		message = strings.TrimSpace(message)

		// 解析消息和命令码
		parts := strings.Split(message, "*")
		if len(parts) >= 4 {
			commandChan <- Command{
				DeviceID: parts[1],
				Type:     parts[2],
				Params:   parts[3],
			}
		} else {
			log.Print("Invalid message format")
		}
	}
}
