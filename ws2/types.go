package ws2

import (
	"fmt"
	"net"
	"unsafe"
)

// Basic types.
type (
	SOCKET           uintptr
	HANDLE           uintptr
	WSAEVENT         HANDLE
	GROUP            uint32
	SERVICETYPE      uint32
	WSAECOMPARATOR   int32
	WSAESETSERVICEOP int32
)

// FDSet is used by various Windows Sockets functions and service providers, such as the select function,
// to place sockets into a "set" for various purposes, such as testing a given socket for readability using the readfds parameter of the select function.
type FDSet struct {
	Count uint32
	Array [FD_SETSIZE]SOCKET
}

// Clear removes the specified socket from the FDSet.
// It iterates through the array, finds the socket, and removes it by shifting the remaining sockets.
func (fd *FDSet) Clear(s SOCKET) {
	for i := 0; i < int(fd.Count); i++ {
		if fd.Array[i] == s {
			for j := i; j < int(fd.Count)-1; j++ {
				fd.Array[j] = fd.Array[j+1]
			}
			fd.Count--
			break
		}
	}
}

// Set adds the specified socket to the FDSet if it is not already present.
// It ensures that the socket is not added if the set is already at its maximum capacity (FD_SETSIZE).
func (fd *FDSet) Set(s SOCKET) {
	for i := 0; i < int(fd.Count); i++ {
		if fd.Array[i] == s {
			return
		}
	}
	if int(fd.Count) < FD_SETSIZE {
		fd.Array[fd.Count] = s
		fd.Count++
	}
}

// IsSet checks whether the specified socket is present in the FDSet.
// It calls the helper function __WSAFDIsSet to perform the check.
func (fd *FDSet) IsSet(s SOCKET) bool {
	ret, _ := __WSAFDIsSet(s, fd)
	return ret
}

// Zero clears all entries from the FDSet by setting the count to zero.
// This effectively removes all sockets from the set.
func (fd *FDSet) Zero() {
	fd.Count = 0
}

// Timeval is used to specify a time interval.
//
// It is associated with the Berkeley Software Distribution (BSD) Time.h header file.
type Timeval struct {
	Sec  int32 // seconds
	USec int32 // microseconds
}

// InAddr represents an IPv4 Internet address.
type InAddr struct {
	Addr [4]byte // IPv4 address.
}

// InAddr6 represents an IPv6 Internet address.
type InAddr6 struct {
	Addr [16]byte // IPv6 address.
}

// SockAddr varies depending on the protocol selected.
type SockAddr struct {
	Family uint16   // Address family (AF_INET/AF_INET6).
	Data   [64]byte // Data.
}

// Is4 checks if the socket address represents an IPv4 address.
func (sa *SockAddr) Is4() bool {
	return sa.Family == AF_INET && len(sa.Data) >= net.IPv4len
}

// Is6 checks if the socket address represents an IPv6 address.
func (sa *SockAddr) Is6() bool {
	return sa.Family == AF_INET6 && len(sa.Data) >= net.IPv6len
}

// To4 returns the socket address converted to an SockAddrInet4 structure or nil.
func (sa *SockAddr) To4() (v4 *SockAddrInet4) {
	v4, _ = sa.toInets()
	return
}

// To6 returns the socket address converted to an SockAddrInet6 structure or nil.
func (sa *SockAddr) To6() (v6 *SockAddrInet6) {
	_, v6 = sa.toInets()
	return
}

// ToIP returns the net.IP representation of the socket IP address.
func (sa *SockAddr) ToIP() (ip net.IP) {
	ip, _ = sa.ToIPPort()
	return
}

// ToIPPort returns the IP as net.IP and port number.
func (sa *SockAddr) ToIPPort() (ip net.IP, port int) {
	if sa.Is4() {
		addr := sa.To4()
		if addr != nil {
			ip = net.IPv4(addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3])
			port = int(Ntohs(addr.Port))
		}
	} else if sa.Is6() {
		addr := sa.To6()
		if addr != nil {
			ip = net.IP(addr.Addr[:16])
			port = int(Ntohs(addr.Port))
		}
	}
	return
}

