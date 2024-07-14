package snippets

import (
	"log"
	"os"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSADuplicateSocketA: Returns a WSAProtocolInfoA structure that can be used to create a new socket descriptor for a shared socket.
- WSADuplicateSocketW: Returns a WSAProtocolInfoA structure that can be used to create a new socket descriptor for a shared socket.
- WSAGetOverlappedResult: Retrieves the results of an overlapped operation on the specified socket.
- WSASocketA: Creates a socket that is bound to a specific transport-service provider.
- WSASocketW: Creates a socket that is bound to a specific transport-service provider.
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
