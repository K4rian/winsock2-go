package ws2

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	modws232                             = syscall.NewLazyDLL("ws2_32.dll")
	procWSAFDIsSet                       = modws232.NewProc("__WSAFDIsSet")                     //
	procAccept                           = modws232.NewProc("accept")                           //
	procBind                             = modws232.NewProc("bind")                             //
	procClosesocket                      = modws232.NewProc("closesocket")                      //
	procConnect                          = modws232.NewProc("connect")                          //
	procFreeAddrInfo                     = modws232.NewProc("freeaddrinfo")                     //
	procFreeAddrInfoW                    = modws232.NewProc("FreeAddrInfoW")                    //
	procGetAddrInfo                      = modws232.NewProc("getaddrinfo")                      //
	procGetAddrInfoW                     = modws232.NewProc("GetAddrInfoW")                     //
	procGetHostByAddr                    = modws232.NewProc("gethostbyaddr")                    //
	procGetHostByName                    = modws232.NewProc("gethostbyname")                    //
	procGetHostNameA                     = modws232.NewProc("gethostname")                      //
	procGetHostNameW                     = modws232.NewProc("GetHostNameW")                     //
	procGetNameInfoA                     = modws232.NewProc("getnameinfo")                      //
	procGetNameInfoW                     = modws232.NewProc("GetNameInfoW")                     //
	procGetPeerName                      = modws232.NewProc("getpeername")                      //
	procGetProtoByName                   = modws232.NewProc("getprotobyname")                   //
	procGetProtoByNumber                 = modws232.NewProc("getprotobynumber")                 //
	procGetServByName                    = modws232.NewProc("getservbyname")                    //
	procGetServByPort                    = modws232.NewProc("getservbyport")                    //
	procGetSockName                      = modws232.NewProc("getsockname")                      //
	procGetSockOpt                       = modws232.NewProc("getsockopt")                       //
	procInetAddr                         = modws232.NewProc("inet_addr")                        //
	procInetNtop                         = modws232.NewProc("inet_ntop")                        //
	procInetNtopW                        = modws232.NewProc("InetNtopW")                        //
	procInetPton                         = modws232.NewProc("inet_pton")                        //
	procInetPtonW                        = modws232.NewProc("InetPtonW")                        //
	procIoctlSocket                      = modws232.NewProc("ioctlsocket")                      //
	procListen                           = modws232.NewProc("listen")                           //
	procRecv                             = modws232.NewProc("recv")                             //
	procRecvfrom                         = modws232.NewProc("recvfrom")                         //
	procSelect                           = modws232.NewProc("select")                           //
	procSend                             = modws232.NewProc("send")                             //
	procSendTo                           = modws232.NewProc("sendto")                           //
	procSetSockOpt                       = modws232.NewProc("setsockopt")                       //
	procShutdown                         = modws232.NewProc("shutdown")                         //
	procSocket                           = modws232.NewProc("socket")                           //
	procWSAAccept                        = modws232.NewProc("WSAAccept")                        // The condition and callbackData parameters are always ignored.
	procWSAAddressToStringA              = modws232.NewProc("WSAAddressToStringA")              //
	procWSAAddressToStringW              = modws232.NewProc("WSAAddressToStringW")              //
	procWSACleanup                       = modws232.NewProc("WSACleanup")                       //
	procWSACloseEvent                    = modws232.NewProc("WSACloseEvent")                    //
	procWSAConnect                       = modws232.NewProc("WSAConnect")                       //
	procWSAConnectByList                 = modws232.NewProc("WSAConnectByList")                 // Broken. Can't get it to work (SocketAddressList issue?).
	procWSAConnectByNameA                = modws232.NewProc("WSAConnectByNameA")                //
	procWSAConnectByNameW                = modws232.NewProc("WSAConnectByNameW")                //
	procWSACreateEvent                   = modws232.NewProc("WSACreateEvent")                   //
	procWSADuplicateSocketA              = modws232.NewProc("WSADuplicateSocketA")              //
	procWSADuplicateSocketW              = modws232.NewProc("WSADuplicateSocketW")              //
	procWSAEnumNameSpaceProvidersA       = modws232.NewProc("WSAEnumNameSpaceProvidersA")       //
	procWSAEnumNameSpaceProvidersExA     = modws232.NewProc("WSAEnumNameSpaceProvidersExA")     //
	procWSAEnumNameSpaceProvidersExW     = modws232.NewProc("WSAEnumNameSpaceProvidersExW")     // Crash when iterating over the buffer after the first element (memory alignment issue?).
	procWSAEnumNameSpaceProvidersW       = modws232.NewProc("WSAEnumNameSpaceProvidersW")       //
	procWSAEnumNetworkEvents             = modws232.NewProc("WSAEnumNetworkEvents")             //
	procWSAEnumProtocolsA                = modws232.NewProc("WSAEnumProtocolsA")                //
	procWSAEnumProtocolsW                = modws232.NewProc("WSAEnumProtocolsW")                //
	procWSAEventSelect                   = modws232.NewProc("WSAEventSelect")                   //
	procWSAGetLastError                  = modws232.NewProc("WSAGetLastError")                  // Broken? Always returns 0.
	procWSAGetOverlappedResult           = modws232.NewProc("WSAGetOverlappedResult")           //
	procWSAGetQOSByName                  = modws232.NewProc("WSAGetQOSByName")                  // Broken? The WSABuf isn't populated, even with a NT string.
	procWSAGetServiceClassInfoA          = modws232.NewProc("WSAGetServiceClassInfoA")          // Can't get it to work: "An invalid argument was supplied.".
	procWSAGetServiceClassInfoW          = modws232.NewProc("WSAGetServiceClassInfoW")          // Can't get it to work: "An invalid argument was supplied.".
	procWSAGetServiceClassNameByClassIdA = modws232.NewProc("WSAGetServiceClassNameByClassIdA") // Crash, tried many approaches without success. (0xc0000005 - Access Violation)
	procWSAGetServiceClassNameByClassIdW = modws232.NewProc("WSAGetServiceClassNameByClassIdW") // Crash, tried many approaches without success. (0xc0000005 - Access Violation)
	procWSAHtonl                         = modws232.NewProc("WSAHtonl")                         //
	procWSAHtons                         = modws232.NewProc("WSAHtons")                         //
	procWSAInstallServiceClassA          = modws232.NewProc("WSAInstallServiceClassA")          // Broken. Raise an "A socket operation was attempted to an unreachable host" error(?).
	procWSAInstallServiceClassW          = modws232.NewProc("WSAInstallServiceClassW")          // Broken. Raise an "A socket operation was attempted to an unreachable host" error(?).
	procWSAIoctl                         = modws232.NewProc("WSAIoctl")                         // The completionRoutine parameter is always ignored.
	procWSAJoinLeaf                      = modws232.NewProc("WSAJoinLeaf")                      // Untested.
	procWSALookupServiceBeginA           = modws232.NewProc("WSALookupServiceBeginA")           //
	procWSALookupServiceBeginW           = modws232.NewProc("WSALookupServiceBeginW")           //
	procWSALookupServiceEnd              = modws232.NewProc("WSALookupServiceEnd")              //
	procWSALookupServiceNextA            = modws232.NewProc("WSALookupServiceNextA")            //
	procWSALookupServiceNextW            = modws232.NewProc("WSALookupServiceNextW")            // Crash, even with a valid handle from WSALookupServiceBeginW.
	procWSANSPIoctl                      = modws232.NewProc("WSANSPIoctl")                      // Seems to works, not depth tested. The completionRoutine parameter is always ignored.
	procWSANtohl                         = modws232.NewProc("WSANtohl")                         //
	procWSANtohs                         = modws232.NewProc("WSANtohs")                         //
	procWSAPoll                          = modws232.NewProc("WSAPoll")                          //
	procWSAProviderConfigChange          = modws232.NewProc("WSAProviderConfigChange")          //
	procWSARecv                          = modws232.NewProc("WSARecv")                          //
	procWSARecvDisconnect                = modws232.NewProc("WSARecvDisconnect")                //
	procWSARecvFrom                      = modws232.NewProc("WSARecvFrom")                      //
	procWSARemoveServiceClass            = modws232.NewProc("WSARemoveServiceClass")            //
	procWSAResetEvent                    = modws232.NewProc("WSAResetEvent")                    //
	procWSASend                          = modws232.NewProc("WSASend")                          //
	procWSASendDisconnect                = modws232.NewProc("WSASendDisconnect")                //
	procWSASendMsg                       = modws232.NewProc("WSASendMsg")                       // Untested.
	procWSASendTo                        = modws232.NewProc("WSASendTo")                        //
	procWSASetEvent                      = modws232.NewProc("WSASetEvent")                      //
	procWSASetLastError                  = modws232.NewProc("WSASetLastError")                  // Seems to works, not sure with the GetLastError issue.
	procWSASetServiceA                   = modws232.NewProc("WSASetServiceA")                   // Seems to works, not depth tested.
	procWSASetServiceW                   = modws232.NewProc("WSASetServiceW")                   // Seems to works, not depth tested.
	procWSASocketA                       = modws232.NewProc("WSASocketA")                       //
	procWSASocketW                       = modws232.NewProc("WSASocketW")                       //
	procWSAStartup                       = modws232.NewProc("WSAStartup")                       //
	procWSAStringToAddressA              = modws232.NewProc("WSAStringToAddressA")              //
	procWSAStringToAddressW              = modws232.NewProc("WSAStringToAddressW")              //
	procWSAWaitForMultipleEvents         = modws232.NewProc("WSAWaitForMultipleEvents")         //
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