// FromNetIP populates the SockAddr structure with the given net.IP and port number.
func (sa *SockAddr) FromNetIP(ip net.IP, port int) error {
	if port < 0 || port > 0xFFFF {
		return fmt.Errorf("port out of range (0-65535)")
	}

	var ipLen int = 0
	var ipOffset int = 2 // v4
	var ipBytes []byte

	if v4 := ip.To4(); v4 != nil {
		sa.Family = AF_INET
		ipLen = net.IPv4len
		ipBytes = v4
	} else {
		sa.Family = AF_INET6
		ipLen = net.IPv6len
		ipOffset = 6
		ipBytes = ip.To16()
	}

	if ipLen == 0 {
		return fmt.Errorf("invalid IP")
	}

	var bePort uint16 = Htons(uint16(port))
	var data [64]byte

	data[0] = byte(uint16(bePort & 0xFF))
	data[1] = byte(uint16(bePort >> 8))

	copy(data[ipOffset:], ipBytes)

	sa.Data = data
	return nil
}

// toInets converts the SockAddr to SockAddrInet4 or SockAddrInet6 based on the address family.
func (sa *SockAddr) toInets() (v4 *SockAddrInet4, v6 *SockAddrInet6) {
	if sa.Is4() {
		in := SockAddrInet4{
			Family: AF_INET,
			Port:   uint16(sa.Data[1]) | uint16(sa.Data[0])<<8,
			Addr:   [4]byte(sa.Data[2:6]),
		}
		v4 = &in
	} else if sa.Is6() {
		in := SockAddrInet6{
			Family: AF_INET6,
			Port:   uint16(sa.Data[1]) | uint16(sa.Data[0])<<8,
			Addr:   [16]byte(sa.Data[2:18]),
		}
		v6 = &in
	}
	return
}

// SockAddrInet4 represents an IPv4 socket address structure.
//
// The address family is in host byte order and the IPv6 address is in network byte order.
type SockAddrInet4 struct {
	Family uint16   // Address family (AF_INET).
	Port   uint16   // Port number.
	Addr   [4]byte  // IPv4 address.
	_      [8]uint8 // Padding.
}

// ToIPPort returns the IP as net.IP and port number.
func (sa *SockAddrInet4) ToIPPort() (ip net.IP, port int) {
	ip = net.IPv4(sa.Addr[0], sa.Addr[1], sa.Addr[2], sa.Addr[3])

	if sa.Port > 0 {
		port = int(Ntohs(sa.Port))
	}
	return
}

// SockAddrInet6 represents an IPv6 socket address structure.
//
// The address family is in host byte order and the IPv6 address is in network byte order.
type SockAddrInet6 struct {
	Family   uint16   // Address family (AF_INET6).
	Port     uint16   // Port number.
	FlowInfo uint32   // IPv6 flow information.
	Addr     [16]byte // IPv6 address.
	ScopeID  uint32   // Scope identifier.
}

// ToIPPort returns the IP as net.IP and port number.
func (sa *SockAddrInet6) ToIPPort() (ip net.IP, port int) {
	ip = net.IP(sa.Addr[:])

	if sa.Port > 0 {
		port = int(Ntohs(sa.Port))
	}
	return
}

// HostEnt is used by functions to store information about a given host, such as host name, IPv4 address, and so forth.
// An application should never attempt to modify this structure or to free any of its components.
//
// Furthermore, only one copy of the hostent structure is allocated per thread, and an application should therefore copy any
// information that it needs before issuing any other Windows Sockets API calls.
type HostEnt struct {
	HName     *byte  //  Official name of host
	HAliases  **byte //  Alias list
	HAddrType int16  //  Host address type
	HLength   int16  //  Length of address
	HAddrList **byte //  List of addresses
}

// It is assumed here that a network number
// fits in 32 bits.
type NetEnt struct {
	NName     *byte  //  Official name of net
	NAliases  **byte //  Alias list
	NAddrType int16  //  Net address type
	NNet      uint32 //  Network #
}

// ServEnt is used to store or return the name and service number for a given service name.
type ServEnt struct {
	SName    *byte  //  Official service name
	SAliases **byte //  Alias list
	SProto   *byte  //  Protocol to use
	SPort    int16  //  Port #
}

