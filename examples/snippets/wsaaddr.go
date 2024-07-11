package snippets

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSAAddressToStringA: Converts all components of a sockaddr structure into a human-readable string representation of the address.
- WSAAddressToStringW: Converts all components of a sockaddr structure into a human-readable string representation of the address.
- WSAStringToAddressA: Converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure.
- WSAStringToAddressW: Converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure.
*/

// WSAStartup() has to be called before using these functions.

func ExampleWSAAddressToStringA() {
	bufferLen := uint32(16384)
	bufferProto := ws2.WSAProtocolInfoA{}
	// protos := []int32{} // Array of protocols

	ret, err := ws2.WSAEnumProtocolsA( /*&protos[0]*/ nil, &bufferProto, &bufferLen) // nil = all available protocols
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAAddressToStringA | WSAEnumProtocolsA error (%d): %v\n", ret, err)
		return
	}

	addrExpected := "172.20.20.90:4046" // Expected
	addr := ws2.SockAddrInet4{
		Family: ws2.AF_INET,
		Port:   uint16(ws2.Htons(4046)),
		Addr:   [4]byte{172, 20, 20, 90},
	}
	addrLen := uint32(unsafe.Sizeof(addr))
	retBufLen := uint32(256)
	retBuf := make([]byte, retBufLen)

	ret, err = ws2.WSAAddressToStringA(
		unsafe.Pointer(&addr),
		addrLen,
		&bufferProto,
		&retBuf[0],
		&retBufLen,
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAAddressToStringA | Error (%d): %v\n", ret, err)
		return
	}

	retStr := string(retBuf[:])

	log.Printf("ExampleWSAAddressToStringA | Original Addr (bytes):    %v\n", addr.Addr)
	log.Printf("ExampleWSAAddressToStringA | Expected Result (string): %s\n", addrExpected)
	log.Printf("ExampleWSAAddressToStringA | Result (string):          %s\n", retStr)
}

func ExampleWSAAddressToStringW() {
	bufferLen := uint32(16384)
	bufferProto := ws2.WSAProtocolInfoW{}
	// protos := []int32{} // Array of protocols

	ret, err := ws2.WSAEnumProtocolsW( /*&protos[0]*/ nil, &bufferProto, &bufferLen) // nil = all available protocols
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAAddressToStringW | WSAEnumProtocolsW error (%d): %v\n", ret, err)
		return
	}

	addrExpected := "[7143:15c7:b1c3:d2d0:5864:431b:1835:24d2]:27015" // Expected
	addr := ws2.SockAddrInet6{
		Family: ws2.AF_INET6,
		Port:   uint16(ws2.Htons(27015)),
		Addr:   [16]byte{113, 67, 21, 199, 177, 195, 210, 208, 88, 100, 67, 27, 24, 53, 36, 210},
	}
	addrLen := uint32(unsafe.Sizeof(addr))
	retBufLen := uint32(256)
	retBuf := make([]uint16, retBufLen)

	ret, err = ws2.WSAAddressToStringW(
		unsafe.Pointer(&addr),
		addrLen,
		&bufferProto,
		&retBuf[0],
		&retBufLen,
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAAddressToStringW | Error (%d): %v\n", ret, err)
		return
	}

	retStr := syscall.UTF16ToString(retBuf[:])

	log.Printf("ExampleWSAAddressToStringW | Original Addr (bytes):    %v\n", addr.Addr)
	log.Printf("ExampleWSAAddressToStringW | Expected Result (string): %s\n", addrExpected)
	log.Printf("ExampleWSAAddressToStringW | Result (string):          %s\n", retStr)
}

func ExampleWSAStringToAddressA() {
	bufferLen := uint32(16384)
	bufferProto := ws2.WSAProtocolInfoA{}
	// protos := []int32{} // Array of protocols

	ret, err := ws2.WSAEnumProtocolsA( /*&protos[0]*/ nil, &bufferProto, &bufferLen) // nil = all available protocols
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAStringToAddressA | WSAEnumProtocolsA error (%d): %v\n", ret, err)
		return
	}

	addrExpected := "{2 0 [192 168 1 1] [0 0 0 0 0 0 0 0]}" // Expected
	addrStr := "192.168.1.1"                                // v4
	addrBytes := append([]byte(addrStr), 0)                 // v4
	addrFamily := uint16(ws2.AF_INET)                       // v4
	addr := ws2.SockAddrInet4{}                             // v4
	addrLen := int32(unsafe.Sizeof(addr))

	ret, err = ws2.WSAStringToAddressA(
		&addrBytes[0],
		addrFamily,
		&bufferProto,
		unsafe.Pointer(&addr),
		&addrLen,
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAStringToAddressA | Error (%d): %v\n", ret, err)
		return
	}

	log.Printf("ExampleWSAStringToAddressA | Original Addr (bytes):  %v\n", addrBytes)
	log.Printf("ExampleWSAStringToAddressA | Original Addr (string): %s\n", addrBytes)
	log.Printf("ExampleWSAStringToAddressA | Expected Result:        %s\n", addrExpected)
	log.Printf("ExampleWSAStringToAddressA | Result (raw):           %v\n", addr)
}

func ExampleWSAStringToAddressW() {
	bufferLen := uint32(16384)
	bufferProto := ws2.WSAProtocolInfoW{}
	// protos := []int32{} // Array of protocols

	ret, err := ws2.WSAEnumProtocolsW( /*&protos[0]*/ nil, &bufferProto, &bufferLen) // nil = all available protocols
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAStringToAddressW | WSAEnumProtocolsW error (%d): %v\n", ret, err)
		return
	}

	addrExpected := "{23 0 0 [113 67 21 199 177 195 210 208 88 100 67 27 24 53 36 210] 0}" // Expected
	addrStr := "7143:15c7:b1c3:d2d0:5864:431b:1835:24d2"                                   // v6
	addrBytes, _ := syscall.UTF16FromString(addrStr)                                       // v6
	addrFamily := uint16(ws2.AF_INET6)                                                     // v6
	addr := ws2.SockAddrInet6{}                                                            // v6
	addrLen := int32(unsafe.Sizeof(addr))

	ret, err = ws2.WSAStringToAddressW(
		&addrBytes[0],
		addrFamily,
		&bufferProto,
		unsafe.Pointer(&addr),
		&addrLen,
	)
	if ret == ws2.SOCKET_ERROR {
		log.Printf("ExampleWSAStringToAddressW | Error (%d): %v\n", ret, err)
		return
	}

	log.Printf("ExampleWSAStringToAddressW | Original Addr (bytes):  %v\n", addrBytes)
	log.Printf("ExampleWSAStringToAddressW | Original Addr (string): %s\n", addrStr)
	log.Printf("ExampleWSAStringToAddressW | Expected Result:        %s\n", addrExpected)
	log.Printf("ExampleWSAStringToAddressW | Result (raw):           %v\n", addr)
}
