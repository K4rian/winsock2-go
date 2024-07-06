package ws2

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	modws232             = syscall.NewLazyDLL("ws2_32.dll")
	procWSAFDIsSet       = modws232.NewProc("__WSAFDIsSet")     //
	procAccept           = modws232.NewProc("accept")           //
	procBind             = modws232.NewProc("bind")             //
	procClosesocket      = modws232.NewProc("closesocket")      //
	procConnect          = modws232.NewProc("connect")          //
	procFreeAddrInfo     = modws232.NewProc("freeaddrinfo")     //
	procFreeAddrInfoW    = modws232.NewProc("FreeAddrInfoW")    //
	procGetAddrInfo      = modws232.NewProc("getaddrinfo")      //
	procGetAddrInfoW     = modws232.NewProc("GetAddrInfoW")     //
	procGetHostByAddr    = modws232.NewProc("gethostbyaddr")    //
	procGetHostByName    = modws232.NewProc("gethostbyname")    //
	procGetHostNameA     = modws232.NewProc("gethostname")      //
	procGetHostNameW     = modws232.NewProc("GetHostNameW")     //
	procGetNameInfoA     = modws232.NewProc("getnameinfo")      //
	procGetNameInfoW     = modws232.NewProc("GetNameInfoW")     //
	procGetPeerName      = modws232.NewProc("getpeername")      //
	procGetProtoByName   = modws232.NewProc("getprotobyname")   //
	procGetProtoByNumber = modws232.NewProc("getprotobynumber") //
	procGetServByName    = modws232.NewProc("getservbyname")    //
	procGetServByPort    = modws232.NewProc("getservbyport")    //
	procGetSockName      = modws232.NewProc("getsockname")      //
	procGetSockOpt       = modws232.NewProc("getsockopt")       //
	procInetAddr         = modws232.NewProc("inet_addr")        //
	procInetNtop         = modws232.NewProc("inet_ntop")        //
	procInetNtopW        = modws232.NewProc("InetNtopW")        //
	procInetPton         = modws232.NewProc("inet_pton")        //
	procInetPtonW        = modws232.NewProc("InetPtonW")        //
	procIoctlSocket      = modws232.NewProc("ioctlsocket")      //
	procListen           = modws232.NewProc("listen")           //
	procRecv             = modws232.NewProc("recv")             //
	procRecvfrom         = modws232.NewProc("recvfrom")         //
	procSelect           = modws232.NewProc("select")           //
	procSend             = modws232.NewProc("send")             //
	procSendTo           = modws232.NewProc("sendto")           //
	procSetSockOpt       = modws232.NewProc("setsockopt")       //
	procShutdown         = modws232.NewProc("shutdown")         //
	procSocket           = modws232.NewProc("socket")           //
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

// GetProtoByName retrieves the protocol information corresponding to a protocol name.
func GetProtoByName(name *byte) (*ProtoEnt, error) {
	ret, _, err := procGetProtoByName.Call(
		uintptr(unsafe.Pointer(name)),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*ProtoEnt)(retPtr), nil
}

// GetProtoByNumber retrieves protocol information corresponding to a protocol number.
func GetProtoByNumber(proto int32) (*ProtoEnt, error) {
	ret, _, err := procGetProtoByNumber.Call(
		uintptr(proto),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*ProtoEnt)(retPtr), nil
}

// GetServByName retrieves service information corresponding to a service name and protocol.
func GetServByName(name *byte, proto *byte) (*ServEnt, error) {
	ret, _, err := procGetServByName.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(proto)),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*ServEnt)(retPtr), nil
}

// GetServByPort retrieves service information corresponding to a port and protocol.
func GetServByPort(port uint16, proto *byte) (*ServEnt, error) {
	ret, _, err := procGetServByPort.Call(
		uintptr(port),
		uintptr(unsafe.Pointer(proto)),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*ServEnt)(retPtr), nil
}