// ProtoEnt  contains the name and protocol numbers that correspond to a given protocol name.
// Applications must never attempt to modify this structure or to free any of its components.
//
// Furthermore, only one copy of this structure is allocated per thread, and therefore, the application should copy any
// information it needs before issuing any other Windows Sockets function calls.
type ProtoEnt struct {
	PName    *byte  //  Official protocol name
	PAliases **byte //  Alias list
	PProto   int16  //  Protocol #
}

// AddrInfoA is used by the GetAddrInfoA function to hold host address information.
type AddrInfoA struct {
	Flags      int32          // AI_PASSIVE, AI_CANONNAME, AI_NUMERICHOST.
	Family     int32          // PF_xxx.
	SockType   int32          // SOCK_xxx.
	Protocol   int32          // 0 or IPPROTO_xxx for IPv4 and IPv6.
	AddrLength uint32         // Length of Addr.
	CanonName  *byte          // Canonical name for nodename.
	Addr       unsafe.Pointer // Binary address.
	Next       *AddrInfoA     // Next structure in linked list.
}

// AddrInfoW is used by the GetAddrInfoW function to hold host address information.
type AddrInfoW struct {
	Flags      int32          // AI_PASSIVE, AI_CANONNAME, AI_NUMERICHOST.
	Family     int32          // PF_xxx.
	SockType   int32          // SOCK_xxx.
	Protocol   int32          // 0 or IPPROTO_xxx for IPv4 and IPv6.
	AddrLength uint32         // Length of Addr.
	CanonName  *uint16        // Canonical name for nodename.
	Addr       unsafe.Pointer // Binary address.
	Next       *AddrInfoW     // Next structure in linked list.
}

// BLOB represents a structure derived from Binary Large Object, contains information about a block of data.
type BLOB struct {
	Data []byte // Binary data.
	Size uint32 // Data length.
}

// FlowSpec  provides quality of service parameters to the RSVP SP.
//
// This allows QOS-aware applications to invoke, modify, or remove QOS settings for a given flow.
type FlowSpec struct {
	TokenRate          uint32      // Rate at which data can be transmitted over the life of the flow
	TokenBucketSize    uint32      // Maximum amount of credits a given direction of a flow can accrue, regardless of time, in bytes.
	PeakBandwidth      uint32      // Upper limit on time-based transmission permission for a given flow, in bytes per second.
	Latency            uint32      // Maximum acceptable delay between transmission of a bit by the sender and its receipt by one or more intended receivers, in microseconds.
	DelayVariation     uint32      // Difference between the maximum and minimum possible delay a packet will experience, in microseconds.
	ServiceType        SERVICETYPE // Specifies the level of service to negotiate for the flow.
	MaxSduSize         uint32      // Specifies the maximum packet size permitted or used in the traffic flow, in bytes.
	MinimumPolicedSize uint32      // Specifies the minimum packet size for which the requested quality of service will be provided, in bytes.
}

// GUID identifies an object such as a COM interfaces, or a COM class object, or a manager entry-point vector(EPV).
//
// A GUID is a 128-bit value consisting of one group of 8 hexadecimal digits, followed by three groups of 4 hexadecimal digits each,
// followed by one group of 12 hexadecimal digits.
//
// The GUID structure stores a GUID.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// AFProtocols supplies a list of protocols to which application programmers can constrain queries.
//
// The AFPROTOCOLS structure is used for query purposes only.
type AFProtocols struct {
	AddressFamily int32 // Address family to which the query is to be constrained.
	Protocol      int32 // Protocol to which the query is to be constrained.
}

// CSAddrInfo contains Windows Sockets address information for a socket, network service, or namespace provider.
type CSAddrInfo struct {
	LocalAddr  SocketAddress // Specifies a Windows Sockets local address.
	RemoteAddr SocketAddress // Specifies a Windows Sockets remote address.
	SocketType int32         // Specifies the type of the Windows Socket.
	Protocol   int32         // Specifies a value to pass as the protocol parameter to the socket (Windows Sockets) function to open a socket for this service.
}

