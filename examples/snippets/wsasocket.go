package snippets

import (
	"bytes"
	"log"
	"os"
	"time"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSADuplicateSocketA:    Returns a WSAProtocolInfoA structure that can be used to create a new socket descriptor for a shared socket.
- WSADuplicateSocketW:    Returns a WSAProtocolInfoA structure that can be used to create a new socket descriptor for a shared socket.
- WSAGetOverlappedResult: Retrieves the results of an overlapped operation on the specified socket.
- WSAIoctl:               Controls the mode of a socket.
- WSASocketA:             Creates a socket that is bound to a specific transport-service provider.
- WSASocketW:             Creates a socket that is bound to a specific transport-service provider.
*/

// WSAStartup() has to be called before using these functions.

func ExampleWSADuplicateSocketA() {
	// Create the "original" socket.
	socket, err := ws2.WSASocketA(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP, nil, 0, ws2.WSA_FLAG_OVERLAPPED)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSADuplicateSocketA | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)
	defer log.Printf("ExampleWSADuplicateSocketA | Original socket %d closed\n", socket)

	log.Printf("ExampleWSADuplicateSocketA | Original socket: %d\n", socket)

	// Get the protocol info of the original socket.
	processID := uint32(os.Getpid())
	protoInfo := &ws2.WSAProtocolInfoA{}
	success, err := ws2.WSADuplicateSocketA(socket, processID, protoInfo)
	if !success {
		log.Printf("ExampleWSADuplicateSocketA | Failed to duplicate original socket info: %v\n", err)
		return
	}

	// Duplicate the original socket.
	cloneSocket, err := ws2.WSASocketA(protoInfo.AddressFamily, protoInfo.SocketType, protoInfo.Protocol, protoInfo, 0, 0)
	if cloneSocket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSADuplicateSocketA | Failed to create duplicate socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(cloneSocket)
	defer log.Printf("ExampleWSADuplicateSocketA | Duplicated socket %d closed\n", cloneSocket)

	log.Printf("ExampleWSADuplicateSocketA | Duplicated socket: %d\n", cloneSocket)
}

func ExampleWSADuplicateSocketW() {
	// Create the "original" socket.
	socket, err := ws2.WSASocketW(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP, nil, 0, 0)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSADuplicateSocketW | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)
	defer log.Printf("ExampleWSADuplicateSocketW | Original socket %d closed\n", socket)

	log.Printf("ExampleWSADuplicateSocketW | Original socket: %d\n", socket)

	// Get the protocol info of the original socket.
	processID := uint32(os.Getpid())
	protoInfo := &ws2.WSAProtocolInfoW{}
	success, err := ws2.WSADuplicateSocketW(socket, processID, protoInfo)
	if !success {
		log.Printf("ExampleWSADuplicateSocketW | Failed to duplicate original socket info: %v\n", err)
		return
	}

	// Duplicate the original socket.
	cloneSocket, err := ws2.WSASocketW(protoInfo.AddressFamily, protoInfo.SocketType, protoInfo.Protocol, protoInfo, 0, 0)
	if cloneSocket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSADuplicateSocketW | Failed to create duplicate socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(cloneSocket)
	defer log.Printf("ExampleWSADuplicateSocketW | Duplicated socket %d closed\n", cloneSocket)

	log.Printf("ExampleWSADuplicateSocketW | Duplicated socket: %d\n", cloneSocket)
}

func ExampleWSAGetOverlappedResult() {
	// Create a dummy TCP socket.
	socket, err := ws2.Socket(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSAGetOverlappedResult | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)

	// Get the overlapped result.
	overlapped := ws2.WSAOverlapped{}
	transfert := uint32(0)
	wait := true
	flags := uint32(0)
	success, err := ws2.WSAGetOverlappedResult(socket, &overlapped, &transfert, wait, &flags)
	if !success {
		log.Printf("ExampleWSAGetOverlappedResult | Error: %v\n", err)
		return
	}

	log.Printf("ExampleWSAGetOverlappedResult | Result (raw): %v\n", overlapped)
	log.Printf("ExampleWSAGetOverlappedResult | Transfert:    %d\n", transfert)
	log.Printf("ExampleWSAGetOverlappedResult | Wait:         %v\n", wait)
	log.Printf("ExampleWSAGetOverlappedResult | Flags:        %v\n", flags)
}

func ExampleWSAIoctl() {
	/*
		Note: You must run this code with elevated privileges (Administrator) to use RAW sockets on Windows.
	*/
	const SIO_RCVALL = 0x98000001

	// Create a raw socket to capture all network packets.
	socket, err := ws2.Socket(ws2.AF_INET, ws2.SOCK_RAW, ws2.IPPROTO_IP)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSAIoctl | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.Shutdown(socket, ws2.SD_BOTH)
	defer ws2.CloseSocket(socket)
	defer log.Printf("ExampleWSAIoctl | Socket closed\n")

	// Bind it to the local IP.
	sockAddr, sockAddrLen := ws2.NewV4SockAddress("192.168.1.1", 0) // Use your local IP address here!
	ret, err := ws2.Bind(socket, unsafe.Pointer(sockAddr), sockAddrLen)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAIoctl | Socket bind failed: %v\n", err)
		return
	}

	// Set the socket to non-blocking mode.
	iMode := uint32(1)
	ret, err = ws2.IoctlSocket(socket, ws2.FIONBIO, &iMode)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAIoctl | Failed to set non-blocking mode on socket: %v\n", err)
		return
	}

	// Enable promiscuous mode.
	ioControlCode := uint32(SIO_RCVALL)
	ioValue := uint32(1)
	ioValueSize := uint32(4) // DWORD
	bytesReturned := uint32(0)
	ret, err = ws2.WSAIoctl(
		socket,
		ioControlCode,
		unsafe.Pointer(&ioValue),
		ioValueSize,
		nil,
		0,
		&bytesReturned,
		nil,
		nil,
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAIoctl | Failed to set promiscuous mode on socket: %v\n", err)
		return
	}

	log.Printf("ExampleWSAIoctl | Socket set to promiscuous mode\n")

	// Let's receive some data with a execution time limit of 10s.
	curTime := time.Now()
	endTime := curTime.Add(10 * time.Second)
	recvBuf := make([]byte, 65536)
	recvBufLen := int32(65536)

	for curTime.Before(endTime) {
		ret, _ = ws2.Recv(socket, unsafe.Pointer(&recvBuf[0]), recvBufLen, 0)

		// Let's skip the socket errors.
		if ret != ws2.SOCKET_ERROR {
			log.Printf("ExampleWSAIoctl | Received raw data (%d bytes): %v\n", ret, bytes.TrimRight(recvBuf, "\x00"))
		}

		// Sleep for a bit then refresh the current time.
		<-time.After(10 * time.Millisecond)
		curTime = time.Now()
	}
}

func ExampleWSASocketA() {
	socket, err := ws2.WSASocketA(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP, nil, 0, ws2.WSA_FLAG_OVERLAPPED)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSASocketA | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)
	defer log.Printf("ExampleWSASocketA | Socket %d closed\n", socket)

	log.Printf("ExampleWSASocketA | Socket %d created successfully\n", socket)
}

func ExampleWSASocketW() {
	socket, err := ws2.WSASocketW(ws2.AF_INET6, ws2.SOCK_STREAM, ws2.IPPROTO_TCP, nil, 0, 0)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleWSASocketW | Failed to create socket: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)
	defer log.Printf("ExampleWSASocketW | Socket %d closed\n", socket)

	log.Printf("ExampleWSASocketW | Socket %d created successfully\n", socket)
}
