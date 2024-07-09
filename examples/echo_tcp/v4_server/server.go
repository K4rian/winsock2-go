package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

var (
	serverAddr            string          = "0.0.0.0"
	serverPort            int             = 70 // Gopher ;)
	clients               int32           = 0
	flagStopHandleClients int32           = 0
	chanShutdown          chan struct{}   = make(chan struct{})
	wg                    *sync.WaitGroup = &sync.WaitGroup{}
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
	log.Printf("tcp_server_v4 | %s\n", fmt.Sprintf(format, v...))
}

func main() {
	if !wsStartup() {
		return
	}
	defer ws2.WSACleanup()

	// When a SIGINT or SIGTERM signal is received, flagStop is set to 1,
	// signaling the server to stop gracefully.
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-chSig

		if atomic.LoadInt32(&clients) == 0 {
			close(chanShutdown)
		} else {
			atomic.StoreInt32(&flagStopHandleClients, 1)
		}
	}()

	consoleLog("Starting TCP server...")

	// Create a TCP socket.
	serverSocket, err := ws2.Socket(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	if serverSocket == ws2.INVALID_SOCKET {
		consoleLog("Failed to create server socket: %v", err)
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
	// addr, addrLen := ws2.NewV4SockAddress(serverAddr, serverPort)

	// Bind it.
	ret, err := ws2.Bind(serverSocket, unsafe.Pointer(addr), addrLen)
	if ret != 0 {
		consoleLog("Bind failed: %v", err)
		return
	}

	// Listen for incoming connections with a maximum backlog size.
	ret, err = ws2.Listen(serverSocket, ws2.SOMAXCONN)
	if ret != 0 {
		consoleLog("Listen failed: %v", err)
		return
	}
	consoleLog("Listening on %s:%d", serverAddr, serverPort)

	// Start a goroutine to accept incoming connections.
	go func() {
		for {
			select {
			case <-chanShutdown:
				return
			default:
				// Define a new empty client socket address.
				clientAddr := ws2.SockAddr{}
				clientAddrLen := int32(unsafe.Sizeof(clientAddr))

				// Attempt to accept a new client connection.
				clientSocket, _ := ws2.Accept(serverSocket, unsafe.Pointer(&clientAddr), &clientAddrLen)
				if clientSocket == ws2.INVALID_SOCKET {
					continue
				}

				// Use the select function to wait for the client socket to become readable
				// within the specified timeout period (1s).
				readfds := ws2.FDSet{}
				timeout := ws2.Timeval{Sec: 1, USec: 0}
				readfds.Set(clientSocket)

				ret, err := ws2.Select(0, &readfds, nil, nil, &timeout)
				if ret == ws2.SOCKET_ERROR {
					consoleLog("Select error: %v", err)
					ws2.Shutdown(clientSocket, ws2.SD_BOTH)
					ws2.CloseSocket(clientSocket)
					continue
				}

				// Set the client socket to non-blocking mode.
				if err := setNonBlockingSocket(clientSocket); err != nil {
					consoleLog("%v", err)
					ws2.Shutdown(clientSocket, ws2.SD_BOTH)
					ws2.CloseSocket(clientSocket)
					continue
				}

				// Handle the client connection in a separate goroutine.
				wg.Add(1)
				atomic.AddInt32(&clients, 1)
				go handleClient(clientSocket)
			}
		}
	}()

	// Wait for shutdown signal.
	<-chanShutdown

	// Wait for all client goroutines to finish.
	wg.Wait()

	consoleLog("Server stopped.")
}

func handleClient(socket ws2.SOCKET) {
	defer ws2.Shutdown(socket, ws2.SD_BOTH)
	defer ws2.CloseSocket(socket)

	defer atomic.AddInt32(&clients, -1)
	defer wg.Done()

	// Get Client socket address.
	addr := ws2.SockAddrInet4{}
	addrLen := int32(unsafe.Sizeof(addr))
	ret, err := ws2.GetPeerName(socket, unsafe.Pointer(&addr), addrLen)
	if ret != 0 {
		consoleLog("Failed to get client's peer name (%d): %v", ret, err)
		return
	}

	// Announce the new client.
	clientIP, clientPort := addr.ToIPPort()
	consoleLog("Client connected from: %s:%d", clientIP, clientPort)

	// Get client local socket name.
	localAddr := ws2.SockAddrInet4{}
	localAddrLen := int32(unsafe.Sizeof(localAddr))
	ret, err = ws2.GetSockName(socket, unsafe.Pointer(&localAddr), &localAddrLen)
	if ret != 0 {
		consoleLog("Failed to get local client socket name (%d): %v", ret, err)
	}
	laIP, laPort := localAddr.ToIPPort()
	consoleLog("Client socket local address: %s:%d", laIP, laPort)

	// Get SO_KEEPALIVE socket option.
	optVal := int32(0)
	optLen := int32(4) // 4 Bytes
	ret, err = ws2.GetSockOpt(socket, ws2.SOL_SOCKET, ws2.SO_KEEPALIVE, unsafe.Pointer(&optVal), &optLen)
	if ret == ws2.SOCKET_ERROR {
		consoleLog("Unable to get client's SO_KEEPALIVE socket option (%d): %v", ret, err)
		return
	}

	// Set SO_KEEPALIVE socket option.
	if optVal == 0 {
		optVal = 1
		ret, err = ws2.SetSockOpt(socket, ws2.SOL_SOCKET, ws2.SO_KEEPALIVE, unsafe.Pointer(&optVal), optLen)
		if ret == ws2.SOCKET_ERROR {
			consoleLog("Unable to set client's SO_KEEPALIVE socket option (%d): %v", ret, err)
			return
		}
	}

	// Set-up a buffer to receive the client message.
	buf := make([]byte, 1024)

	for atomic.LoadInt32(&flagStopHandleClients) == 0 {
		n, _ := ws2.Recv(
			socket,                  // Client socket.
			unsafe.Pointer(&buf[0]), // Pointer to the first byte in the buffer.
			int32(len(buf)),         // Buffer length.
			0,                       // Optional flags.
		)

		if n > 0 {
			consoleLog("Received client message: '%s'", string(buf[:n]))

			// ... then send it back.
			ret, err := ws2.Send(
				socket,                  // Client socket.
				unsafe.Pointer(&buf[0]), // Pointer to the first byte in the buffer.
				int32(len(buf)),         // Buffer length.
				0,                       // Optional flags.
			)
			if ret == ws2.SOCKET_ERROR {
				consoleLog("Failed to send back the message to the client: %v", err)
			}
			break
		}
	}
}
