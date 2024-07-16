package snippets

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSAConnect:        Establishes a connection to another socket application, exchanges connect data, and specifies required quality of service based on the specified FlowSpec structure.
- WSAConnectByNameA: Establishes a connection to a specified host and port.
- WSAConnectByNameW: Establishes a connection to a specified host and port.
*/

// WSAStartup() has to be called before using these functions.

func ExampleWSAConnect() {
	// Let's assume we have a local TCP v6 server running on port 70.

	// Create a TCP socket.
	socket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSAConnect | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.Shutdown(socket, ws2.SD_BOTH)
	defer ws2.CloseSocket(socket)

	// Define the service address.
	name, nameLen := ws2.NewV6SockAddress("::1", 70)
	callerBuf := &ws2.WSABuf{}
	calleeBuf := &ws2.WSABuf{}
	sQOS := &ws2.QOS{}

	// Connect to the service.
	ret, err := ws2.WSAConnect(
		socket,
		unsafe.Pointer(name),
		nameLen,
		callerBuf,
		calleeBuf,
		sQOS,
		nil, // Should be nil
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAConnect | Error: %v\n", err)
		return
	}
	defer log.Printf("ExampleWSAConnect | Disconnected\n")

	remoteIP, remotePort := name.ToIPPort()
	log.Printf("ExampleWSAConnect | Connected to %s:%d\n", remoteIP, remotePort)
}

func ExampleWSAConnectByNameA() {
	// Let's assume we have a local TCP v6 server running on port 70.

	// Create a TCP socket.
	socket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSAConnectByNameA | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.Shutdown(socket, ws2.SD_BOTH)
	defer ws2.CloseSocket(socket)

	// Define the node and service name.
	nodename := append([]byte("::1"), 0)
	servicename := append([]byte("gopher"), 0) // 70
	localAddress := ws2.SockAddrInet6{}
	localAddressLen := uint32(unsafe.Sizeof(localAddress))
	remoteAddress := ws2.SockAddrInet6{}
	remoteAddressLen := uint32(0)
	timeout := ws2.Timeval{Sec: 5, USec: 0}

	// Connect to the service.
	success, err := ws2.WSAConnectByNameA(
		socket,
		&nodename[0],
		&servicename[0],
		&localAddressLen,
		unsafe.Pointer(&localAddress),
		&remoteAddressLen,
		unsafe.Pointer(&remoteAddress),
		&timeout,
		nil,
	)
	if !success {
		log.Printf("ExampleWSAConnectByNameA | Error: %v\n", err)
		return
	}
	defer log.Printf("ExampleWSAConnectByNameA | Disconnected\n")

	remoteIP, remotePort := remoteAddress.ToIPPort()
	log.Printf("ExampleWSAConnectByNameA | Connected to %s:%d\n", remoteIP, remotePort)
}

func ExampleWSAConnectByNameW() {
	// Let's assume we have a local TCP v6 server running on port 70.

	// Create a TCP socket.
	socket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSAConnectByNameW | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.Shutdown(socket, ws2.SD_BOTH)
	defer ws2.CloseSocket(socket)

	// Define the node and service name.
	nodename, _ := syscall.UTF16PtrFromString("::1")
	servicename, _ := syscall.UTF16PtrFromString("gopher")
	localAddress := ws2.SockAddrInet6{}
	localAddressLen := uint32(unsafe.Sizeof(localAddress))
	remoteAddress := ws2.SockAddrInet6{}
	remoteAddressLen := uint32(0)
	timeout := ws2.Timeval{Sec: 5, USec: 0}

	// Connect to the service.
	success, err := ws2.WSAConnectByNameW(
		socket,
		nodename,
		servicename,
		&localAddressLen,
		unsafe.Pointer(&localAddress),
		&remoteAddressLen,
		unsafe.Pointer(&remoteAddress),
		&timeout,
		nil,
	)
	if !success {
		log.Printf("ExampleWSAConnectByNameW | Error: %v\n", err)
		return
	}
	defer log.Printf("ExampleWSAConnectByNameW | Disconnected\n")

	remoteIP, remotePort := remoteAddress.ToIPPort()
	log.Printf("ExampleWSAConnectByNameW | Connected to %s:%d\n", remoteIP, remotePort)
}