// Overlapped contains information used in asynchronous (or overlapped) input and output (I/O).
type Overlapped struct {
	Internal     uint32   // Reserved for internal use. The Internal member is used internally by the entity that implements overlapped I/O.
	InternalHigh uint32   // Reserved. Used internally by the entity that implements overlapped I/O.
	Offset       uint32   // Reserved for use by service providers.
	OffsetHigh   uint32   // Reserved for use by service providers.
	HEvent       WSAEVENT // Either contain a valid handle to a WSAEVENT object or be null.
}

// SocketAddress stores protocol-specific address information.
type SocketAddress struct {
	Sockaddr       unsafe.Pointer // Pointer to a socket address represented as a SockAddr structure.
	SockaddrLength int32          // length, in bytes, of the socket address.
}

// SocketAddressList stores an array of SocketAddress structures that contain protocol-specific address information.
type SocketAddressList struct {
	AddressCount int32           // Number of address structures in the Address member.
	Address      []SocketAddress // Slice of SocketAddress structures that are specific to a protocol family.
}

// QOS provides the means by which QOS-enabled applications can specify quality of service parameters for sent and received traffic on a particular flow.
type QOS struct {
	SendingFlowspec   FlowSpec // Specifies QOS parameters for the sending direction of a particular flow.
	ReceivingFlowspec FlowSpec // Specifies QOS parameters for the receiving direction of a particular flow.
	ProviderSpecific  WSABuf   // Pointer to a structure of type WSABUF that can provide additional provider-specific quality of service parameters to the RSVP SP for a given flow.
}

// WSAData contains information about the Windows Sockets implementation.
type WSAData struct {
	Version      uint16                       // Version of the Windows Sockets specification that the Ws2_32.dll expects the caller to use.
	HighVersion  uint16                       // Highest version of the Windows Sockets specification that the Ws2_32.dll can support.
	Description  [WSADESCRIPTION_LEN + 1]byte // NULL-terminated ASCII string into which the Ws2_32.dll copies a description of the Windows Sockets implementation.
	SystemStatus [WSASYS_STATUS_LEN + 1]byte  // NULL-terminated ASCII string into which the Ws2_32.dll copies relevant status or configuration information.
	// MaxSockets   uint16                       // Maximum number of sockets that may be opened. (ignored)
	// MaxUdpDg     uint16                       // The maximum datagram message size. (ignored)
	// VendorInfo   *byte                        // Pointer to vendor-specific information. (ignored)
}

// WSANameSpaceInfoA contains all registration information for a namespace provider.
type WSANameSpaceInfoA struct {
	NSProviderId GUID   // Unique GUID for this namespace provider.
	NameSpace    uint32 // Namespace supported by this provider.
	Active       bool   // If TRUE, indicates that this namespace provider is active.
	Version      uint32 // Version number of the namespace provider.
	Identifier   *byte  // Display string that identifies the namespace provider.
}

// WSANameSpaceInfoExA contains all registration information for a namespace provider.
type WSANameSpaceInfoExA struct {
	NSProviderId     GUID   // Unique GUID for this namespace provider.
	NameSpace        uint32 // Namespace supported by this provider.
	Active           bool   // If TRUE, indicates that this namespace provider is active.
	Version          uint32 // Version number of the namespace provider.
	Identifier       *byte  // Display string that identifies the namespace provider.
	ProviderSpecific BLOB   // Provider-specific data blob associated with namespace entry.
}

// WSANameSpaceInfoExW contains all registration information for a namespace provider.
type WSANameSpaceInfoExW struct {
	NSProviderId     GUID    // Unique GUID for this namespace provider.
	NameSpace        uint32  // Namespace supported by this provider.
	Active           bool    // If TRUE, indicates that this namespace provider is active.
	Version          uint32  // Version number of the namespace provider.
	Identifier       *uint16 // Display string that identifies the namespace provider.
	ProviderSpecific BLOB    // Provider-specific data blob associated with namespace entry.
}

// WSANameSpaceInfoW contains all registration information for a namespace provider.
type WSANameSpaceInfoW struct {
	NSProviderId GUID    // Unique GUID for this namespace provider.
	NameSpace    uint32  // Namespace supported by this provider.
	Active       bool    // If TRUE, indicates that this namespace provider is active.
	Version      uint32  // Version number of the namespace provider.
	Identifier   *uint16 // Display string that identifies the namespace provider.
}