// GetSockName retrieves the local name for a socket.
func GetSockName(s SOCKET, name unsafe.Pointer, nameLen *int32) (int, error) {
	ret, _, err := procGetSockName.Call(
		uintptr(s),
		uintptr(name),
		uintptr(unsafe.Pointer(nameLen)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// GetSockOpt retrieves a socket option.
func GetSockOpt(s SOCKET, level int32, optName int32, optVal unsafe.Pointer, optLen *int32) (int, error) {
	ret, _, err := procGetSockOpt.Call(
		uintptr(s),
		uintptr(level),
		uintptr(optName),
		uintptr(optVal),
		uintptr(unsafe.Pointer(optLen)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// Htond converts a double (float64) from host to TCP/IP network byte order (which is big-endian).
// Reimplemented in utils.go

// Htonf converts a float from host to TCP/IP network byte order (which is big-endian).
// Reimplemented in utils.go

// Htonl converts a unsigned long (uint32) from host to TCP/IP network byte order (which is big-endian).
// Reimplemented in utils.go

// Htonll converts an unsigned int64 (uint64) from host to TCP/IP network byte order (which is big-endian).
// Reimplemented in utils.go

// Htons converts a unsigned short (uint16) from host to TCP/IP network byte order (which is big-endian).
// Reimplemented in utils.go

// InetAddr converts a string containing an IPv4 dotted-decimal address into a proper address for the InAddr structure.
func InetAddr(cp *byte) uint32 {
	ret, _, _ := procInetAddr.Call(
		uintptr(unsafe.Pointer(cp)),
	)
	return uint32(ret)
}

// InetNtoa converts an (IPv4) Internet network address into an ASCII string in Internet standard dotted-decimal format.
// Reimplemented in utils.go

// InetNtop converts an IPv4 or IPv6 Internet network address into a string in Internet standard format (ANSI).
func InetNtop(family uint16, addr unsafe.Pointer, stringBuf *byte, stringBufSize int32) (*byte, error) {
	ret, _, err := procInetNtop.Call(
		uintptr(family),
		uintptr(addr),
		uintptr(unsafe.Pointer(stringBuf)),
		uintptr(stringBufSize),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*byte)(retPtr), nil
}

// InetNtopW converts an IPv4 or IPv6 Internet network address into a string in Internet standard format (Unicode).
func InetNtopW(family uint16, addr unsafe.Pointer, stringBuf *uint16, stringBufSize int32) (*uint16, error) {
	ret, _, err := procInetNtopW.Call(
		uintptr(family),
		uintptr(addr),
		uintptr(unsafe.Pointer(stringBuf)),
		uintptr(stringBufSize),
	)
	retPtr := unsafe.Pointer(ret)
	if retPtr == nil {
		if !isValidErr(err) {
			err = nil
		}
		return nil, err
	}
	return (*uint16)(retPtr), nil
}

// InetPton converts an IPv4 or IPv6 Internet network address in its standard text presentation form into its numeric binary form (ANSI).
func InetPton(family uint16, addrString *byte, addrBuf unsafe.Pointer) (int, error) {
	ret, _, err := procInetPton.Call(
		uintptr(family),
		uintptr(unsafe.Pointer(addrString)),
		uintptr(addrBuf),
	)
	iRet := int32(ret)
	if iRet != 1 {
		return int(ret), err
	}
	return 1, nil
}

// InetPtonW converts an IPv4 or IPv6 Internet network address in its standard text presentation form into its numeric binary form (Unicode).
func InetPtonW(family uint16, addrString *uint16, addrBuf unsafe.Pointer) (int, error) {
	ret, _, err := procInetPtonW.Call(
		uintptr(family),
		uintptr(unsafe.Pointer(addrString)),
		uintptr(addrBuf),
	)
	iRet := int32(ret)
	if iRet != 1 {
		return int(ret), err
	}
	return 1, nil
}

// IoctlSocket controls the I/O mode of a socket.
func IoctlSocket(s SOCKET, cmd uint32, argp *uint32) (int, error) {
	ret, _, err := procIoctlSocket.Call(
		uintptr(s),
		uintptr(cmd),
		uintptr(unsafe.Pointer(argp)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// Listen places a socket in a state in which it is listening for an incoming connection.
func Listen(s SOCKET, backlog int32) (int, error) {
	ret, _, err := procListen.Call(
		uintptr(s),
		uintptr(backlog),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// Ntohd
// Ignored

// Ntohf
// Ignored

// Ntohl converts a unsigned long (uint32) from TCP/IP network order to host byte order (which is little-endian on Intel processors).
// Reimplemented in utils.go

// Ntohll
// Ignored

// Ntohs converts a unsigned short (uint16) from TCP/IP network byte order to host byte order (which is little-endian on Intel processors).
// Reimplemented in utils.go

// ProcessSocketNotifications
// Ignored

// Recv function receives data from a connected socket or a bound connectionless socket.
func Recv(s SOCKET, buf unsafe.Pointer, len int32, flags int32) (int, error) {
	ret, _, err := procRecv.Call(
		uintptr(s),
		uintptr(buf),
		uintptr(len),
		uintptr(flags),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// RecvFrom receives a datagram, and stores the source address.
func RecvFrom(s SOCKET, buf unsafe.Pointer, len int32, flags int32, from unsafe.Pointer, fromLen *int32) (int, error) {
	ret, _, err := procRecvfrom.Call(
		uintptr(s),
		uintptr(buf),
		uintptr(len),
		uintptr(flags),
		uintptr(from),
		uintptr(unsafe.Pointer(fromLen)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// Select determines the status of one or more sockets, waiting if necessary, to perform synchronous I/O.
func Select(nfds int32, readfds *FDSet, writefds *FDSet, exceptfds *FDSet, timeout *Timeval) (int, error) {
	ret, _, err := procSelect.Call(
		uintptr(nfds),
		uintptr(unsafe.Pointer(readfds)),
		uintptr(unsafe.Pointer(writefds)),
		uintptr(unsafe.Pointer(exceptfds)),
		uintptr(unsafe.Pointer(timeout)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// Send sends data on a connected socket.
func Send(s SOCKET, buf unsafe.Pointer, len, flags int32) (int, error) {
	ret, _, err := procSend.Call(
		uintptr(s),
		uintptr(buf),
		uintptr(len),
		uintptr(flags),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// SendTo sends data to a specific destination.
func SendTo(s SOCKET, buf unsafe.Pointer, len int32, flags int32, to unsafe.Pointer, toLen int32) (int, error) {
	ret, _, err := procSendTo.Call(
		uintptr(s),
		uintptr(buf),
		uintptr(len),
		uintptr(flags),
		uintptr(to),
		uintptr(toLen),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// SetSockOpt sets a socket option.
func SetSockOpt(s SOCKET, level int32, optName int32, optVal unsafe.Pointer, optLen int32) (int, error) {
	ret, _, err := procSetSockOpt.Call(
		uintptr(s),
		uintptr(level),
		uintptr(optName),
		uintptr(optVal),
		uintptr(optLen),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// Shutdown function disables sends or receives on a socket.
func Shutdown(s SOCKET, how int32) (int, error) {
	ret, _, err := procShutdown.Call(
		uintptr(s),
		uintptr(how),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// Socket creates a socket that is bound to a specific transport service provider.
func Socket(af int32, stype int32, protocol int32) (SOCKET, error) {
	ret, _, err := procSocket.Call(
		uintptr(af),
		uintptr(stype),
		uintptr(protocol),
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// SocketNotificationRetrieveEvents
// Ignored

// isValidErr returns true if the given error is a "real" one.
func isValidErr(err error) bool {
	if err != nil && !strings.Contains(err.Error(), "success") {
		return true
	}
	return false
}