// GetNameInfoA provides protocol-independent name resolution from an address to an ANSI
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

// Recv receives data from a connected socket or a bound connectionless socket.
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
func Send(s SOCKET, buf unsafe.Pointer, len int32, flags int32) (int, error) {
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

// Shutdown disables sends or receives on a socket.
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
// WARNING: Partially implemented, the condition and callbackData parameters are always ignored.
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
/*
func WSAGetLastError() int {
	ret, _, _ := procWSAGetLastError.Call()
	return int(int32(ret))
}
*/

// WSAGetOverlappedResult retrieves the results of an overlapped operation on the specified socket.
func WSAGetOverlappedResult(s SOCKET, overlapped *WSAOverlapped, transfer *uint32, wait bool, flags *uint32) (bool, error) {
	iWait := 0
	if wait {
		iWait = 1
	}
	ret, _, err := procWSAGetOverlappedResult.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(transfer)),
		uintptr(iWait),
		uintptr(unsafe.Pointer(flags)),
	)
	if int32(ret) == 0 {
		return false, err
	}
	return true, nil
}

// WSAGetQOSByName initializes a QOS structure based on a named template, or it supplies a buffer to retrieve
// an enumeration of the available template names.
/*
func WSAGetQOSByName(s SOCKET, qOSName *WSABuf, qOS *QOS) (bool, error) {
	ret, _, err := procWSAGetQOSByName.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(qOSName)),
		uintptr(unsafe.Pointer(qOS)),
	)
	if int32(ret) == 0 {
		return false, err
	}
	return true, nil
}
*/

