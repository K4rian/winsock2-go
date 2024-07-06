package ws2

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"syscall"
	"unsafe"
)

// BytePtrToString converts a byte pointer to a string.
//
// It iterates over the memory pointed to by the byte pointer until it encounters a null terminator (0 byte),
// appending each byte to a byte slice. It then constructs a string from the byte slice and returns it.
func BytePtrToString(bptr *byte) string {
	var dataSlice []byte
	ptr := uintptr(unsafe.Pointer(bptr))
	for {
		b := *(*byte)(unsafe.Pointer(ptr))
		if b == 0 {
			break
		}
		dataSlice = append(dataSlice, b)
		ptr++
	}
	return string(dataSlice)
}

// GUIDToString returns the string representation of the given GUID.
func GUIDToString(guid GUID) string {
	return fmt.Sprintf("{%08X-%04X-%04X-%04X-%012X}", guid.Data1, guid.Data2, guid.Data3, guid.Data4[:2], guid.Data4[2:])
}

// Htond converts a float64 to network byte order (big-endian) as a uint64.
func Htond(value float64) uint64 {
	bits := math.Float64bits(value)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, bits)
	return binary.BigEndian.Uint64(buf)
}

// Htonf converts a float32 to network byte order (big-endian) as a uint32.
func Htonf(value float32) uint32 {
	bits := math.Float32bits(value)
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, bits)
	return binary.BigEndian.Uint32(buf)
}

// Htonl converts a unsigned long (uint32) from host to TCP/IP network byte order (which is big-endian).
func Htonl(hostlong uint32) uint32 {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, hostlong)
	return binary.BigEndian.Uint32(b)
}

// Htonll converts a uint64 to network byte order (big-endian).
func Htonll(value uint64) uint64 {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	return uint64(binary.BigEndian.Uint64(buf))
}

// Htons converts a unsigned short (uint16) from host to TCP/IP network byte order (which is big-endian).
func Htons(hostshort uint16) uint16 {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, hostshort)
	return binary.BigEndian.Uint16(b)
}

// InetNtoa converts an (IPv4) Internet network address into an ASCII string in Internet standard dotted-decimal format.
func InetNtoa(addr *InAddr) string {
	ip := net.IPv4(addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3])
	return ip.String()
}

// Ntohl converts a unsigned long (uint32) from TCP/IP network order to host byte order (which is little-endian on Intel processors).
func Ntohl(netlong uint32) uint32 {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, netlong)
	return binary.LittleEndian.Uint32(b)
}

// Ntohs converts a unsigned short (uint16) from TCP/IP network byte order to host byte order (which is little-endian on Intel processors).
func Ntohs(netshort uint16) uint16 {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, netshort)
	return binary.LittleEndian.Uint16(b)
}

// NewSockAddress creates a new generic socket address structure.
//
// It takes an IPv4 or IPv6 address string and a port number as input.
//
// Returns the SockAddr structure and its length or nil if the IP is invalid.
func NewSockAddress(address string, port int) (addr *SockAddr, length int32) {
	ip := net.ParseIP(address)
	if ip != nil {
		sockAddr := SockAddr{}
		if sockAddr.FromNetIP(ip, port) == nil {
			addr = &sockAddr
			length = int32(unsafe.Sizeof(sockAddr))
		}
	}
	return
}

// NewV4SockAddress creates a new IPv4 socket address structure.
//
// It takes an IPv4 address string and a port number as input.
//
// Returns the SockAddrInet4 structure and its length or nil if the IP is invalid.
func NewV4SockAddress(address string, port int) (addr *SockAddrInet4, length int32) {
	ip := net.ParseIP(address)
	if ip != nil {
		if v4 := ip.To4(); v4 != nil {
			v4Addr := SockAddrInet4{
				Family: AF_INET,
				Port:   Htons(uint16(port)),
				Addr:   [4]byte(v4[:4]),
			}
			addr = &v4Addr
			length = int32(unsafe.Sizeof(v4Addr))
		}
	}
	return
}

// NewV6SockAddress creates a new IPv6 socket address structure.
//
// It takes an IPv6 address string and a port number as input.
//
// Returns the SockAddrInet6 structure and its length or nil if the IP is invalid.
func NewV6SockAddress(address string, port int) (addr *SockAddrInet6, length int32) {
	ip := net.ParseIP(address)
	if ip != nil {
		if v6 := ip.To16(); v6 != nil {
			v6Addr := SockAddrInet6{
				Family: AF_INET6,
				Port:   Htons(uint16(port)),
				Addr:   [16]byte(v6[:16]),
			}
			addr = &v6Addr
			length = int32(unsafe.Sizeof(v6Addr))
		}
	}
	return
}

// PtrToSockAddr converts an pointer to a SockAddr.
//
// Returns the SockAddr and nil if valid, otherwise returns an error.
func PtrToSockAddr(p unsafe.Pointer) (*SockAddr, error) {
	if p != nil {
		addr := *(*SockAddr)(p)
		if addr.Family == AF_INET || addr.Family == AF_INET6 {
			return &addr, nil
		}
	}
	return nil, fmt.Errorf("invalid pointer")
}

// UTF16PtrToString converts a uint16 pointer to a string.
//
// Returns the populated string and nil if valid, otherwise returns a empty string and an error.
func UTF16PtrToString(ptr *uint16) (string, error) {
	if ptr != nil {
		sl := (*[1 << 30]uint16)(unsafe.Pointer(ptr))[:]

		var length int
		for sl[length] != 0 {
			length++
		}

		str := syscall.UTF16ToString(sl[:length])
		return str, nil
	}
	return "", fmt.Errorf("invalid pointer")
}
