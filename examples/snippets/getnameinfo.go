package snippets

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- GetNameInfoA: Converts a socket address into a node name and service name using ANSI encoding.
- GetNameInfoW: Converts a socket address into a node name and service name using Unicode encoding.
*/

func ExampleGetNameInfoA(sock6 *ws2.SockAddrInet6) {
	var nodeBuffer [ws2.NI_MAXHOST]byte
	var serviceBuffer [ws2.NI_MAXSERV]byte
	var flags int32 = 0

	ret, err := ws2.GetNameInfoA(
		unsafe.Pointer(sock6),
		int32(unsafe.Sizeof(*sock6)),
		&nodeBuffer[0],
		int32(ws2.NI_MAXHOST),
		&serviceBuffer[0],
		int32(ws2.NI_MAXSERV),
		flags,
	)

	if err != nil {
		log.Printf("ExampleGetNameInfoA | Error (%d): %v\n", ret, err)
		return
	}

	node := ws2.BytePtrToString(&nodeBuffer[0])
	service := ws2.BytePtrToString(&serviceBuffer[0])

	log.Printf("ExampleGetNameInfoA | Node:    %s\n", node)
	log.Printf("ExampleGetNameInfoA | Service: %s\n", service)
}

func ExampleGetNameInfoW(sock6 *ws2.SockAddrInet6) {
	var nodeBuffer [ws2.NI_MAXHOST]uint16
	var serviceBuffer [ws2.NI_MAXSERV]uint16
	var flags int32 = 0

	ret, err := ws2.GetNameInfoW(
		unsafe.Pointer(sock6),
		int32(unsafe.Sizeof(*sock6)),
		&nodeBuffer[0],
		int32(ws2.NI_MAXHOST),
		&serviceBuffer[0],
		int32(ws2.NI_MAXSERV),
		flags,
	)

	if err != nil {
		log.Printf("ExampleGetNameInfoW | Error (%d): %v\n", ret, err)
		return
	}

	node := syscall.UTF16ToString(nodeBuffer[:])
	service := syscall.UTF16ToString(serviceBuffer[:])

	log.Printf("ExampleGetNameInfoW | Node:    %s\n", node)
	log.Printf("ExampleGetNameInfoW | Service: %s\n", service)
}
