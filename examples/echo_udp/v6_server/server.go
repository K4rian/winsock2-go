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

var (
	serverAddr         string = "::1"
	serverPort         int    = 70 // Gopher ;)
	flagStopServer     int32  = 0
	flagStopHandleData int32  = 0
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

// setNonBlockingSocket sets the given socket to non-blocking mode.
// Returns an error if setting the non-blocking mode fails, otherwise returns nil.
func setNonBlockingSocket(socket ws2.SOCKET) error {
	var iMode uint32 = 1
	ret, err := ws2.IoctlSocket(socket, ws2.FIONBIO, &iMode)
	if ret == ws2.SOCKET_ERROR {
		return fmt.Errorf("failed to set non-blocking mode on socket: %v", err)
	}
	return nil
}

// consoleLog formats and prints a log message to the console in the specified format.
func consoleLog(format string, v ...any) {
	log.Printf("udp_server_v6 | %s\n", fmt.Sprintf(format, v...))
}

func main() {
	if !wsStartup() {
		return
	}
	defer ws2.WSACleanup()

	// When a SIGINT or SIGTERM signal is received, flagStopHandleData is set to 1,
	// signaling the server to stop gracefully.
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-chSig
		atomic.StoreInt32(&flagStopHandleData, 1)
	}()

	consoleLog("Starting UDP server...")

	// Create a UDP socket.
	serverSocket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_DGRAM, ws2.IPPROTO_UDP)
	if serverSocket == ws2.INVALID_SOCKET {
		consoleLog("Failed to create server socket: %v\n", err)
		return
	}
	// Shutdown the server socket to disable sending and receiving data
	// and close it to free up resources.
	defer ws2.Shutdown(serverSocket, ws2.SD_BOTH)
	defer ws2.CloseSocket(serverSocket)

	// Set the server socket to non-blocking mode.
	if err := setNonBlockingSocket(serverSocket); err != nil {
		consoleLog("%v", err)
	}

	// Create a new socket address.
	addr, addrLen := ws2.NewSockAddress(serverAddr, serverPort)
	/* OR */
	// addr, addrLen := ws2.NewV6SockAddress(serverAddr, serverPort)

	// Bind it.
	ret, err := ws2.Bind(serverSocket, unsafe.Pointer(addr), addrLen)
	if ret != 0 {
		consoleLog("Bind failed: %v", err)
		return
	}
	consoleLog("Bound on %s:%d", serverAddr, serverPort)

	// Handle data in a separate goroutine.
	go handleData(serverSocket)

	// Wait till the SIGINT or SIGTERM signal is received.
	for atomic.LoadInt32(&flagStopServer) == 0 {
		<-time.After(10 * time.Millisecond)
	}
	consoleLog("Server stopped.")
}

func handleData(socket ws2.SOCKET) {
	defer atomic.StoreInt32(&flagStopServer, 1)

	// Set-up a buffer to receive the client message.
	buf := make([]byte, 1024)

	for atomic.LoadInt32(&flagStopHandleData) == 0 {
		// Define a new empty client socket address.
		clientAddr := ws2.SockAddr{}
		clientAddrLen := int32(unsafe.Sizeof(clientAddr))

		// Receive the client message...
		n, _ := ws2.RecvFrom(
			socket,                      // Server socket.
			unsafe.Pointer(&buf[0]),     // Pointer to the first byte in the buffer.
			int32(len(buf)),             // Buffer length.
			0,                           // Optional flags.
			unsafe.Pointer(&clientAddr), // Pointer to the client socket address struct.
			&clientAddrLen,              // Client socket address struct length.
		)

		if n > 0 {
			clientIP, clientPort := clientAddr.ToIPPort()
			consoleLog("Received client message: '%s' from %s:%d", string(buf[:n]), clientIP, clientPort)

			// ... then send it back.
			ret, err := ws2.SendTo(
				socket,                      // Server socket.
				unsafe.Pointer(&buf[0]),     // Pointer to the first byte in the buffer.
				int32(len(buf)),             // Buffer length.
				0,                           // Optional flags.
				unsafe.Pointer(&clientAddr), // Pointer to the client socket address struct.
				clientAddrLen,               // Client socket address struct length.
			)
			if ret == ws2.SOCKET_ERROR {
				consoleLog("Failed to send back the message to the client: %v", err)
			}
		}
	}
}
