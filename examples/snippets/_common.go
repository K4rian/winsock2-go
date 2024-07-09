package snippets

import (
	"fmt"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

func doubleBytePtrToStrSlice(b **byte) (ret []string) {
	for i := 0; *b != nil; i++ {
		str := ws2.BytePtrToString(*b)
		if len(str) > 0 {
			ret = append(ret, str)
		}
		b = (**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + unsafe.Sizeof(*b)))
	}
	return
}

func addrFamilyToString(family int32) (ret string) {
	switch family {
	case ws2.AF_UNSPEC:
		ret = "Unspecified"
	case ws2.AF_INET:
		ret = "AF_INET (IPv4)"
	case ws2.AF_INET6:
		ret = "AF_INET6 (IPv6)"
	case ws2.AF_NETBIOS:
		ret = "AF_NETBIOS (NetBIOS)"
	default:
		ret = fmt.Sprintf("Other (%d)", family)
	}
	return
}

func socketTypeToString(stype int32) (ret string) {
	switch stype {
	case 0:
		ret = "Unspecified"
	case ws2.SOCK_STREAM:
		ret = "SOCK_STREAM (stream)"
	case ws2.SOCK_DGRAM:
		ret = "SOCK_DGRAM (datagram)"
	case ws2.SOCK_RAW:
		ret = "SOCK_RAW (raw)"
	case ws2.SOCK_RDM:
		ret = "SOCK_RDM (reliable message datagram)"
	case ws2.SOCK_SEQPACKET:
		ret = "SOCK_SEQPACKET (pseudo-stream packet)"
	default:
		ret = fmt.Sprintf("Other (%d)", stype)
	}
	return
}

func ipProtocolToString(proto int32) (ret string) {
	switch proto {
	case 0:
		ret = "Unspecified"
	case ws2.IPPROTO_TCP:
		ret = "IPPROTO_TCP (TCP)"
	case ws2.IPPROTO_UDP:
		ret = "IPPROTO_UDP (UDP)"
	default:
		ret = fmt.Sprintf("Other (%d)", proto)
	}
	return
}
