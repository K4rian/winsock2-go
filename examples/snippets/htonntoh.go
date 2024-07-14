package snippets

import (
	"log"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- Htond:  Converts a double (float64) from host to TCP/IP network byte order (which is big-endian).
- Htonf:  Converts a float from host to TCP/IP network byte order (which is big-endian).
- Htonl:  Converts a unsigned long (uint32) from host to TCP/IP network byte order (which is big-endian).
- Htonll: Converts an unsigned int64 (uint64) from host to TCP/IP network byte order (which is big-endian).
- Htons:  Converts a unsigned short (uint16) from host to TCP/IP network byte order (which is big-endian).
- Ntohl:  Converts a unsigned long (uint32) from TCP/IP network order to host byte order (which is little-endian on Intel processors).
- Ntohs:  Converts a unsigned short (uint16) from TCP/IP network byte order to host byte order (which is little-endian on Intel processors).
*/

func ExampleHtond() {
	val := float64(145197.2024)
	log.Printf("ExampleHtond  | %.4f -> %d\n", val, ws2.Htond(val))
}

func ExampleHtonf() {
	val := float32(451.8)
	log.Printf("ExampleHtonf  | %.1f -> %d\n", val, ws2.Htonf(val))
}

func ExampleHtonl() {
	val := uint32(27015)
	log.Printf("ExampleHtonl  | %d -> %d\n", val, ws2.Htonl(val))
}

func ExampleHtonll() {
	val := uint64(2701573903245214522)
	log.Printf("ExampleHtonll | %d -> %d\n", val, ws2.Htonll(val))
}

func ExampleHtons() {
	val := uint16(4046)
	log.Printf("ExampleHtons  | %d -> %d\n", val, ws2.Htons(val))
}

func ExampleNtohl() {
	val := uint32(2271805440)
	log.Printf("ExampleNtohl  | %d -> %d\n", val, ws2.Ntohl(val))
}

func ExampleNtohs() {
	val := uint16(52751)
	log.Printf("ExampleNtohs  | %d -> %d\n", val, ws2.Ntohs(val))
}
