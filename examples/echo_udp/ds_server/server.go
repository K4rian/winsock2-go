package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

var (
	serverPort         int   = 70 // Gopher ;)
	flagStopServer     int32 = 0
	flagStopHandleData int32 = 0
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

// enableDualStackSocket enables dual-stack mode on the specified socket, allowing it to handle both IPv4 and IPv6.
func enableDualStackSocket(socket ws2.SOCKET) error {
	var opt uint32 = 0
	var optLen int32 = 4 // DWORD - 4 bytes

	ret, err := ws2.SetSockOpt(socket, ws2.IPPROTO_IPV6, ws2.IPV6_V6ONLY, unsafe.Pointer(&opt), optLen)
	if ret == ws2.SOCKET_ERROR {
		return fmt.Errorf("failed to enable dual-stack mode on socket: %v", err)
	}
	return nil
}

// consoleLog formats and prints a log message to the console in the specified format.
func consoleLog(format string, v ...any) {
	log.Printf("udp_ds_server | %s\n", fmt.Sprintf(format, v...))
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

	// Address information hints for resolving the socket address.
	hints := ws2.AddrInfoW{
		Family:   ws2.AF_UNSPEC,   // Either IPv4 or IPv6.
		SockType: ws2.SOCK_DGRAM,  // Datagram-oriented.
		Protocol: ws2.IPPROTO_UDP, // UDP protocol.
		Flags:    ws2.AI_PASSIVE,  // The socket address will be used in a call to bind() for a listening socket.
	}
	result := &ws2.AddrInfoW{}
	portStr, _ := syscall.UTF16FromString(strconv.Itoa(serverPort))

	// Resolve the server address and port.
	ret, err := ws2.GetAddrInfoW(nil, &portStr[0], &hints, &result)
	if ret != 0 {
		consoleLog("GetAddrInfoW failed (%d): %v", ret, err)
		return
	}
	defer ws2.FreeAddrInfoW(result)

	consoleLog("Starting dual-stack UDP server...")

	// Create a UDP socket.
	serverSocket, err := ws2.Socket(result.Family, result.SockType, result.Protocol)
	if serverSocket == ws2.INVALID_SOCKET {
		consoleLog("Failed to create server socket: %v", err)
		return
	}
	// Shutdown the server socket to disable sending and receiving data
	// and close it to free up resources.
	defer ws2.Shutdown(serverSocket, ws2.SD_BOTH)
	defer ws2.CloseSocket(serverSocket)

	// Set the IPV6_V6ONLY socket option to zero to enable dual-stack.
	if err := enableDualStackSocket(serverSocket); err != nil {
		consoleLog("%v", err)
		return
	}

	// Set the server socket to non-blocking mode.
	if err := setNonBlockingSocket(serverSocket); err != nil {
		consoleLog("%v", err)
		return
	}

	// Bind it.
	ret, err = ws2.Bind(serverSocket, result.Addr, int32(result.AddrLength))
	if ret != 0 {
		consoleLog("Bind failed: %v", err)
		return
	}

	serverAddr := ""
	socketAddr, _ := ws2.PtrToSockAddr(result.Addr)
	if socketAddr != nil {
		serverAddr = socketAddr.ToIP().String()
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
			consoleLog("Received client message: '%s' from %s:%d", string(buf[:n]), clientIP.String(), clientPort)

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
