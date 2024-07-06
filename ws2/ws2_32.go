package ws2

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	modws232                         = syscall.NewLazyDLL("ws2_32.dll")
	procWSAFDIsSet                   = modws232.NewProc("__WSAFDIsSet")                 //
	procAccept                       = modws232.NewProc("accept")                       //
	procBind                         = modws232.NewProc("bind")                         //
	procClosesocket                  = modws232.NewProc("closesocket")                  //
	procConnect                      = modws232.NewProc("connect")                      //
	procFreeAddrInfo                 = modws232.NewProc("freeaddrinfo")                 //
	procFreeAddrInfoW                = modws232.NewProc("FreeAddrInfoW")                //
	procGetAddrInfo                  = modws232.NewProc("getaddrinfo")                  //
	procGetAddrInfoW                 = modws232.NewProc("GetAddrInfoW")                 //
	procGetHostByAddr                = modws232.NewProc("gethostbyaddr")                //
	procGetHostByName                = modws232.NewProc("gethostbyname")                //
	procGetHostNameA                 = modws232.NewProc("gethostname")                  //
	procGetHostNameW                 = modws232.NewProc("GetHostNameW")                 //
	procGetNameInfoA                 = modws232.NewProc("getnameinfo")                  //
	procGetNameInfoW                 = modws232.NewProc("GetNameInfoW")                 //
	procGetPeerName                  = modws232.NewProc("getpeername")                  //
	procGetProtoByName               = modws232.NewProc("getprotobyname")               //
	procGetProtoByNumber             = modws232.NewProc("getprotobynumber")             //
	procGetServByName                = modws232.NewProc("getservbyname")                //
	procGetServByPort                = modws232.NewProc("getservbyport")                //
	procGetSockName                  = modws232.NewProc("getsockname")                  //
	procGetSockOpt                   = modws232.NewProc("getsockopt")                   //
	procInetAddr                     = modws232.NewProc("inet_addr")                    //
	procInetNtop                     = modws232.NewProc("inet_ntop")                    //
	procInetNtopW                    = modws232.NewProc("InetNtopW")                    //
	procInetPton                     = modws232.NewProc("inet_pton")                    //
	procInetPtonW                    = modws232.NewProc("InetPtonW")                    //
	procIoctlSocket                  = modws232.NewProc("ioctlsocket")                  //
	procListen                       = modws232.NewProc("listen")                       //
	procRecv                         = modws232.NewProc("recv")                         //
	procRecvfrom                     = modws232.NewProc("recvfrom")                     //
	procSelect                       = modws232.NewProc("select")                       //
	procSend                         = modws232.NewProc("send")                         //
	procSendTo                       = modws232.NewProc("sendto")                       //
	procSetSockOpt                   = modws232.NewProc("setsockopt")                   //
	procShutdown                     = modws232.NewProc("shutdown")                     //
	procSocket                       = modws232.NewProc("socket")                       //
	procWSAAccept                    = modws232.NewProc("WSAAccept")                    // The condition and callbackData parameters are always ignored.
	procWSAAddressToStringA          = modws232.NewProc("WSAAddressToStringA")          //
	procWSAAddressToStringW          = modws232.NewProc("WSAAddressToStringW")          //
	procWSACleanup                   = modws232.NewProc("WSACleanup")                   //
	procWSACloseEvent                = modws232.NewProc("WSACloseEvent")                //
	procWSAConnect                   = modws232.NewProc("WSAConnect")                   //
	procWSAConnectByList             = modws232.NewProc("WSAConnectByList")             // Broken. Can't get it to work (SocketAddressList issue?).
	procWSAConnectByNameA            = modws232.NewProc("WSAConnectByNameA")            //
	procWSAConnectByNameW            = modws232.NewProc("WSAConnectByNameW")            //
	procWSACreateEvent               = modws232.NewProc("WSACreateEvent")               //
	procWSADuplicateSocketA          = modws232.NewProc("WSADuplicateSocketA")          //
	procWSADuplicateSocketW          = modws232.NewProc("WSADuplicateSocketW")          //
	procWSAEnumNameSpaceProvidersA   = modws232.NewProc("WSAEnumNameSpaceProvidersA")   //
	procWSAEnumNameSpaceProvidersExA = modws232.NewProc("WSAEnumNameSpaceProvidersExA") //
	procWSAEnumNameSpaceProvidersExW = modws232.NewProc("WSAEnumNameSpaceProvidersExW") // Crash when iterating over the buffer after the first element (memory alignment issue?).
	procWSAEnumNameSpaceProvidersW   = modws232.NewProc("WSAEnumNameSpaceProvidersW")   //
	procWSAEnumNetworkEvents         = modws232.NewProc("WSAEnumNetworkEvents")         //
	procWSAEnumProtocolsA            = modws232.NewProc("WSAEnumProtocolsA")            //
	procWSAEnumProtocolsW            = modws232.NewProc("WSAEnumProtocolsW")            //
	procWSAEventSelect               = modws232.NewProc("WSAEventSelect")               //
	procWSAGetLastError              = modws232.NewProc("WSAGetLastError")              // Broken? Always returns 0.
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

