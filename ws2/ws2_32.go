package ws2

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	modws232          = syscall.NewLazyDLL("ws2_32.dll")
	procWSAFDIsSet    = modws232.NewProc("__WSAFDIsSet")  //
	procAccept        = modws232.NewProc("accept")        //
	procBind          = modws232.NewProc("bind")          //
	procClosesocket   = modws232.NewProc("closesocket")   //
	procConnect       = modws232.NewProc("connect")       //
	procFreeAddrInfo  = modws232.NewProc("freeaddrinfo")  //
	procFreeAddrInfoW = modws232.NewProc("FreeAddrInfoW") //
	procGetAddrInfo   = modws232.NewProc("getaddrinfo")   //
	procGetAddrInfoW  = modws232.NewProc("GetAddrInfoW")  //
	procGetHostByAddr = modws232.NewProc("gethostbyaddr") //
	procGetHostByName = modws232.NewProc("gethostbyname") //
	procGetHostNameA  = modws232.NewProc("gethostname")   //
	procGetHostNameW  = modws232.NewProc("GetHostNameW")  //
	procGetNameInfoA  = modws232.NewProc("getnameinfo")   //
	procGetNameInfoW  = modws232.NewProc("GetNameInfoW")  //
	procGetPeerName   = modws232.NewProc("getpeername")   //
)

// __WSAFDIsSet returns a value indicating whether a socket is included in a set of socket descriptors.
func __WSAFDIsSet(socket SOCKET, set *FDSet) (bool, error) {
	ret, _, err := procWSAFDIsSet.Call(
		uintptr(socket),
		uintptr(unsafe.Pointer(set)),
	)
	if int32(ret) == 0 {
		if !isValidErr(err) {
			err = nil
		}
		return false, err
	}
	return true, nil
}

