package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

var (
	serverAddr string = "127.0.0.1"
	serverPort int    = 70 // Gopher ;)
)

// wsStartup initializes Winsock and checks if it's successfully initialized.
// Returns true if initialization is successful, otherwise false.
func wsStartup() bool {
	version := uint16(ws2.WINSOCK_VERSION)
	data := &ws2.WSAData{}

	ret, err := ws2.WSAStartup(version, data)
	if err != nil {
		consoleLog("Startup failed (%d): %v", ret, err)
		return false
	}
	return true
}

// consoleLog formats and prints a log message to the console in the specified format.
func consoleLog(format string, v ...any) {
	log.Printf("udp_client_v4 | %s\n", fmt.Sprintf(format, v...))
}

func main() {
	if !wsStartup() {
		return
	}
	defer ws2.WSACleanup()

	// Create a UDP socket.
	clientSocket, err := ws2.Socket(ws2.AF_INET, ws2.SOCK_DGRAM, ws2.IPPROTO_UDP)
	if clientSocket == ws2.INVALID_SOCKET {
		consoleLog("Failed to create client socket: %v", err)
		return
	}
	// Shutdown the client socket to disable sending and receiving data
	// and close it to free up resources.
	defer ws2.Shutdown(clientSocket, ws2.SD_BOTH)
	defer ws2.CloseSocket(clientSocket)

	// Create a new socket address.
	addr, addrLen := ws2.NewSockAddress(serverAddr, serverPort)
	/* OR */
	// addr, addrLen := ws2.NewV4SockAddress(serverAddr, serverPort)

	// Prepare the message.
	message := []byte("Hello from the v4 udp client!")

	// Send the message to the server.
	// Note: SendTo() is used instead of Send because Connect() is not called.
	consoleLog("Sending the message: '%s'...", string(message[:]))
	ret, err := ws2.SendTo(
		clientSocket,                // Client socket.
		unsafe.Pointer(&message[0]), // Pointer to the first character in the message.
		int32(len(message)),         // Message length.
		0,                           // Optional flags.
		unsafe.Pointer(addr),        // Pointer to the server socket address struct.
		addrLen,                     // Server socket address struct length.
	)
	if ret == ws2.SOCKET_ERROR {
		consoleLog("Failed to send message to the server: %v", err)
		return
	}
	consoleLog("Message sent, waiting for response...")

	// Set-up a buffer to receive the message back from the server.
	buf := make([]byte, 1024)

	// Receive the server message.
	n, err := ws2.RecvFrom(
		clientSocket,            // Client socket.
		unsafe.Pointer(&buf[0]), // Pointer to the first byte in the buffer.
		int32(len(buf)),         // Buffer length.
		0,                       // Optional flags.
		unsafe.Pointer(addr),    // Pointer to the server socket address struct.
		&addrLen,                // Server socket address struct length.
	)

	if n < 0 || err != nil {
		consoleLog("RecvFrom failed (%d): %v", n, err)
	} else if n > 0 {
		consoleLog("Received message: '%s'", string(buf[:]))
	}
}
