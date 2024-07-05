package ws2

import (
	"encoding/binary"
	"math"
	"net"
)

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