// WSAGetServiceClassInfoA retrieves the class information (schema) pertaining to a specified service class from a specified namespace provider.
/*
func WSAGetServiceClassInfoA(
	providerId *GUID,
	serviceClassId *GUID,
	bufSize *uint32,
	serviceClassInfo *WSAServiceClassInfoA,
) (int, error) {
	ret, _, err := procWSAGetServiceClassInfoA.Call(
		uintptr(unsafe.Pointer(providerId)),
		uintptr(unsafe.Pointer(serviceClassId)),
		uintptr(unsafe.Pointer(&bufSize)),
		uintptr(unsafe.Pointer(serviceClassInfo)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAGetServiceClassInfoW retrieves the class information (schema) pertaining to a specified service class from a specified namespace provider.
/*
func WSAGetServiceClassInfoW(
	providerId *GUID,
	serviceClassId *GUID,
	bufSize *uint32,
	serviceClassInfo *WSAServiceClassInfoW,
) (int, error) {
	ret, _, err := procWSAGetServiceClassInfoW.Call(
		uintptr(unsafe.Pointer(providerId)),
		uintptr(unsafe.Pointer(serviceClassId)),
		uintptr(unsafe.Pointer(&bufSize)),
		uintptr(unsafe.Pointer(serviceClassInfo)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAGetServiceClassNameByClassIdA retrieves the name of the service associated with the specified type.
//
// This name is the generic service name, like FTP or SNA, and not the name of a specific instance of that service.
/*
func WSAGetServiceClassNameByClassIdA(serviceClassId *GUID, serviceClassName *byte, bufferLength *uint32) (int, error) {
	ret, _, err := procWSAGetServiceClassNameByClassIdA.Call(
		uintptr(unsafe.Pointer(serviceClassId)),
		uintptr(unsafe.Pointer(serviceClassName)),
		uintptr(unsafe.Pointer(bufferLength)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAGetServiceClassNameByClassIdW retrieves the name of the service associated with the specified type.
//
// This name is the generic service name, like FTP or SNA, and not the name of a specific instance of that service.
/*
func WSAGetServiceClassNameByClassIdW(serviceClassId *GUID, serviceClassName *uint16, bufferLength *uint32) (int, error) {
	ret, _, err := procWSAGetServiceClassNameByClassIdW.Call(
		uintptr(unsafe.Pointer(serviceClassId)),
		uintptr(unsafe.Pointer(serviceClassName)),
		uintptr(unsafe.Pointer(bufferLength)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAHtonl converts a unsigned long (uint32) from host byte order to network byte order.
func WSAHtonl(s SOCKET, hostLong uint32, netLong *uint32) (int, error) {
	ret, _, err := procWSAHtonl.Call(
		uintptr(s),
		uintptr(hostLong),
		uintptr(unsafe.Pointer(netLong)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAHtons converts a unsigned short (uint16) from host byte order to network byte order.
func WSAHtons(s SOCKET, hostShort uint16, netShort *uint16) (int, error) {
	ret, _, err := procWSAHtons.Call(
		uintptr(s),
		uintptr(hostShort),
		uintptr(unsafe.Pointer(netShort)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAInstallServiceClassA registers a service class schema within a namespace.
//
// This schema includes the class name, class identifier, and any namespace-specific information that is common to all
// instances of the service, such as the SAP identifier or object identifier.
/*
func WSAInstallServiceClassA(serviceClassInfo *WSAServiceClassInfoA) (int, error) {
	ret, _, err := procWSAInstallServiceClassA.Call(
		uintptr(unsafe.Pointer(serviceClassInfo)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAInstallServiceClassW registers a service class schema within a namespace.
//
// This schema includes the class name, class identifier, and any namespace-specific information that is common to all
// instances of the service, such as the SAP identifier or object identifier.
/*
func WSAInstallServiceClassW(serviceClassInfo *WSAServiceClassInfoW) (int, error) {
	ret, _, err := procWSAInstallServiceClassW.Call(
		uintptr(unsafe.Pointer(serviceClassInfo)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSAIoctl controls the mode of a socket.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSAIoctl(
	s SOCKET,
	ioControlCode uint32,
	inBuffer unsafe.Pointer,
	inBufferSize uint32,
	outBuffer unsafe.Pointer,
	outBufferSize uint32,
	bytesReturned *uint32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSAIoctl.Call(
		uintptr(s),
		uintptr(ioControlCode),
		uintptr(inBuffer),
		uintptr(inBufferSize),
		uintptr(outBuffer),
		uintptr(outBufferSize),
		uintptr(unsafe.Pointer(bytesReturned)),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAIsBlocking
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSAJoinLeaf joins a leaf node into a multipoint session, exchanges connect data,
// and specifies needed quality of service based on the specified FlowSpec structures.
func WSAJoinLeaf(
	s SOCKET,
	sockaddr unsafe.Pointer,
	nameLen int32,
	callerData *WSABuf,
	calleeData *WSABuf,
	sQOS *QOS,
	gQOS *QOS,
	flags uint32,
) (SOCKET, error) {
	ret, _, err := procWSAJoinLeaf.Call(
		uintptr(s),
		uintptr(sockaddr),
		uintptr(nameLen),
		uintptr(unsafe.Pointer(callerData)),
		uintptr(unsafe.Pointer(calleeData)),
		uintptr(unsafe.Pointer(sQOS)),
		uintptr(unsafe.Pointer(gQOS)),
		uintptr(flags),
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// WSALookupServiceBeginA initiates a client query that is constrained by the information contained within a WSAQuerySetA structure.
//
// WSALookupServiceBeginA only returns a handle, which should be used by subsequent calls to WSALookupServiceNext to get the actual results.
func WSALookupServiceBeginA(restrictions *WSAQuerySetA, controlFlags uint32, lookup *HANDLE) (int, error) {
	ret, _, err := procWSALookupServiceBeginA.Call(
		uintptr(unsafe.Pointer(restrictions)),
		uintptr(controlFlags),
		uintptr(unsafe.Pointer(lookup)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSALookupServiceBeginW initiates a client query that is constrained by the information contained within a WSAQuerySetW structure.
//
// WSALookupServiceBeginW only returns a handle, which should be used by subsequent calls to WSALookupServiceNext to get the actual results.
func WSALookupServiceBeginW(restrictions *WSAQuerySetW, controlFlags uint32, lookup *HANDLE) (int, error) {
	ret, _, err := procWSALookupServiceBeginW.Call(
		uintptr(unsafe.Pointer(restrictions)),
		uintptr(controlFlags),
		uintptr(unsafe.Pointer(lookup)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSALookupServiceEnd is called to free the handle after previous calls to WSALookupServiceBegin and WSALookupServiceNext.
//
// If you call WSALookupServiceEnd from another thread while an existing WSALookupServiceNext is blocked, the end call will have
// the same effect as a cancel and will cause the WSALookupServiceNext call to return immediately.
func WSALookupServiceEnd(lookup HANDLE) (int, error) {
	ret, _, err := procWSALookupServiceEnd.Call(
		uintptr(lookup),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSALookupServiceNextA is called after obtaining a handle from a previous call to WSALookupServiceBegin in order to retrieve the requested service information.
//
// The provider will pass back a WSAQuerySet structure in the results buffer.
//
// The client should continue to call this function until it returns WSA_E_NO_MORE, indicating that all of WSAQuerySet has been returned.
func WSALookupServiceNextA(lookup HANDLE, controlFlags uint32, bufferLength *uint32, results *WSAQuerySetA) (int, error) {
	ret, _, err := procWSALookupServiceNextA.Call(
		uintptr(lookup),
		uintptr(controlFlags),
		uintptr(unsafe.Pointer(&bufferLength)),
		uintptr(unsafe.Pointer(results)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSALookupServiceNextW is called after obtaining a handle from a previous call to WSALookupServiceBegin in order to retrieve the requested service information.
//
// The provider will pass back a WSAQuerySet structure in the results buffer.
//
// The client should continue to call this function until it returns WSA_E_NO_MORE, indicating that all of WSAQuerySet has been returned.
/*
func WSALookupServiceNextW(lookup HANDLE, controlFlags uint32, bufferLength *uint32, results *WSAQuerySetW) (int, error) {
	ret, _, err := procWSALookupServiceNextW.Call(
		uintptr(lookup),
		uintptr(controlFlags),
		uintptr(unsafe.Pointer(&bufferLength)),
		uintptr(unsafe.Pointer(results)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}
*/

