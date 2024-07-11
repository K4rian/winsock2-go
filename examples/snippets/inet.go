package snippets

import (
	"encoding/binary"
	"log"
	"syscall"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- InetAddr:  Converts a string containing an IPv4 dotted-decimal address into a proper address for the InAddr structure.
- InetNtoa:  Converts an (IPv4) Internet network address into an ASCII string in Internet standard dotted-decimal format.
- InetNtop:  Converts an IPv4 or IPv6 Internet network address into a string in Internet standard format (ANSI).
- InetNtopW: Converts an IPv4 or IPv6 Internet network address into a string in Internet standard format (Unicode).
- InetPton:  Converts an IPv4 or IPv6 Internet network address in its standard text presentation form into its numeric binary form (ANSI).
- InetPtonW: Converts an IPv4 or IPv6 Internet network address in its standard text presentation form into its numeric binary form (Unicode).
*/

func ExampleInetAddr() {
	ip := append([]byte("192.168.10.100"), 0)
	addr := ws2.InetAddr(&ip[0])

	if addr == ws2.INADDR_NONE {
		log.Printf("InetAddr | Invalid IP address: %s\n", ip)
		return
	}

	addrBE := ws2.Ntohl(addr)
	addrBytesLE := make([]byte, 4)
	addrBytesBE := make([]byte, 4)

	binary.LittleEndian.PutUint32(addrBytesLE, addr)
	binary.BigEndian.PutUint32(addrBytesBE, addr)

	log.Printf("InetAddr | Original IP (string): %s\n", ip)
	log.Printf("InetAddr | Binary (LE):          %d\n", addr)
	log.Printf("InetAddr | Binary (BE):          %d\n", addrBE)
	log.Printf("InetAddr | Bytes (LE):           %v\n", addrBytesLE)
	log.Printf("InetAddr | Bytes (BE):           %v\n", addrBytesBE)
}

func ExampleInetNtoa() {
	addrStr := "7143:15c7:b1c3:d2d0:5864:431b:1835:24d2"
	addr := ws2.InAddr{
		Addr: [4]byte{192, 168, 10, 100},
	}

	retStr := ws2.InetNtoa(&addr)

	log.Printf("InetNtoa | Original Addr (bytes):    %v\n", addr.Addr)
	log.Printf("InetNtoa | Expected Result (string): %s\n", addrStr)
	log.Printf("InetNtoa | Result (string):          %s\n", retStr)
}

func ExampleInetNtop() {
	addrStr := "7143:15c7:b1c3:d2d0:5864:431b:1835:24d2"
	addr := &ws2.InAddr6{
		Addr: [16]byte{113, 67, 21, 199, 177, 195, 210, 208, 88, 100, 67, 27, 24, 53, 36, 210},
	}

	bufLen := int32(128)
	buf := make([]byte, bufLen)
	ret, err := ws2.InetNtop(ws2.AF_INET6, unsafe.Pointer(addr), &buf[0], bufLen)

	if ret == nil {
		log.Printf("InetNtop | The function returns a nil value: %v\n", err)
	}

	// buf could also be used here.
	retStr := ws2.BytePtrToString(ret)

	log.Printf("InetNtop | Original IP (bytes):      %v\n", addr.Addr)
	log.Printf("InetNtop | Expected Result (string): %s\n", addrStr)
	log.Printf("InetNtop | Result (string):          %s\n", retStr)
}

func ExampleInetNtopW() {
	addrStr := "7143:15c7:b1c3:d2d0:5864:431b:1835:24d2"
	addr := &ws2.InAddr6{
		Addr: [16]byte{113, 67, 21, 199, 177, 195, 210, 208, 88, 100, 67, 27, 24, 53, 36, 210},
	}

	bufLen := int32(128)
	buf := make([]uint16, bufLen)
	ret, err := ws2.InetNtopW(ws2.AF_INET6, unsafe.Pointer(addr), &buf[0], bufLen)

	if ret == nil {
		log.Printf("InetNtopW | The function returns a nil value: %v\n", err)
	}

	// buf could also be used here.
	retStr, _ := ws2.UTF16PtrToString(ret)

	log.Printf("InetNtopW | Original IP (bytes):      %v\n", addr.Addr)
	log.Printf("InetNtopW | Expected Result (string): %s\n", addrStr)
	log.Printf("InetNtopW | Result (string):          %s\n", retStr)
}

func ExampleInetPton() {
	family := ws2.AF_INET
	addrStr := "192.168.1.100"
	addrBytes := append([]byte(addrStr), 0)
	addrBuf := ws2.InAddr{}
	exResult := "[192 168 1 100]"

	ret, err := ws2.InetPton(uint16(family), &addrBytes[0], unsafe.Pointer(&addrBuf))
	if err != nil {
		log.Printf("InetPton | Error (%d): %v\n", ret, err)
		return
	}

	log.Printf("InetPton | Original IP (string):    %s\n", addrStr)
	log.Printf("InetPton | Expected Result (bytes): %s\n", exResult)
	log.Printf("InetPton | Result (bytes):          %v\n", addrBuf.Addr)
}

func ExampleInetPtonW() {
	family := ws2.AF_INET6
	addrStr := "3942:5c49:086a:4db7:eaf6:8784:1bad:6cc8"
	addr16, _ := syscall.UTF16FromString(addrStr)
	addrBuf := ws2.InAddr6{}
	exResult := "[57 66 92 73 8 106 77 183 234 246 135 132 27 173 108 200]"

	ret, err := ws2.InetPtonW(uint16(family), &addr16[0], unsafe.Pointer(&addrBuf))
	if err != nil {
		log.Printf("InetPtonW | Error (%d): %v\n", ret, err)
		return
	}

	log.Printf("InetPtonW | Original IP (string):    %s\n", addrStr)
	log.Printf("InetPtonW | Expected Result (bytes): %s\n", exResult)
	log.Printf("InetPtonW | Result (bytes):          %v\n", addrBuf.Addr)
}