// WSANetworkEvents is used to store a socket's internal information about network events.
type WSANetworkEvents struct {
	NetworkEvents int32                // Indicates which of the FD_XXX network events have occurred.
	ErrorCode     [FD_MAX_EVENTS]int32 // Array that contains any associated error codes.
}

// WSANSClassInfoA provides individual parameter information for a specific Windows Sockets namespace.
type WSANSClassInfoA struct {
	Name      *byte          // String value associated with the parameter, such as SAPID, TCPPORT, and so forth.
	NameSpace uint32         // GUID associated with the namespace.
	ValueType uint32         // Value type for the parameter, such as REG_DWORD or REG_SZ, and so forth.
	ValueSize uint32         // Size of the parameter provided in Value, in bytes.
	Value     unsafe.Pointer // Pointer to the value of the parameter.
}

// WSANSClassInfoW provides individual parameter information for a specific Windows Sockets namespace.
type WSANSClassInfoW struct {
	Name      *uint16        // String value associated with the parameter, such as SAPID, TCPPORT, and so forth.
	NameSpace uint32         // GUID associated with the namespace.
	ValueType uint32         // Value type for the parameter, such as REG_DWORD or REG_SZ, and so forth.
	ValueSize uint32         // Size of the parameter provided in Value, in bytes.
	Value     unsafe.Pointer // Pointer to the value of the parameter.
}

// WSABuf enables the creation or manipulation of a data buffer used by some Winsock functions.
type WSABuf struct {
	Length uint32 // Length of the buffer, in bytes.
	Buf    *byte  // Pointer to the buffer.
}

// WSAMsg is used with the WSARecvMsg and WSASendMsg functions to store address and optional control
// information about connected and unconnected sockets as well as an array of buffers used to store message data.
type WSAMsg struct {
	Name        unsafe.Pointer // Pointer to a SocketAddress structure that stores information about the remote address. Used only with unconnected sockets.
	NameLength  int32          // Length, in bytes, of the SocketAddress structure pointed to in the pAddr member. Used only with unconnected sockets.
	Buffers     *WSABuf        // Array of WSABUF structures used to receive the message data. The capability of the Buffers member to contain multiple buffers enables the use of scatter/gather I/O.
	BufferCount uint32         // Number of buffers pointed to in the Buffers member.
	Control     WSABuf         // A structure of WSABUF type used to specify optional control data.
	Flags       uint32         // One or more control flags, specified as the logical OR of values.
}

// WSAPollFD stores socket information used by the WSAPoll function.
type WSAPollFD struct {
	FD      SOCKET // Identifier of the socket for which to find status. This parameter is ignored if set to a negative value.
	Events  int16  // Set of flags indicating the type of status being requested.
	Revents int16  // Set of flags that indicate, upon return from the WSAPoll function call, the results of the status query.
}

// WSAProtocolChain contains a counted list of Catalog Entry identifiers that comprise a protocol chain.
type WSAProtocolChain struct {
	ChainLength  int32   // Length of the chain, in bytes. (0 = layered protocol, 1 = base protocol, >1 = protocol chain)
	ChainEntries *uint32 // Array of protocol chain entries. /*[MAX_PROTOCOL_CHAIN]uint32*/
}

