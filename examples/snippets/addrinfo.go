package snippets

import (
	"log"
	"syscall"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- GetAddrInfoA:  Resolves a hostname and service name into a list of socket addresses using ANSI encoding.
- GetAddrInfoW:  Resolves a hostname and service name into a list of socket addresses using Unicode encoding.
- FreeAddrInfoA: Frees memory allocated by the GetAddrInfoA function.
- FreeAddrInfoW: Frees memory allocated by the GetAddrInfoW function.
*/

func ExampleGetAddrInfoA() {
	domain := "github.com"
	service := "https"

	hints := ws2.AddrInfoA{
		Flags:    ws2.AI_CANONNAME,
		Family:   ws2.AF_UNSPEC,
		SockType: ws2.SOCK_STREAM,
		Protocol: ws2.IPPROTO_TCP,
	}
	result := &ws2.AddrInfoA{}

	nodeName := append([]byte(domain), 0)
	serviceName := append([]byte(service), 0)

	ret, err := ws2.GetAddrInfoA(&nodeName[0], &serviceName[0], &hints, &result)
	if err != nil {
		log.Printf("GetAddrInfoA | Failed (%d): %v\n", ret, err)
		return
	}
	defer ws2.FreeAddrInfoA(result)

	for addr := result; addr != nil; addr = addr.Next {
		log.Printf("GetAddrInfoA | Flags: %d\n", addr.Flags)
		log.Printf("GetAddrInfoA | Family: %s\n", addrFamilyToString(addr.Family))        // _common.go#addrFamilyToString
		log.Printf("GetAddrInfoA | Socket Type: %s\n", socketTypeToString(addr.SockType)) // _common.go#socketTypeToString
		log.Printf("GetAddrInfoA | Protocol: %s\n", ipProtocolToString(addr.Protocol))    // _common.go#ipProtocolToString
		log.Printf("GetAddrInfoA | Address Length: %d\n", addr.AddrLength)

		if addr.CanonName != nil {
			canonicalName := ws2.BytePtrToString(addr.CanonName)
			if canonicalName != "" {
				log.Printf("GetAddrInfoA | Canonical Name: %s\n", canonicalName)
			}
		}

		sa, err := ws2.PtrToSockAddr(addr.Addr)
		if err != nil {
			log.Printf("GetAddrInfoA | Failed to cast the socket address struct: %v\n", err)
			return
		}
		ip, port := sa.ToIPPort()
		log.Printf("GetAddrInfoA | Address: %s:%d\n", ip, port)
	}
}

func ExampleGetAddrInfoW() {
	domain := "github.com"
	service := "https"

	hints := ws2.AddrInfoW{
		Flags:    ws2.AI_CANONNAME,
		Family:   ws2.AF_UNSPEC,
		SockType: ws2.SOCK_STREAM,
		Protocol: ws2.IPPROTO_TCP,
	}
	result := &ws2.AddrInfoW{}

	nodeName, _ := syscall.UTF16PtrFromString(domain)
	serviceName, _ := syscall.UTF16PtrFromString(service)

	ret, err := ws2.GetAddrInfoW(nodeName, serviceName, &hints, &result)
	if err != nil {
		log.Printf("GetAddrInfoW | Failed (%d): %v\n", ret, err)
		return
	}
	defer ws2.FreeAddrInfoW(result)

	for addr := result; addr != nil; addr = addr.Next {
		log.Printf("GetAddrInfoW | Flags: %d\n", addr.Flags)
		log.Printf("GetAddrInfoW | Family: %s\n", addrFamilyToString(addr.Family))        // _common.go#addrFamilyToString
		log.Printf("GetAddrInfoW | Socket Type: %s\n", socketTypeToString(addr.SockType)) // _common.go#socketTypeToString
		log.Printf("GetAddrInfoW | Protocol: %s\n", ipProtocolToString(addr.Protocol))    // _common.go#ipProtocolToString
		log.Printf("GetAddrInfoW | Address Length: %d\n", addr.AddrLength)

		if addr.CanonName != nil {
			canonicalName, err := ws2.UTF16PtrToString(addr.CanonName)
			if err == nil {
				log.Printf("GetAddrInfoW | Canonical Name: %s\n", canonicalName)
			}
		}

		sa, err := ws2.PtrToSockAddr(addr.Addr)
		if err != nil {
			log.Printf("GetAddrInfoW | Failed to cast the socket address struct: %v\n", err)
			return
		}
		ip, port := sa.ToIPPort()
		log.Printf("GetAddrInfoW | Address: %s:%d\n", ip, port)
	}
}
