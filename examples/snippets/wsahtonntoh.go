package snippets

import (
	"log"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSAHtonl: Converts a unsigned long (int32) from host byte order to network byte order.
- WSAHtons: Converts a unsigned short (int16) from host byte order to network byte order.
- WSANtohl: Converts a unsigned long (uint32) from network byte order to host byte order.
- WSANtohs: Converts a unsigned short (uint16) from network byte order to host byte order.
*/

// WSAStartup() has to be called before using these functions.

func ExampleWSAHtonl() {
	socket, _ := ws2.Socket(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	valIn := uint32(27015)
	valOut := uint32(0)
	ws2.WSAHtonl(socket, valIn, &valOut)

	log.Printf("WSAHtonl | %d -> %d\n", valIn, valOut)
}

func ExampleWSAHtons() {
	socket, _ := ws2.Socket(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	valIn := uint16(4046)
	valOut := uint16(0)
	ws2.WSAHtons(socket, valIn, &valOut)

	log.Printf("WSAHtons | %d -> %d\n", valIn, valOut)
}

func ExampleWSANtohl() {
	socket, _ := ws2.Socket(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	valIn := uint32(2271805440)
	valOut := uint32(0)
	ws2.WSANtohl(socket, valIn, &valOut)

	log.Printf("WSANtohl | %d -> %d\n", valIn, valOut)
}

func ExampleWSANtohs() {
	socket, _ := ws2.Socket(ws2.AF_INET, ws2.SOCK_STREAM, ws2.IPPROTO_TCP)
	valIn := uint16(52751)
	valOut := uint16(0)
	ws2.WSANtohs(socket, valIn, &valOut)

	log.Printf("WSANtohs | %d -> %d\n", valIn, valOut)
}