// WSAProtocolInfoA is used to store or retrieve complete information for a given protocol.
// The protocol name is represented as an array of ANSI characters.
type WSAProtocolInfoA struct {
	ServiceFlags1     uint32                    // Bitmask that describes the services provided by the protocol.
	ServiceFlags2     uint32                    // Reserved for additional protocol-attribute definitions.
	ServiceFlags3     uint32                    // Reserved for additional protocol-attribute definitions.
	ServiceFlags4     uint32                    // Reserved for additional protocol-attribute definitions.
	ProviderFlags     uint32                    // Set of flags that provides information on how this protocol is represented in the Winsock catalog.
	ProviderId        GUID                      // A globally unique identifier (GUID) assigned to the provider by the service provider vendor.
	CatalogEntryId    uint32                    // Unique identifier assigned by the WS2_32.DLL for each WSAProtocolInfo structure.
	ProtocolChain     WSAProtocolChain          // The WSAProtocolChain structure associated with the protocol.
	Version           int32                     // Protocol version identifier.
	AddressFamily     int32                     // Value to pass as the address family parameter to the socket or WSASocket function in order to open a socket for this protocol.
	MaxSockAddr       int32                     // Maximum address length, in bytes.
	MinSockAddr       int32                     // Minimum address length, in bytes.
	SocketType        int32                     // Value to pass as the socket type parameter to the socket or WSASocket function in order to open a socket for this protocol.
	Protocol          int32                     // Value to pass as the protocol parameter to the socket or WSASocket function in order to open a socket for this protocol.
	ProtocolMaxOffset int32                     // Maximum value that may be added to iProtocol when supplying a value for the protocol parameter to socket or WSASocket function.
	NetworkByteOrder  int32                     // Currently these values are manifest constants (BIGENDIAN and LITTLEENDIAN) that indicate either big-endian or little-endian with the values 0 and 1 respectively.
	SecurityScheme    int32                     // Type of security scheme employed (if any). A value of SECURITY_PROTOCOL_NONE (0) is used for protocols that do not incorporate security provisions.
	MessageSize       uint32                    // Maximum message size, in bytes, supported by the protocol.
	ProviderReserved  uint32                    // Reserved for use by service providers.
	ProtocolName      [WSAPROTOCOL_LEN + 1]byte // Array of characters that contains a human-readable name identifying the protocol (max. 255 char.).
}

// WSAProtocolInfoW is used to store or retrieve complete information for a given protocol.
// The protocol name is represented as an array of Unicode characters.
type WSAProtocolInfoW struct {
	ServiceFlags1     uint32                      // Bitmask that describes the services provided by the protocol.
	ServiceFlags2     uint32                      // Reserved for additional protocol-attribute definitions.
	ServiceFlags3     uint32                      // Reserved for additional protocol-attribute definitions.
	ServiceFlags4     uint32                      // Reserved for additional protocol-attribute definitions.
	ProviderFlags     uint32                      // Set of flags that provides information on how this protocol is represented in the Winsock catalog.
	ProviderId        GUID                        // A globally unique identifier (GUID) assigned to the provider by the service provider vendor.
	CatalogEntryId    uint32                      // Unique identifier assigned by the WS2_32.DLL for each WSAProtocolInfo structure.
	ProtocolChain     WSAProtocolChain            // The WSAProtocolChain structure associated with the protocol.
	Version           int32                       // Protocol version identifier.
	AddressFamily     int32                       // Value to pass as the address family parameter to the socket or WSASocket function in order to open a socket for this protocol.
	MaxSockAddr       int32                       // Maximum address length, in bytes.
	MinSockAddr       int32                       // Minimum address length, in bytes.
	SocketType        int32                       // Value to pass as the socket type parameter to the socket or WSASocket function in order to open a socket for this protocol.
	Protocol          int32                       // Value to pass as the protocol parameter to the socket or WSASocket function in order to open a socket for this protocol.
	ProtocolMaxOffset int32                       // Maximum value that may be added to iProtocol when supplying a value for the protocol parameter to socket or WSASocket function.
	NetworkByteOrder  int32                       // Currently these values are manifest constants (BIGENDIAN and LITTLEENDIAN) that indicate either big-endian or little-endian with the values 0 and 1 respectively.
	SecurityScheme    int32                       // Type of security scheme employed (if any). A value of SECURITY_PROTOCOL_NONE (0) is used for protocols that do not incorporate security provisions.
	MessageSize       uint32                      // Maximum message size, in bytes, supported by the protocol.
	ProviderReserved  uint32                      // Reserved for use by service providers.
	ProtocolName      [WSAPROTOCOL_LEN + 1]uint16 // Array of characters that contains a human-readable name identifying the protocol (max. 255 char.).
}