// Accept permits an incoming connection attempt on a socket.
func Accept(s SOCKET, addr unsafe.Pointer, addrLen *int32) (SOCKET, error) {
	ret, _, err := procAccept.Call(
		uintptr(s),
		uintptr(addr),
		uintptr(unsafe.Pointer(addrLen)),
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// Bind associates a local address with a socket.
func Bind(s SOCKET, addr unsafe.Pointer, addrLen int32) (int, error) {
	ret, _, err := procBind.Call(
		uintptr(s),
		uintptr(addr),
		uintptr(addrLen),
	)
	if int32(ret) != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// CloseSocket closes an existing socket.
func CloseSocket(s SOCKET) (int, error) {
	ret, _, err := procClosesocket.Call(
		uintptr(s),
	)
	if int32(ret) != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// Connect establishes a connection to a specified socket.
func Connect(s SOCKET, name unsafe.Pointer, nameLen int32) (int, error) {
	ret, _, err := procConnect.Call(
		uintptr(s),
		uintptr(name),
		uintptr(nameLen),
	)
	if int32(ret) != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// FreeAddrInfoA frees address information that the GetAddrInfoA function dynamically allocates in AddrInfoA structures.
func FreeAddrInfoA(addrInfo *AddrInfoA) error {
	_, _, err := procFreeAddrInfo.Call(
		uintptr(unsafe.Pointer(addrInfo)),
	)
	if !isValidErr(err) {
		err = nil
	}
	return err
}

// FreeAddrInfoW frees address information that the GetAddrInfoW function dynamically allocates in AddrInfoW structures.
func FreeAddrInfoW(addrInfo *AddrInfoW) error {
	_, _, err := procFreeAddrInfoW.Call(
		uintptr(unsafe.Pointer(addrInfo)),
	)
	if !isValidErr(err) {
		err = nil
	}
	return err
}

// GetAddrInfoA provides protocol-independent translation from a ANSI host name to an address.
func GetAddrInfoA(nodeName *byte, serviceName *byte, hints *AddrInfoA, result **AddrInfoA) (int, error) {
	ret, _, err := procGetAddrInfo.Call(
		uintptr(unsafe.Pointer(nodeName)),
		uintptr(unsafe.Pointer(serviceName)),
		uintptr(unsafe.Pointer(hints)),
		uintptr(unsafe.Pointer(result)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return int(iRet), err
	}
	return 0, nil
}

// GetAddrInfoW provides protocol-independent translation from a Unicode host name to an address.
func GetAddrInfoW(nodeName *uint16, serviceName *uint16, hints *AddrInfoW, result **AddrInfoW) (int, error) {
	ret, _, err := procGetAddrInfoW.Call(
		uintptr(unsafe.Pointer(nodeName)),
		uintptr(unsafe.Pointer(serviceName)),
		uintptr(unsafe.Pointer(hints)),
		uintptr(unsafe.Pointer(result)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return int(iRet), err
	}
	return 0, nil
}

// GetHostByAddr retrieves the host information corresponding to a network address.
//
// Deprecated: Use GetNameInfo instead.
func GetHostByAddr(addr unsafe.Pointer, addrLen int32, htype int32) (*HostEnt, error) {
	ret, _, err := procGetHostByAddr.Call(
		uintptr(addr),
		uintptr(addrLen),
		uintptr(htype),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*HostEnt)(retPtr), nil
}

// GetHostByName retrieves host information corresponding to a host name from a host database.
//
// Deprecated: Use GetHostNameA/GetHostNameW instead.
func GetHostByName(name *byte) (*HostEnt, error) {
	ret, _, err := procGetHostByName.Call(
		uintptr(unsafe.Pointer(name)),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*HostEnt)(retPtr), nil
}

// GetHostNameA retrieves the standard host name for the local computer.
func GetHostNameA(name *byte, namelen int) (int, error) {
	ret, _, err := procGetHostNameA.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(namelen),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// GetHostNameW retrieves the standard host name for the local computer as a Unicode string.
func GetHostNameW(name *uint16, namelen int) (int, error) {
	ret, _, err := procGetHostNameW.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(namelen),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// GetNameInfo provides protocol-independent name resolution from an address to an ANSI
// host name and from a port number to the ANSI service name.
func GetNameInfoA(
	sockaddr unsafe.Pointer,
	sockaddrLength int32,
	nodeBuffer *byte,
	nodeBufferSize int32,
	serviceBuffer *byte,
	serviceBufferSize int32,
	flags int32,
) (int, error) {
	ret, _, err := procGetNameInfoA.Call(
		uintptr(sockaddr),
		uintptr(sockaddrLength),
		uintptr(unsafe.Pointer(nodeBuffer)),
		uintptr(nodeBufferSize),
		uintptr(unsafe.Pointer(serviceBuffer)),
		uintptr(serviceBufferSize),
		uintptr(flags),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return int(iRet), err
	}
	return 0, nil
}

// GetNameInfoW provides protocol-independent name resolution from an address to a Unicode
// host name and from a port number to the Unicode service name.
func GetNameInfoW(
	sockaddr unsafe.Pointer,
	sockaddrLength int32,
	nodeBuffer *uint16,
	nodeBufferSize int32,
	serviceBuffer *uint16,
	serviceBufferSize int32,
	flags int32,
) (int, error) {
	ret, _, err := procGetNameInfoW.Call(
		uintptr(sockaddr),
		uintptr(sockaddrLength),
		uintptr(unsafe.Pointer(nodeBuffer)),
		uintptr(nodeBufferSize),
		uintptr(unsafe.Pointer(serviceBuffer)),
		uintptr(serviceBufferSize),
		uintptr(flags),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return int(iRet), err
	}
	return 0, nil
}

// GetPeerName retrieves the address of the peer to which a socket is connected.
func GetPeerName(s SOCKET, name unsafe.Pointer, nameLen int32) (int, error) {
	ret, _, err := procGetPeerName.Call(
		uintptr(s),
		uintptr(name),
		uintptr(unsafe.Pointer(&nameLen)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// isValidErr returns true if the given error is a "real" one.
func isValidErr(err error) bool {
	if err != nil && !strings.Contains(err.Error(), "success") {
		return true
	}
	return false
}