// WSAAccept accepts a connection based on the return value of a condition function,
// provides quality of service flow specifications, and allows the transfer of connection data.
//
// Warning: Partially implemented, the condition and callbackData parameters are always ignored.
func WSAAccept(s SOCKET, addr unsafe.Pointer, addrLen *int32, condition unsafe.Pointer, callbackData uintptr) (SOCKET, error) {
	ret, _, err := procWSAAccept.Call(
		uintptr(s),
		uintptr(addr),
		uintptr(unsafe.Pointer(&addrLen)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// WSAAddressToStringA converts all components of a sockaddr structure into a human-readable string representation of the address.
func WSAAddressToStringA(
	address unsafe.Pointer,
	addressLength uint32,
	protocolInfo *WSAProtocolInfoA,
	addressString *byte,
	addressStringLength *uint32,
) (int, error) {
	ret, _, err := procWSAAddressToStringA.Call(
		uintptr(address),
		uintptr(addressLength),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(unsafe.Pointer(addressString)),
		uintptr(unsafe.Pointer(addressStringLength)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAAddressToStringW converts all components of a sockaddr structure into a human-readable string representation of the address.
func WSAAddressToStringW(
	address unsafe.Pointer,
	addressLength uint32,
	protocolInfo *WSAProtocolInfoW,
	addressString *uint16,
	addressStringLength *uint32,
) (int, error) {
	ret, _, err := procWSAAddressToStringW.Call(
		uintptr(address),
		uintptr(addressLength),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(unsafe.Pointer(addressString)),
		uintptr(unsafe.Pointer(addressStringLength)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAAsyncGetHostByAddr asynchronously retrieves host information that corresponds to an address.
// Ignored

// WSAAsyncGetHostByName asynchronously retrieves host information that corresponds to a host name.
// Ignored

// WSAAsyncGetProtoByName asynchronously retrieves protocol information that corresponds to a protocol name.
// Ignored

// WSAAsyncGetProtoByNumber asynchronously retrieves protocol information that corresponds to a protocol number.
// Ignored

// WSAAsyncGetServByName asynchronously retrieves service information that corresponds to a service name and port.
// Ignored

// WSAAsyncGetServByPort asynchronously retrieves service information that corresponds to a port and protocol.
// Ignored

// WSAAsyncSelect requests Windows message-based notification of network events for a socket.
// Ignored

// WSACancelAsyncRequest cancels an incomplete asynchronous operation.
// Ignored

// WSACleanup terminates use of the Winsock DLL.
func WSACleanup() (bool, error) {
	ret, _, err := procWSACleanup.Call()
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return false, err
	}
	return true, nil
}

// WSACloseEvent closes an open event object handle.
func WSACloseEvent(event WSAEVENT) (bool, error) {
	ret, _, err := procWSACloseEvent.Call(
		uintptr(event),
	)
	if int32(ret) == 0 {
		return false, err
	}
	return true, nil
}

// WSAConnect establishes a connection to another socket application, exchanges connect data,
// and specifies required quality of service based on the specified FlowSpec structure.
func WSAConnect(
	s SOCKET,
	name unsafe.Pointer,
	namelen int32,
	callerData *WSABuf,
	calleeData *WSABuf,
	cQOS *QOS,
	gQOS *QOS,
) (int, error) {
	ret, _, err := procWSAConnect.Call(
		uintptr(s),
		uintptr(name),
		uintptr(unsafe.Pointer(&namelen)),
		uintptr(unsafe.Pointer(callerData)),
		uintptr(unsafe.Pointer(calleeData)),
		uintptr(unsafe.Pointer(&cQOS)),
		uintptr(unsafe.Pointer(nil)), // Should be nil (see the official doc for more info)
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAConnectByList establishes a connection to one out of a collection of possible endpoints represented by a set of destination addresses (host names and ports).
// This function takes all the destination addresses passed to it and all of the local computer's source addresses,
// and tries connecting using all possible address combinations before giving up.
//
// This function supports both IPv4 and IPv6 addresses.
/*
func WSAConnectByList(
	s SOCKET,
	socketAddress *SocketAddressList,
	localAddressLength *uint32,
	localAddress unsafe.Pointer,
	remoteAddressLength *uint32,
	remoteAddress unsafe.Pointer,
	timeout *Timeval,
	reserved *WSAOverlapped,
) (bool, error) {
	ret, _, err := procWSAConnectByList.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(socketAddress)),
		uintptr(unsafe.Pointer(localAddressLength)),
		uintptr(localAddress),
		uintptr(unsafe.Pointer(remoteAddressLength)),
		uintptr(remoteAddress),
		uintptr(unsafe.Pointer(timeout)),
		uintptr(unsafe.Pointer(nil)), // Must be set to nil
	)
	bRet := int32(ret) != 0
	if isValidErr(err) {
		return bRet, err
	}
	return bRet, nil
}
*/

// WSAConnectByNameA establishes a connection to a specified host and port.
//
// This function is provided to allow a quick connection to a network endpoint given a host name and port.
//
// This function supports both IPv4 and IPv6 addresses.
func WSAConnectByNameA(
	s SOCKET,
	nodename *byte,
	servicename *byte,
	localAddressLength *uint32,
	localAddress unsafe.Pointer,
	remoteAddressLength *uint32,
	remoteAddress unsafe.Pointer,
	timeout *Timeval,
	reserved *WSAOverlapped,
) (bool, error) {
	ret, _, err := procWSAConnectByNameA.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(nodename)),
		uintptr(unsafe.Pointer(servicename)),
		uintptr(unsafe.Pointer(&localAddressLength)),
		uintptr(localAddress),
		uintptr(unsafe.Pointer(&remoteAddressLength)),
		uintptr(remoteAddress),
		uintptr(unsafe.Pointer(timeout)),
		uintptr(unsafe.Pointer(nil)),
	)
	bRet := int32(ret) != 0
	if isValidErr(err) {
		return bRet, err
	}
	return bRet, nil
}

// WSAConnectByNameW establishes a connection to a specified host and port.
//
// This function is provided to allow a quick connection to a network endpoint given a host name and port.
//
// This function supports both IPv4 and IPv6 addresses.
func WSAConnectByNameW(
	s SOCKET,
	nodename *uint16,
	servicename *uint16,
	localAddressLength *uint32,
	localAddress unsafe.Pointer,
	remoteAddressLength *uint32,
	remoteAddress unsafe.Pointer,
	timeout *Timeval,
	reserved *WSAOverlapped,
) (bool, error) {
	ret, _, err := procWSAConnectByNameW.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(nodename)),
		uintptr(unsafe.Pointer(servicename)),
		uintptr(unsafe.Pointer(&localAddressLength)),
		uintptr(localAddress),
		uintptr(unsafe.Pointer(&remoteAddressLength)),
		uintptr(remoteAddress),
		uintptr(unsafe.Pointer(timeout)),
		uintptr(unsafe.Pointer(nil)),
	)
	bRet := int32(ret) != 0
	if isValidErr(err) {
		return bRet, err
	}
	return bRet, nil
}

// WSACreateEvent creates a new event object.
func WSACreateEvent() (WSAEVENT, error) {
	ret, _, err := procWSACreateEvent.Call()
	event := WSAEVENT(ret)
	if event == WSA_INVALID_EVENT {
		return WSA_INVALID_EVENT, err
	}
	return event, nil
}

// WSADuplicateSocketA returns a WSAProtocolInfoA structure that can be used to create a new socket descriptor for a shared socket.
//
// The WSADuplicateSocket function cannot be used on a QOS-enabled socket.
func WSADuplicateSocketA(s SOCKET, processId uint32, protocolInfo *WSAProtocolInfoA) (bool, error) {
	ret, _, err := procWSADuplicateSocketA.Call(
		uintptr(s),
		uintptr(processId),
		uintptr(unsafe.Pointer(protocolInfo)),
	)
	if int32(ret) != 0 {
		return false, err
	}
	return true, nil
}

// WSADuplicateSocketW returns a WSAProtocolInfoW structure that can be used to create a new socket descriptor for a shared socket.
//
// The WSADuplicateSocket function cannot be used on a QOS-enabled socket.
func WSADuplicateSocketW(s SOCKET, processId uint32, protocolInfo *WSAProtocolInfoW) (bool, error) {
	ret, _, err := procWSADuplicateSocketW.Call(
		uintptr(s),
		uintptr(processId),
		uintptr(unsafe.Pointer(protocolInfo)),
	)
	if int32(ret) != 0 {
		return false, err
	}
	return true, nil
}

// WSAEnumNameSpaceProvidersA retrieves information on available namespace providers.
//
// NOTE: lpnspBuffer should be a pointer to the first element of a WSANameSpaceInfoA slice.
func WSAEnumNameSpaceProvidersA(bufferLength *uint32, buffer *WSANameSpaceInfoA) (int, error) {
	ret, _, err := procWSAEnumNameSpaceProvidersA.Call(
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(buffer)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAEnumNameSpaceProvidersExA retrieves information on available namespace providers.
//
// NOTE: buffer should be a pointer to the first element of a WSANameSpaceInfoExA slice.
//
// NOTE: buffer[x].ProviderSpecific doesn't contain any data, only the length.
func WSAEnumNameSpaceProvidersExA(bufferLength *uint32, buffer *WSANameSpaceInfoExA) (int, error) {
	ret, _, err := procWSAEnumNameSpaceProvidersExA.Call(
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(buffer)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAEnumNameSpaceProvidersExW retrieves information on available namespace providers.
//
// NOTE: buffer should be a pointer to the first element of a WSANameSpaceInfoExW slice.
//
// NOTE: buffer[x].ProviderSpecific doesn't contain any data, only the length.
/*
func WSAEnumNameSpaceProvidersExW(bufferLength *uint32, buffer *WSANameSpaceInfoExW) (int, error) {
	ret, _, err := procWSAEnumNameSpaceProvidersExW.Call(
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(buffer)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}
*/

// WSAEnumNameSpaceProvidersW retrieves information on available namespace providers.
//
// NOTE: buffer should be a pointer to the first element of a WSANameSpaceInfoW slice.
func WSAEnumNameSpaceProvidersW(bufferLength *uint32, buffer *WSANameSpaceInfoW) (int, error) {
	ret, _, err := procWSAEnumNameSpaceProvidersW.Call(
		uintptr(unsafe.Pointer(bufferLength)),
		uintptr(unsafe.Pointer(buffer)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAEnumNetworkEvents discovers occurrences of network events for the indicated socket,
// clear internal network event records, and reset event objects (optional).
func WSAEnumNetworkEvents(s SOCKET, eventObject WSAEVENT, networkEvents *WSANetworkEvents) (int, error) {
	ret, _, err := procWSAEnumNetworkEvents.Call(
		uintptr(s),
		uintptr(eventObject),
		uintptr(unsafe.Pointer(networkEvents)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAEnumProtocolsA retrieves information about available transport protocols.
func WSAEnumProtocolsA(protocols *int32, protocolBuffer *WSAProtocolInfoA, bufferLength *uint32) (int, error) {
	ret, _, err := procWSAEnumProtocolsA.Call(
		uintptr(unsafe.Pointer(protocols)),
		uintptr(unsafe.Pointer(protocolBuffer)),
		uintptr(unsafe.Pointer(bufferLength)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAEnumProtocolsW retrieves information about available transport protocols.
func WSAEnumProtocolsW(protocols *int32, protocolBuffer *WSAProtocolInfoW, bufferLength *uint32) (int, error) {
	ret, _, err := procWSAEnumProtocolsW.Call(
		uintptr(unsafe.Pointer(protocols)),
		uintptr(unsafe.Pointer(protocolBuffer)),
		uintptr(unsafe.Pointer(bufferLength)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAEventSelect specifies an event object to be associated with the specified set of FD_XXX network events.
func WSAEventSelect(s SOCKET, eventObject WSAEVENT, networkEvents int32) (int, error) {
	ret, _, err := procWSAEventSelect.Call(
		uintptr(s),
		uintptr(eventObject),
		uintptr(networkEvents),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAGetLastError returns the error status for the last Windows Sockets operation that failed.
func WSAGetLastError() int {
	ret, _, _ := procWSAGetLastError.Call()
	return int(int32(ret))
}

// WSAGetOverlappedResult retrieves the results of an overlapped operation on the specified socket.
// TODO

// WSAGetQOSByName initializes a QOS structure based on a named template, or it supplies a buffer to retrieve
// an enumeration of the available template names.
// TODO

// WSAGetServiceClassInfoA retrieves the class information (schema) pertaining to a specified service class from a specified namespace provider.
// TODO

// WSAGetServiceClassInfoW retrieves the class information (schema) pertaining to a specified service class from a specified namespace provider.
// TODO

// WSAGetServiceClassNameByClassIdA retrieves the name of the service associated with the specified type.
//
// This name is the generic service name, like FTP or SNA, and not the name of a specific instance of that service.
// TODO

// WSAGetServiceClassNameByClassIdW retrieves the name of the service associated with the specified type.
//
// This name is the generic service name, like FTP or SNA, and not the name of a specific instance of that service.
// TODO

// WSAHtonl converts a signed long (int32) from host byte order to network byte order.
// TODO

// WSAHtons converts a signed short (int16) from host byte order to network byte order.
// TODO

// WSAInstallServiceClassA registers a service class schema within a namespace.
//
// This schema includes the class name, class identifier, and any namespace-specific information that is common to all
// instances of the service, such as the SAP identifier or object identifier.
// TODO

// WSAInstallServiceClassW registers a service class schema within a namespace.
//
// This schema includes the class name, class identifier, and any namespace-specific information that is common to all
// instances of the service, such as the SAP identifier or object identifier.
// TODO

// WSAIoctl controls the mode of a socket.
// TODO

// WSAIsBlocking
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSAJoinLeaf joins a leaf node into a multipoint session, exchanges connect data,
// and specifies needed quality of service based on the specified FlowSpec structures.
// TODO

// WSALookupServiceBeginA initiates a client query that is constrained by the information contained within a WSAQuerySetA structure.
//
// WSALookupServiceBeginA only returns a handle, which should be used by subsequent calls to WSALookupServiceNext to get the actual results.
// TODO

// WSALookupServiceBeginW initiates a client query that is constrained by the information contained within a WSAQuerySetW structure.
//
// WSALookupServiceBeginW only returns a handle, which should be used by subsequent calls to WSALookupServiceNext to get the actual results.
// TODO

// WSALookupServiceEnd is called to free the handle after previous calls to WSALookupServiceBegin and WSALookupServiceNext.
//
// If you call WSALookupServiceEnd from another thread while an existing WSALookupServiceNext is blocked, the end call will have
// the same effect as a cancel and will cause the WSALookupServiceNext call to return immediately.
// TODO

// WSALookupServiceNextA is called after obtaining a handle from a previous call to WSALookupServiceBegin in order to retrieve the requested service information.
//
// The provider will pass back a WSAQuerySet structure in the results buffer.
//
// The client should continue to call this function until it returns WSA_E_NO_MORE, indicating that all of WSAQuerySet has been returned.
// TODO

// WSALookupServiceNextW is called after obtaining a handle from a previous call to WSALookupServiceBegin in order to retrieve the requested service information.
//
// The provider will pass back a WSAQuerySet structure in the results buffer.
//
// The client should continue to call this function until it returns WSA_E_NO_MORE, indicating that all of WSAQuerySet has been returned.
// TODO

// WSANSPIoctl enables developers to make I/O control calls to a registered namespace.
// TODO

// WSANtohl converts a unsigned long (uint32) from network byte order to host byte order.
// TODO

// WSANtohs converts a unsigned short (uint16) from network byte order to host byte order.
// TODO

// WSAPoll determines status of one or more sockets.
// TODO

// WSAProviderConfigChange notifies the application when the provider configuration is changed.
// TODO

// WSARecv receives data from a connected socket or a bound connectionless socket.
// TODO

// WSARecvDisconnect terminates reception on a socket, and retrieves the disconnect data if the socket is connection oriented.
// TODO

// WSARecvFrom receives a datagram and stores the source address.
// TODO

// WSARemoveServiceClass permanently removes the service class schema from the registry.
// TODO

// WSAResetEvent resets the state of the specified event object to nonsignaled.
// TODO

// WSASend sends data on a connected socket.
// TODO

// WSASendDisconnect initiates termination of the connection for the socket and sends disconnect data.
// TODO

// WSASendMsg sends data and optional control information from connected and unconnected sockets.
// TODO

// WSASendTo sends data to a specific destination, using overlapped I/O where applicable.
// TODO

// WSASetBlockingHook
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSASetEvent sets the state of the specified event object to signaled.
// TODO

// WSASetLastError sets the error code that can be retrieved through the WSAGetLastError function.
// TODO

// WSASetServiceA registers or removes from the registry a service instance within one or more namespaces.
// TODO

// WSASetServiceW registers or removes from the registry a service instance within one or more namespaces.
// TODO

// WSASocketA creates a socket that is bound to a specific transport-service provider.
// TODO

// WSASocketW creates a socket that is bound to a specific transport-service provider.
// TODO

// WSAStartup initiates use of the Winsock DLL by a process.
// TODO

// WSAStringToAddressA converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure,
// suitable for passing to Windows Sockets routines that take such a structure.
// TODO

// WSAStringToAddressW converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure,
// suitable for passing to Windows Sockets routines that take such a structure.
// TODO

// WSAUnhookBlockingHook
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSAWaitForMultipleEvents returns when one or all of the specified event objects are in the signaled state,
// when the time-out interval expires, or when an I/O completion routine has executed.
// TODO

// isValidErr returns true if the given error is a "real" one.
func isValidErr(err error) bool {
	if err != nil && !strings.Contains(err.Error(), "success") {
		return true
	}
	return false
}
