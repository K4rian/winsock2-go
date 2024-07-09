package snippets

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- GetHostByAddr:
- GetHostByName:
- GetHostNameA:
- GetHostNameW:
*/

func ExampleGetHostByAddr() {
	addr := ws2.InAddr{Addr: [4]byte{140, 82, 113, 4}}
	addrLen := unsafe.Sizeof(addr)
	htype := ws2.AF_INET
	host, err := ws2.GetHostByAddr(unsafe.Pointer(&addr), int32(addrLen), int32(htype))

	if host == nil {
		log.Printf("GetHostByAddr | Error: %v\n", err)
		return
	}

	log.Printf("GetHostByAddr | Name:         %s\n", ws2.BytePtrToString(host.HName))
	log.Printf("GetHostByAddr | Address Type: %d\n", host.HAddrType)
	log.Printf("GetHostByAddr | Length:       %d\n", host.HLength)
}

func ExampleGetHostByName() {
	name := []byte("github.com")
	host, err := ws2.GetHostByName(&name[0])

	if host == nil {
		log.Printf("GetHostByAddr | Error: %v\n", err)
		return
	}

	log.Printf("GetHostByName | Name:         %s\n", ws2.BytePtrToString(host.HName))
	log.Printf("GetHostByName | Address Type: %d\n", host.HAddrType)
	log.Printf("GetHostByName | Length:       %d\n", host.HLength)
}

func ExampleGetHostNameA() {
	buf := [512]byte{}
	bufLen := 512

	ws2.GetHostNameA(&buf[0], bufLen)
	log.Printf("GetHostNameA | Name: %s\n", string(buf[:]))
}

func ExampleGetHostNameW() {
	buf := [512]uint16{}
	bufLen := 512

	ws2.GetHostNameW(&buf[0], bufLen)
	log.Printf("GetHostNameW | Name: %s\n", syscall.UTF16ToString(buf[:]))
}