// WSANSPIoctl enables developers to make I/O control calls to a registered namespace.
//
// WARNING: Partially implemented, the completion parameter is always ignored.
func WSANSPIoctl(
	lookup HANDLE,
	controlCode uint32,
	inBuffer unsafe.Pointer,
	inBufferSize uint32,
	outBuffer unsafe.Pointer,
	outBufferSize uint32,
	bytesReturned *uint32,
	completion unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSANSPIoctl.Call(
		uintptr(lookup),
		uintptr(controlCode),
		uintptr(inBuffer),
		uintptr(inBufferSize),
		uintptr(outBuffer),
		uintptr(outBufferSize),
		uintptr(unsafe.Pointer(&bytesReturned)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSANtohl converts a unsigned long (uint32) from network byte order to host byte order.
func WSANtohl(s SOCKET, netlong uint32, hostlong *uint32) (int, error) {
	ret, _, err := procWSANtohl.Call(
		uintptr(s),
		uintptr(netlong),
		uintptr(unsafe.Pointer(hostlong)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSANtohs converts a unsigned short (uint16) from network byte order to host byte order.
func WSANtohs(s SOCKET, netshort uint16, hostshort *uint16) (int, error) {
	ret, _, err := procWSANtohs.Call(
		uintptr(s),
		uintptr(netshort),
		uintptr(unsafe.Pointer(hostshort)),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAPoll determines status of one or more sockets.
func WSAPoll(fdArray []WSAPollFD, fds uint32, timeout int32) (int, error) {
	ret, _, err := procWSAPoll.Call(
		uintptr(unsafe.Pointer(&fdArray[0])),
		uintptr(fds),
		uintptr(timeout),
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSAProviderConfigChange notifies the application when the provider configuration is changed.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSAProviderConfigChange(
	notificationHandle *HANDLE,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSAProviderConfigChange.Call(
		uintptr(unsafe.Pointer(notificationHandle)),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSARecv receives data from a connected socket or a bound connectionless socket.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSARecv(
	s SOCKET,
	buffers *WSABuf,
	bufferCount uint32,
	numberOfBytesRecvd *uint32,
	flags *uint32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSARecv.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(buffers)),
		uintptr(bufferCount),
		uintptr(unsafe.Pointer(numberOfBytesRecvd)),
		uintptr(unsafe.Pointer(flags)),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	iRet := int32(ret)
	if iRet == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return int(iRet), nil
}

// WSARecvDisconnect terminates reception on a socket, and retrieves the disconnect data if the socket is connection oriented.
func WSARecvDisconnect(s SOCKET, inboundDisconnectData *WSABuf) (int, error) {
	ret, _, err := procWSARecvDisconnect.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(inboundDisconnectData)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSARecvFrom receives a datagram and stores the source address.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSARecvFrom(
	s SOCKET,
	buffers *WSABuf,
	bufferCount uint32,
	numberOfBytesRecvd *uint32,
	flags *uint32,
	from unsafe.Pointer,
	fromLen *int32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSARecvFrom.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(buffers)),
		uintptr(bufferCount),
		uintptr(unsafe.Pointer(numberOfBytesRecvd)),
		uintptr(unsafe.Pointer(flags)),
		uintptr(from),
		uintptr(unsafe.Pointer(fromLen)),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSARemoveServiceClass permanently removes the service class schema from the registry.
func WSARemoveServiceClass(serviceClassId *GUID) (int, error) {
	ret, _, err := procWSARemoveServiceClass.Call(
		uintptr(unsafe.Pointer(serviceClassId)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAResetEvent resets the state of the specified event object to nonsignaled.
func WSAResetEvent(event WSAEVENT) (bool, error) {
	ret, _, err := procWSAResetEvent.Call(
		uintptr(event),
	)
	if int32(ret) == 0 {
		return false, err
	}
	return true, nil
}

// WSASend sends data on a connected socket.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSASend(
	s SOCKET,
	buffers *WSABuf,
	bufferCount uint32,
	numberOfBytesSent *uint32,
	flags uint32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSASend.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(buffers)),
		uintptr(bufferCount),
		uintptr(unsafe.Pointer(&numberOfBytesSent)),
		uintptr(flags),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASendDisconnect initiates termination of the connection for the socket and sends disconnect data.
func WSASendDisconnect(s SOCKET, inboundDisconnectData *WSABuf) (int, error) {
	ret, _, err := procWSASendDisconnect.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(inboundDisconnectData)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASendMsg sends data and optional control information from connected and unconnected sockets.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSASendMsg(
	socket SOCKET,
	msg *WSAMsg,
	flags uint32,
	numberOfBytesSent *uint32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSASendMsg.Call(
		uintptr(socket),
		uintptr(unsafe.Pointer(&msg)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&numberOfBytesSent)),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASendTo sends data to a specific destination, using overlapped I/O where applicable.
//
// WARNING: Partially implemented, the completionRoutine parameter is always ignored.
func WSASendTo(
	s SOCKET,
	buffers *WSABuf,
	bufferCount uint32,
	numberOfBytesSent *uint32,
	flags uint32,
	to unsafe.Pointer,
	tolen int32,
	overlapped *WSAOverlapped,
	completionRoutine unsafe.Pointer,
) (int, error) {
	ret, _, err := procWSASendTo.Call(
		uintptr(s),
		uintptr(unsafe.Pointer(buffers)),
		uintptr(bufferCount),
		uintptr(unsafe.Pointer(numberOfBytesSent)),
		uintptr(flags),
		uintptr(to),
		uintptr(tolen),
		uintptr(unsafe.Pointer(overlapped)),
		uintptr(unsafe.Pointer(nil)), // unimplemented
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASetBlockingHook
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSASetEvent sets the state of the specified event object to signaled.
func WSASetEvent(event WSAEVENT) (bool, error) {
	ret, _, err := procWSASetEvent.Call(
		uintptr(event),
	)
	if int32(ret) == 0 {
		return false, nil
	}
	return true, err
}

// WSASetLastError sets the error code that can be retrieved through the WSAGetLastError function.
func WSASetLastError(errorCode int) error {
	_, _, err := procWSASetLastError.Call(
		uintptr(errorCode),
	)
	if isValidErr(err) {
		return err
	}
	return nil
}

// WSASetServiceA registers or removes from the registry a service instance within one or more namespaces.
func WSASetServiceA(regInfo *WSAQuerySetA, essoperation WSAESETSERVICEOP, controlFlags uint32) (int, error) {
	ret, _, err := procWSASetServiceA.Call(
		uintptr(unsafe.Pointer(regInfo)),
		uintptr(unsafe.Pointer(&essoperation)),
		uintptr(controlFlags),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASetServiceW registers or removes from the registry a service instance within one or more namespaces.
func WSASetServiceW(regInfo *WSAQuerySetW, essoperation WSAESETSERVICEOP, controlFlags uint32) (int, error) {
	ret, _, err := procWSASetServiceW.Call(
		uintptr(unsafe.Pointer(regInfo)),
		uintptr(unsafe.Pointer(&essoperation)),
		uintptr(controlFlags),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSASocketA creates a socket that is bound to a specific transport-service provider.
func WSASocketA(af int32, stype int32, protocol int32, protocolInfo *WSAProtocolInfoA, g GROUP, flags uint32) (SOCKET, error) {
	ret, _, err := procWSASocketA.Call(
		uintptr(af),
		uintptr(stype),
		uintptr(protocol),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(g),
		uintptr(flags),
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// WSASocketW creates a socket that is bound to a specific transport-service provider.
func WSASocketW(af int32, stype int32, protocol int32, protocolInfo *WSAProtocolInfoW, g GROUP, flags uint32) (SOCKET, error) {
	ret, _, err := procWSASocketW.Call(
		uintptr(af),
		uintptr(stype),
		uintptr(protocol),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(g),
		uintptr(flags),
	)
	sock := SOCKET(ret)
	if sock == INVALID_SOCKET {
		return INVALID_SOCKET, err
	}
	return sock, nil
}

// WSAStartup initiates use of the Winsock DLL by a process.
func WSAStartup(versionRequired uint16, wsData *WSAData) (int, error) {
	ret, _, err := procWSAStartup.Call(
		uintptr(versionRequired),
		uintptr(unsafe.Pointer(wsData)),
	)
	iRet := int32(ret)
	if iRet != 0 {
		return int(iRet), err
	}
	return 0, nil
}

// WSAStringToAddressA converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure,
// suitable for passing to Windows Sockets routines that take such a structure.
func WSAStringToAddressA(
	addressString *byte,
	addressFamily uint16,
	protocolInfo *WSAProtocolInfoA,
	address unsafe.Pointer,
	addressLength *int32,
) (int, error) {
	ret, _, err := procWSAStringToAddressA.Call(
		uintptr(unsafe.Pointer(addressString)),
		uintptr(addressFamily),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(address),
		uintptr(unsafe.Pointer(addressLength)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAStringToAddressW converts a network address in its standard text presentation form into its numeric binary form in a sockaddr structure,
// suitable for passing to Windows Sockets routines that take such a structure.
func WSAStringToAddressW(
	addressString *uint16,
	addressFamily uint16,
	protocolInfo *WSAProtocolInfoW,
	address unsafe.Pointer,
	addressLength *int32,
) (int, error) {
	ret, _, err := procWSAStringToAddressW.Call(
		uintptr(unsafe.Pointer(addressString)),
		uintptr(addressFamily),
		uintptr(unsafe.Pointer(protocolInfo)),
		uintptr(address),
		uintptr(unsafe.Pointer(addressLength)),
	)
	if int32(ret) == SOCKET_ERROR {
		return SOCKET_ERROR, err
	}
	return 0, nil
}

// WSAUnhookBlockingHook
// This function has been removed in compliance with the Windows Sockets 2 specification, revision 2.2.0.

// WSAWaitForMultipleEvents returns when one or all of the specified event objects are in the signaled state,
// when the time-out interval expires, or when an I/O completion routine has executed.
func WSAWaitForMultipleEvents(eventsCount uint32, events *WSAEVENT, waitAll bool, timeout uint32, alertable bool) (int, error) {
	iWaitAll := 0
	iAlertable := 0

	if waitAll {
		iWaitAll = 1
	}
	if alertable {
		iAlertable = 1
	}
	ret, _, err := procWSAWaitForMultipleEvents.Call(
		uintptr(eventsCount),
		uintptr(unsafe.Pointer(events)),
		uintptr(iWaitAll),
		uintptr(timeout),
		uintptr(iAlertable),
	)
	if !isValidErr(err) {
		err = nil
	}
	return int(ret), err
}

// isValidErr returns true if the given error is a "real" one.
func isValidErr(err error) bool {
	if err != nil && !strings.Contains(err.Error(), "success") {
		return true
	}
	return false
}