// WSAQuerySetA provides relevant information about a given service, including service class ID, service name,
// applicable namespace identifier and protocol information, as well as a set of transport addresses at which the service listens.
type WSAQuerySetA struct {
	Size                uint32       // Size, in bytes, of the WSAQuerySetA structure.
	ServiceInstanceName *byte        // Pointer to an optional NULL-terminated string that contains service name.
	ServiceClassId      *GUID        // GUID corresponding to the service class. This member is required to be set.
	Version             *WSAVersion  // Pointer to an optional desired version number of the namespace provider.
	Comment             *byte        // Ignored for queries.
	NameSpace           uint32       // Namespace identifier that determines which namespace providers are queried.
	NSProviderId        *GUID        // Pointer to an optional GUID of a specific namespace provider to query in the case where multiple namespace providers are registered under a single namespace such as NS_DNS.
	Context             *byte        // Pointer to an optional starting point of the query in a hierarchical namespace.
	NumberOfProtocols   uint32       // Size, in bytes, of the protocol constraint array. This member can be zero.
	AFProtocols         *AFProtocols // Pointer to an optional array of AFProtocols structures. Only services that utilize these protocols will be returned.
	QueryString         *byte        // Pointer to an optional NULL-terminated query string.
	NumberOfCsAddrs     uint32       // Ignored for queries.
	CSABuffer           *CSAddrInfo  // Ignored for queries.
	OutputFlags         uint32       // Ignored for queries.
	Blob                *WSABuf      // Optional pointer to data that is used to query or set provider-specific namespace information.
}

// WSAQuerySetW provides relevant information about a given service, including service class ID, service name,
// applicable namespace identifier and protocol information, as well as a set of transport addresses at which the service listens.
type WSAQuerySetW struct {
	Size                uint32       // Size, in bytes, of the WSAQuerySetW structure.
	ServiceInstanceName *uint16      // Pointer to an optional NULL-terminated string that contains service name.
	ServiceClassId      *GUID        // GUID corresponding to the service class. This member is required to be set.
	Version             *WSAVersion  // Pointer to an optional desired version number of the namespace provider.
	Comment             *uint16      // Ignored for queries.
	NameSpace           uint32       // Namespace identifier that determines which namespace providers are queried.
	NSProviderId        *GUID        // Pointer to an optional GUID of a specific namespace provider to query in the case where multiple namespace providers are registered under a single namespace such as NS_DNS.
	Context             *uint16      // Pointer to an optional starting point of the query in a hierarchical namespace.
	NumberOfProtocols   uint32       // Size, in bytes, of the protocol constraint array. This member can be zero.
	AFProtocols         *AFProtocols // Pointer to an optional array of AFProtocols structures. Only services that utilize these protocols will be returned.
	QueryString         *uint16      // Pointer to an optional NULL-terminated query string.
	NumberOfCsAddrs     uint32       // Ignored for queries.
	CSABuffer           *CSAddrInfo  // Ignored for queries.
	OutputFlags         uint32       // Ignored for queries.
	Blob                *WSABuf      // Optional pointer to data that is used to query or set provider-specific namespace information.
}

// WSAOverlapped provides a communication medium between the initiation of an overlapped I/O operation and its subsequent completion.
//
// The WSAOverlapped structure is compatible with the Windows Overlapped structure.
type WSAOverlapped struct {
	Overlapped
}

// WSAServiceClassInfoA contains information about a specified service class.
type WSAServiceClassInfoA struct {
	ServiceClassId   *GUID            // Unique Identifier (GUID) for the service class.
	ServiceClassName *byte            // Well known name associated with the service class.
	Count            uint32           // Number of entries in ClassInfos.
	ClassInfos       *WSANSClassInfoA // Array of WSANSClassInfoA structures that contains information about the service class.
}

// WSAServiceClassInfoW contains information about a specified service class.
type WSAServiceClassInfoW struct {
	ServiceClassId   *GUID            // Unique Identifier (GUID) for the service class.
	ServiceClassName *uint16          // Well known name associated with the service class.
	Count            uint32           // Number of entries in ClassInfos.
	ClassInfos       *WSANSClassInfoW // Array of WSANSClassInfoW structures that contains information about the service class.
}

// WSAVersion provides version comparison in Windows Sockets.
type WSAVersion struct {
	Version uint32         // Version of Windows Sockets.
	ECHow   WSAECOMPARATOR // WSAECOMPARATOR enumeration, used in the comparison.
}
