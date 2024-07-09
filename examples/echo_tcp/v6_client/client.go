package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

const MAX_RECV_ERRORS = 10 // Max Recv errors threshold.

var (
	serverAddr string = "::1"
	serverPort int    = 70 // Gopher ;)
	flagStop   int32  = 0
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
	log.Printf("tcp_client_v6 | %s\n", fmt.Sprintf(format, v...))
}

func main() {
	if !wsStartup() {
		return
	}
	defer ws2.WSACleanup()

	// When a SIGINT or SIGTERM signal is received, flagStop is set to 1,
	// signaling the client to stop gracefully.
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-chSig
		atomic.StoreInt32(&flagStop, 1)
	}()

	// Create a TCP socket.
	clientSocket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
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
	// addr, addrLen := ws2.NewV6SockAddress(serverAddr, serverPort)

	// Connect to the server.
	ret, err := ws2.Connect(clientSocket, unsafe.Pointer(addr), addrLen)
	if err != nil {
		consoleLog("Connect error (%d): %v", ret, err)
		return
	}
	consoleLog("Connected to %s:%d", serverAddr, serverPort)

	// Prepare the message.
	message := []byte("Hello from the v6 tcp client!")

	// Send a message to the server.
	consoleLog("Sending the message: '%s'...", string(message[:]))
	ret, err = ws2.Send(
		clientSocket,                // Client socket.
		unsafe.Pointer(&message[0]), // Pointer to the first byte of the message.
		int32(len(message)),         // Message length.
		0,                           // Optional flags.
	)
	if ret == ws2.SOCKET_ERROR {
		consoleLog("Failed to send message to the server: %v", err)
		return
	}
	consoleLog("Message sent, waiting for response...")

	// Set-up a buffer to receive the message back from the server.
	buf := make([]byte, 1024)

	// Recv errors counter.
	recvErrors := 0

	// Receive the server message.
	for atomic.LoadInt32(&flagStop) == 0 {
		n, err := ws2.Recv(
			clientSocket,            // Client socket.
			unsafe.Pointer(&buf[0]), // Pointer to the first byte in the buffer.
			int32(len(buf)),         // Buffer length.
			0,                       // Optional flags.
		)

		if n < 0 || err != nil {
			recvErrors++

			if recvErrors >= MAX_RECV_ERRORS {
				consoleLog("Recv failed after 10 attempts: %v", err)
				break
			}
			<-time.After(5 * time.Millisecond)
		} else if n > 0 {
			consoleLog("Received message: '%s'", string(buf[:]))
			break
		}
	}
}
