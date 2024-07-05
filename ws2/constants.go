package ws2

const (
	IOC_VOID     = 0x20000000
	IOC_OUT      = 0x40000000
	IOC_IN       = 0x80000000
	IOC_INOUT    = IOC_IN | IOC_OUT
	IOCPARM_MASK = 0x7f
)

const (
	COMP_EQUAL   WSAECOMPARATOR = 0
	COMP_NOTLESS WSAECOMPARATOR = 1
)

/*
  ws2def.h
*/

// Address families
const (
	AF_UNSPEC     = 0      // Unspecified
	AF_UNIX       = 1      // Local to host (pipes, portals)
	AF_INET       = 2      // Internetwork: UDP, TCP, etc.
	AF_IMPLINK    = 3      // Arpanet imp addresses
	AF_PUP        = 4      // Pup protocols: e.g. BSP
	AF_CHAOS      = 5      // Mit CHAOS protocols
	AF_NS         = 6      // XEROX NS protocols
	AF_IPX        = AF_NS  // IPX protocols: IPX, SPX, etc.
	AF_ISO        = 7      // ISO protocols
	AF_OSI        = AF_ISO // OSI is ISO
	AF_ECMA       = 8      // European computer manufacturers
	AF_DATAKIT    = 9      // Datakit protocols
	AF_CCITT      = 10     // CCITT protocols, X.25 etc
	AF_SNA        = 11     // IBM SNA
	AF_DECnet     = 12     // DECnet
	AF_DLI        = 13     // Direct data link interface
	AF_LAT        = 14     // LAT
	AF_HYLINK     = 15     // NSC Hyperchannel
	AF_APPLETALK  = 16     // AppleTalk
	AF_NETBIOS    = 17     // NetBios-style addresses
	AF_VOICEVIEW  = 18     // VoiceView
	AF_FIREFOX    = 19     // Protocols from Firefox
	AF_UNKNOWN1   = 20     // Somebody is using this!
	AF_BAN        = 21     // Banyan
	AF_ATM        = 22     // Native ATM Services
	AF_INET6      = 23     // Internetwork Version 6
	AF_CLUSTER    = 24     // Microsoft Wolfpack
	AF_12844      = 25     // IEEE 1284.4 WG AF
	AF_IRDA       = 26     // IrDA
	AF_NETDES     = 28     // Network Designers OSI & gateway
	AF_TCNPROCESS = 29     //
	AF_TCNMESSAGE = 30     //
	AF_ICLFXBM    = 31     //
	AF_BTH        = 32     // Bluetooth RFCOMM/L2CAP protocols
	AF_LINK       = 33     //
	AF_HYPERV     = 34     //
	AF_MAX        = 35     //
)

// Levels for socket I/O controls
const (
	SOL_SOCKET = 0xffff
	SOL_IP     = (SOL_SOCKET - 4)
	SOL_IPV6   = (SOL_SOCKET - 5)
)

// Socket options
const (
	SO_DEBUG               = 0x0001          // Turn on debugging info recording
	SO_ACCEPTCONN          = 0x0002          // Docket has had listen()
	SO_REUSEADDR           = 0x0004          // Sllow local address reuse
	SO_KEEPALIVE           = 0x0008          // Keep connections alive
	SO_DONTROUTE           = 0x0010          // Just use interface addresses
	SO_BROADCAST           = 0x0020          // Permit sending of broadcast msgs
	SO_USELOOPBACK         = 0x0040          // Bypass hardware when possible
	SO_LINGER              = 0x0080          // Linger on close if data present
	SO_OOBINLINE           = 0x0100          // Leave received OOB data in line
	SO_DONTLINGER          = ^(SO_LINGER)    //
	SO_EXCLUSIVEADDRUSE    = ^(SO_REUSEADDR) // Disallow local address reuse
	SO_SNDBUF              = 0x1001          // Send buffer size
	SO_RCVBUF              = 0x1002          // Receive buffer size
	SO_SNDLOWAT            = 0x1003          // Send low-water mark
	SO_RCVLOWAT            = 0x1004          // Receive low-water mark
	SO_SNDTIMEO            = 0x1005          // Send timeout
	SO_RCVTIMEO            = 0x1006          // Receive timeout
	SO_ERROR               = 0x1007          // Get error status and clear
	SO_TYPE                = 0x1008          // Get socket type
	SO_BSP_STATE           = 0x1009          // Get socket 5-tuple state
	SO_GROUP_ID            = 0x2001          // ID of a socket group
	SO_GROUP_PRIORITY      = 0x2002          // The relative priority within a group
	SO_MAX_MSG_SIZE        = 0x2003          // Maximum message size
	SO_CONDITIONAL_ACCEPT  = 0x3002          // Rnable true conditional accept: connection is not ack-ed to the other side until conditional function returns CF_ACCEPT
	SO_PAUSE_ACCEPT        = 0x3003          // Pause accepting new connections
	SO_COMPARTMENT_ID      = 0x3004          // Get/set the compartment for a socket
	SO_RANDOMIZE_PORT      = 0x3005          // Randomize assignment of wildcard ports
	SO_PORT_SCALABILITY    = 0x3006          // Enable port scalability
	SO_REUSE_UNICASTPORT   = 0x3007          // Defer ephemeral port allocation for outbound connections
	SO_REUSE_MULTICASTPORT = 0x3008          // Enable port reuse and disable unicast reception.
	SO_ORIGINAL_DST        = 0x300F          // Query the original destination address of a redirected connection.
	IP6T_SO_ORIGINAL_DST   = SO_ORIGINAL_DST //
)

// Base constant used for defining WSK-specific options.
const WSK_SO_BASE = 0x4000

// Options to use with [gs]etsockopt at the IPPROTO_TCP level.
const (
	TCP_NODELAY = 0x0001
)

// WinSock 2 extension -- manifest constants for WSAIoctl()
const (
	IOC_UNIX     = 0x00000000
	IOC_WS2      = 0x08000000
	IOC_PROTOCOL = 0x10000000
	IOC_VENDOR   = 0x18000000
)

// WSK-specific IO control codes are Winsock2 codes with the highest-order
// 3 bits of the Vendor/AddressFamily-specific field set to 1.
const (
	IOC_WSK = IOC_WS2 | 0x07000000
)

const (
	SIO_ASSOCIATE_HANDLE               = uint32(IOC_IN | IOC_WS2 | 1)
	SIO_ENABLE_CIRCULAR_QUEUEING       = uint32(IOC_VOID | IOC_WS2 | 2)
	SIO_FIND_ROUTE                     = uint32(IOC_OUT | IOC_WS2 | 3)
	SIO_FLUSH                          = uint32(IOC_VOID | IOC_WS2 | 4)
	SIO_GET_BROADCAST_ADDRESS          = uint32(IOC_OUT | IOC_WS2 | 5)
	SIO_GET_EXTENSION_FUNCTION_POINTER = uint32(IOC_INOUT | IOC_WS2 | 6)
	SIO_GET_QOS                        = uint32(IOC_INOUT | IOC_WS2 | 7)
	SIO_GET_GROUP_QOS                  = uint32(IOC_INOUT | IOC_WS2 | 8)
	SIO_MULTIPOINT_LOOPBACK            = uint32(IOC_IN | IOC_WS2 | 9)
	SIO_MULTICAST_SCOPE                = uint32(IOC_IN | IOC_WS2 | 10)
	SIO_SET_QOS                        = uint32(IOC_IN | IOC_WS2 | 11)
	SIO_SET_GROUP_QOS                  = uint32(IOC_IN | IOC_WS2 | 12)
	SIO_TRANSLATE_HANDLE               = uint32(IOC_INOUT | IOC_WS2 | 13)
	SIO_ROUTING_INTERFACE_QUERY        = uint32(IOC_INOUT | IOC_WS2 | 20)
	SIO_ROUTING_INTERFACE_CHANGE       = uint32(IOC_IN | IOC_WS2 | 21)
	SIO_ADDRESS_LIST_QUERY             = uint32(IOC_OUT | IOC_WS2 | 22)
	SIO_ADDRESS_LIST_CHANGE            = uint32(IOC_VOID | IOC_WS2 | 23)
	SIO_QUERY_TARGET_PNP_HANDLE        = uint32(IOC_OUT | IOC_WS2 | 24)
	SIO_NSP_NOTIFY_CHANGE              = uint32(IOC_IN | IOC_WS2 | 25)
	SIO_ADDRESS_LIST_SORT              = uint32(IOC_INOUT | IOC_WS2 | 25)
	SIO_QUERY_RSS_PROCESSOR_INFO       = uint32(IOC_INOUT | IOC_WS2 | 37)
	SIO_RESERVED_1                     = uint32(IOC_IN | IOC_WS2 | 26)
	SIO_RESERVED_2                     = uint32(IOC_IN | IOC_WS2 | 33)
)
const SIO_GET_MULTIPLE_EXTENSION_FUNCTION_POINTER uint32 = 0xC8000024

// N.B. required for backwards compatability to support 0 = IP for the
// level argument to get/setsockopt.
const IPPROTO_IP = 0

// Protocols.
const (
	IPPROTO_HOPOPTS  = 0   // IPv6 Hop-by-Hop options
	IPPROTO_ICMP     = 1   //
	IPPROTO_IGMP     = 2   //
	IPPROTO_GGP      = 3   //
	IPPROTO_IPV4     = 4   //
	IPPROTO_ST       = 5   //
	IPPROTO_TCP      = 6   //
	IPPROTO_CBT      = 7   //
	IPPROTO_EGP      = 8   //
	IPPROTO_IGP      = 9   //
	IPPROTO_PUP      = 12  //
	IPPROTO_UDP      = 17  //
	IPPROTO_IDP      = 22  //
	IPPROTO_RDP      = 27  //
	IPPROTO_IPV6     = 41  // IPv6 header
	IPPROTO_ROUTING  = 43  // IPv6 Routing header
	IPPROTO_FRAGMENT = 44  // IPv6 fragmentation header
	IPPROTO_ESP      = 50  // Encapsulating security payload
	IPPROTO_AH       = 51  // Authentication header
	IPPROTO_ICMPV6   = 58  // ICMPv6
	IPPROTO_NONE     = 59  // IPv6 no next header
	IPPROTO_DSTOPTS  = 60  // IPv6 Destination options
	IPPROTO_ND       = 77  //
	IPPROTO_ICLFXBM  = 78  //
	IPPROTO_PIM      = 103 //
	IPPROTO_PGM      = 113 //
	IPPROTO_L2TP     = 115 //
	IPPROTO_SCTP     = 132 //
	IPPROTO_RAW      = 255 //
	IPPROTO_MAX      = 256 //
	// These are reserved for internal use by Windows.
	IPPROTO_RESERVED_RAW          = 257
	IPPROTO_RESERVED_IPSEC        = 258
	IPPROTO_RESERVED_IPSECOFFLOAD = 259
	IPPROTO_RESERVED_WNV          = 260
	IPPROTO_RESERVED_MAX          = 261
)

// Port/socket numbers: network standard functions
const (
	IPPORT_TCPMUX     = 1
	IPPORT_ECHO       = 7
	IPPORT_DISCARD    = 9
	IPPORT_SYSTAT     = 11
	IPPORT_DAYTIME    = 13
	IPPORT_NETSTAT    = 15
	IPPORT_QOTD       = 17
	IPPORT_MSP        = 18
	IPPORT_CHARGEN    = 19
	IPPORT_FTP_DATA   = 20
	IPPORT_FTP        = 21
	IPPORT_TELNET     = 23
	IPPORT_SMTP       = 25
	IPPORT_TIMESERVER = 37
	IPPORT_NAMESERVER = 42
	IPPORT_WHOIS      = 43
	IPPORT_MTP        = 57
)

// Port/socket numbers: host specific functions
const (
	IPPORT_TFTP    = 69
	IPPORT_RJE     = 77
	IPPORT_FINGER  = 79
	IPPORT_TTYLINK = 87
	IPPORT_SUPDUP  = 95
)

// UNIX TCP sockets
const (
	IPPORT_POP3         = 110
	IPPORT_NTP          = 123
	IPPORT_EPMAP        = 135
	IPPORT_NETBIOS_NS   = 137
	IPPORT_NETBIOS_DGM  = 138
	IPPORT_NETBIOS_SSN  = 139
	IPPORT_IMAP         = 143
	IPPORT_SNMP         = 161
	IPPORT_SNMP_TRAP    = 162
	IPPORT_IMAP3        = 220
	IPPORT_LDAP         = 389
	IPPORT_HTTPS        = 443
	IPPORT_MICROSOFT_DS = 445
	IPPORT_EXECSERVER   = 512
	IPPORT_LOGINSERVER  = 513
	IPPORT_CMDSERVER    = 514
	IPPORT_EFSSERVER    = 520
)

// UNIX UDP sockets
const (
	IPPORT_BIFFUDP     = 512
	IPPORT_WHOSERVER   = 513
	IPPORT_ROUTESERVER = 520
)

// Ports < IPPORT_RESERVED are reserved for privileged processes (e.g. root).
const (
	IPPORT_RESERVED       = 1024
	IPPORT_REGISTERED_MIN = IPPORT_RESERVED
	IPPORT_REGISTERED_MAX = 0xbfff
	IPPORT_DYNAMIC_MIN    = 0xc000
	IPPORT_DYNAMIC_MAX    = 0xffff
)

// Common internet addresses
const (
	INADDR_ANY       = 0x00000000
	INADDR_LOOPBACK  = 0x7f000001
	INADDR_BROADCAST = 0xffffffff
	INADDR_NONE      = 0xffffffff
)

// Scope ID definition
//
// ScopeLevel enumeration
type ScopeLevel int

const (
	ScopeLevelInterface    ScopeLevel = 1
	ScopeLevelLink         ScopeLevel = 2
	ScopeLevelSubnet       ScopeLevel = 3
	ScopeLevelAdmin        ScopeLevel = 4
	ScopeLevelSite         ScopeLevel = 5
	ScopeLevelOrganization ScopeLevel = 8
	ScopeLevelGlobal       ScopeLevel = 14
	ScopeLevelCount        ScopeLevel = 16
)

// Definition for flags member of the WSAMsg structure
const (
	MSG_TRUNC    = 0x0100
	MSG_CTRUNC   = 0x0200
	MSG_BCAST    = 0x0400
	MSG_MCAST    = 0x0800
	MSG_ERRQUEUE = 0x1000
)

// Flags used in "hints" argument to GetAddrInfo()
const (
	AI_PASSIVE                  = 0x00000001 // Socket address will be used in bind() call
	AI_CANONNAME                = 0x00000002 // Return canonical name in first ai_canonname
	AI_NUMERICHOST              = 0x00000004 // Nodename must be a numeric address string
	AI_NUMERICSERV              = 0x00000008 // Servicename must be a numeric port number
	AI_DNS_ONLY                 = 0x00000010 // Restrict queries to unicast DNS only (no LLMNR, netbios, etc.)
	AI_FORCE_CLEAR_TEXT         = 0x00000020 // Force clear text DNS query
	AI_BYPASS_DNS_CACHE         = 0x00000040 // Bypass DNS cache
	AI_RETURN_TTL               = 0x00000080 // Return record TTL
	AI_ALL                      = 0x00000100 // Query both IP6 and IP4 with AI_V4MAPPED
	AI_ADDRCONFIG               = 0x00000400 // Resolution only if global address configured
	AI_V4MAPPED                 = 0x00000800 // On v6 failure, query v4 and convert to V4MAPPED format
	AI_NON_AUTHORITATIVE        = 0x00004000 // LUP_NON_AUTHORITATIVE
	AI_SECURE                   = 0x00008000 // LUP_SECURE
	AI_RETURN_PREFERRED_NAMES   = 0x00010000 // LUP_RETURN_PREFERRED_NAMES
	AI_FQDN                     = 0x00020000 // Return the FQDN in ai_canonname
	AI_FILESERVER               = 0x00040000 // Resolving fileserver name resolution
	AI_DISABLE_IDN_ENCODING     = 0x00080000 // Disable Internationalized Domain Names handling
	AI_SECURE_WITH_FALLBACK     = 0x00100000 // Forces clear text fallback if the secure DNS query fails
	AI_EXCLUSIVE_CUSTOM_SERVERS = 0x00200000 // Use exclusively the custom DNS servers
	AI_RETURN_RESPONSE_FLAGS    = 0x10000000 // Requests extra information about the DNS results
	AI_REQUIRE_SECURE           = 0x20000000 // Forces the DNS query to be done over secure protocols
	AI_RESOLUTION_HANDLE        = 0x40000000 // Request resolution handle
	AI_EXTENDED                 = 0x80000000 // Indicates this is extended AddrInfoEx(2/..) struct
)

// AddrInfoEx versions
const (
	ADDRINFOEX_VERSION_2 = 2
	ADDRINFOEX_VERSION_3 = 3
	ADDRINFOEX_VERSION_4 = 4
	ADDRINFOEX_VERSION_5 = 5
	ADDRINFOEX_VERSION_6 = 6
)

// Types of custom DNS servers specified in the ai_servers parameter.
const (
	AI_DNS_SERVER_TYPE_UDP = 0x1
	AI_DNS_SERVER_TYPE_DOH = 0x2
)

// Flags for custom servers.
const (
	AI_DNS_SERVER_UDP_FALLBACK = 0x1
)

// Flags returned through ai_returnflags, when AI_RETURN_RESPONSE_FLAGS is set.
const (
	AI_DNS_RESPONSE_SECURE   = 0x1 // Present if the resolution was done through secure protocols.
	AI_DNS_RESPONSE_HOSTFILE = 0x2
)

// Name Spaces
const (
	NS_ALL         = 0  //
	NS_SAP         = 1  //
	NS_NDS         = 2  //
	NS_PEER_BROWSE = 3  //
	NS_SLP         = 5  //
	NS_DHCP        = 6  //
	NS_TCPIP_LOCAL = 10 //
	NS_TCPIP_HOSTS = 11 //
	NS_DNS         = 12 //
	NS_NETBT       = 13 //
	NS_WINS        = 14 //
	NS_NLA         = 15 // Network Location Awareness
	NS_BTH         = 16 // Bluetooth SDP Namespace
	NS_NBP         = 20 //
	NS_MS          = 30 //
	NS_STDA        = 31 //
	NS_NTDS        = 32 //
	NS_EMAIL       = 37 //
	NS_PNRPNAME    = 38 //
	NS_PNRPCLOUD   = 39 //
	NS_X500        = 40 //
	NS_NIS         = 41 //
	NS_NISPLUS     = 42 //
	NS_WRQ         = 50 //
	NS_NETDES      = 60 // Network Designers Limited
)

// Flags for getnameinfo()
const (
	NI_NOFQDN      = 0x01 // Only return nodename portion for local hosts
	NI_NUMERICHOST = 0x02 // Return numeric form of the host's address
	NI_NAMEREQD    = 0x04 // Error if the host's name not in DNS
	NI_NUMERICSERV = 0x08 // Return numeric form of the service (port #)
	NI_DGRAM       = 0x10 // Service is a datagram service
	NI_MAXHOST     = 1025 // Max size of a fully-qualified domain name
	NI_MAXSERV     = 32   // Max size of a service name
)

/*
  ws2ipdef.h
*/

// Interface flags
const (
	IFF_UP           = 0x00000001 // Interface is up.
	IFF_BROADCAST    = 0x00000002 // Broadcast is supported.
	IFF_LOOPBACK     = 0x00000004 // This is a loopback interface.
	IFF_POINTTOPOINT = 0x00000008 // This is a point-to-point interface.
	IFF_MULTICAST    = 0x00000010 // Multicast is supported.
)

// Options to use with [gs]etsockopt at the IPPROTO_IP level.
// The values should be consistent with the IPv6 equivalents.
//
// Options to use with setsockopt and getsockopt at the IPPROTO_IP level.
const (
	IP_OPTIONS                     = 1  // Set/get IP options.
	IP_HDRINCL                     = 2  // Header is included with data.
	IP_TOS                         = 3  // IP type of service.
	IP_TTL                         = 4  // IP TTL (hop limit).
	IP_MULTICAST_IF                = 9  // IP multicast interface.
	IP_MULTICAST_TTL               = 10 // IP multicast TTL (hop limit).
	IP_MULTICAST_LOOP              = 11 // IP multicast loopback.
	IP_ADD_MEMBERSHIP              = 12 // Add an IP group membership.
	IP_DROP_MEMBERSHIP             = 13 // Drop an IP group membership.
	IP_DONTFRAGMENT                = 14 // Don't fragment IP datagrams.
	IP_ADD_SOURCE_MEMBERSHIP       = 15 // Join IP group/source.
	IP_DROP_SOURCE_MEMBERSHIP      = 16 // Leave IP group/source.
	IP_BLOCK_SOURCE                = 17 // Block IP group/source.
	IP_UNBLOCK_SOURCE              = 18 // Unblock IP group/source.
	IP_PKTINFO                     = 19 // Receive packet information.
	IP_HOPLIMIT                    = 21 // Receive packet hop limit.
	IP_RECVTTL                     = 21 // Receive packet Time To Live (TTL).
	IP_RECEIVE_BROADCAST           = 22 // Allow/block broadcast reception.
	IP_RECVIF                      = 24 // Receive arrival interface.
	IP_RECVDSTADDR                 = 25 // Receive destination address.
	IP_IFLIST                      = 28 // Enable/Disable an interface list.
	IP_ADD_IFLIST                  = 29 // Add an interface list entry.
	IP_DEL_IFLIST                  = 30 // Delete an interface list entry.
	IP_UNICAST_IF                  = 31 // IP unicast interface.
	IP_RTHDR                       = 32 // Set/get IPv6 routing header.
	IP_GET_IFLIST                  = 33 // Get an interface list.
	IP_RECVRTHDR                   = 38 // Receive the routing header.
	IP_TCLASS                      = 39 // Packet traffic class.
	IP_RECVTCLASS                  = 40 // Receive packet traffic class.
	IP_RECVTOS                     = 40 // Receive packet Type Of Service (TOS).
	IP_ORIGINAL_ARRIVAL_IF         = 47 // Original Arrival Interface Index.
	IP_ECN                         = 50 // IP ECN codepoint.
	IP_RECVECN                     = 50 // Receive ECN codepoints in the IP header.
	IP_PKTINFO_EX                  = 51 // Receive extended packet information.
	IP_WFP_REDIRECT_RECORDS        = 60 // WFP's Connection Redirect Records.
	IP_WFP_REDIRECT_CONTEXT        = 70 // WFP's Connection Redirect Context.
	IP_MTU_DISCOVER                = 71 // Set/get path MTU discover state.
	IP_MTU                         = 73 // Get path MTU.
	IP_NRT_INTERFACE               = 74 // Set NRT interface constraint (outbound).
	IP_RECVERR                     = 75 // Receive ICMP errors.
	IP_USER_MTU                    = 76 // Set/get app defined upper bound IP layer MTU.
	IP_UNSPECIFIED_TYPE_OF_SERVICE = -1
	IP_UNSPECIFIED_USER_MTU        = 1<<32 - 1
)

// N.B. These addresses are in network byte order.
var (
	IPV6_ADDRESS_BITS = 128

	IN6ADDR_ANY_INIT = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	IN6ADDR_LOOPBACK_INIT = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	IN6ADDR_ALLNODESONNODE_INIT = [16]byte{
		0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	IN6ADDR_ALLNODESONLINK_INIT = [16]byte{
		0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
	}

	IN6ADDR_ALLROUTERSONLINK_INIT = [16]byte{
		0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
	}

	IN6ADDR_ALLMLDV2ROUTERSONLINK_INIT = [16]byte{
		0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x16,
	}

	IN6ADDR_TEREDOINITIALLINKLOCALADDRESS_INIT = [16]byte{
		0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe,
	}

	IN6ADDR_TEREDOOLDLINKLOCALADDRESSXP_INIT = [16]byte{
		0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 'T', 'E', 'R', 'E', 'D', 'O',
	}

	IN6ADDR_TEREDOOLDLINKLOCALADDRESSVISTA_INIT = [16]byte{
		0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}

	IN6ADDR_LINKLOCALPREFIX_INIT = [2]byte{0xfe, 0x80}

	IN6ADDR_MULTICASTPREFIX_INIT = [2]byte{0xff, 0x00}

	IN6ADDR_SOLICITEDNODEMULTICASTPREFIX_INIT = [13]byte{
		0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0xff,
	}

	IN6ADDR_V4MAPPEDPREFIX_INIT = [12]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
	}

	IN6ADDR_6TO4PREFIX_INIT = [2]byte{0x20, 0x02}

	IN6ADDR_TEREDOPREFIX_INIT = [4]byte{0x20, 0x01, 0x00, 0x00}

	IN6ADDR_TEREDOPREFIX_INIT_OLD = [4]byte{0x3f, 0xfe, 0x83, 0x1f}

	IN6ADDR_ULAPREFIX_INIT = [1]byte{0xfc}

	IN6ADDR_SITELOCALPREFIX_INIT = [2]byte{0xfe, 0xc0}

	IN6ADDR_6BONETESTPREFIX_INIT = [2]byte{0x3f, 0xfe}
)

// Prefix lengths
const (
	IN6ADDR_LINKLOCALPREFIX_LENGTH              = 64
	IN6ADDR_MULTICASTPREFIX_LENGTH              = 8
	IN6ADDR_SOLICITEDNODEMULTICASTPREFIX_LENGTH = 104
	IN6ADDR_V4MAPPEDPREFIX_LENGTH               = 96
	IN6ADDR_6TO4PREFIX_LENGTH                   = 16
	IN6ADDR_TEREDOPREFIX_LENGTH                 = 32
)

// TCP/IP specific Ioctl codes.
const (
	SIO_GET_INTERFACE_LIST    = 0x80000007 | 0x20000000 | (0x74 << 16) | (127 << 8) | 4
	SIO_GET_INTERFACE_LIST_EX = 0x80000007 | 0x20000000 | (0x74 << 16) | (126 << 8) | 4
	SIO_SET_MULTICAST_FILTER  = 0x80000008 | 0x40000000 | (0x74 << 16) | (125 << 8) | 4
	SIO_GET_MULTICAST_FILTER  = 0x80000008 | 0x40000000 | (0x74 << 16) | (124 << 8) | 4
	SIOCSIPMSFILTER           = SIO_SET_MULTICAST_FILTER
	SIOCGIPMSFILTER           = SIO_GET_MULTICAST_FILTER
)

// Protocol independent ioctls for setting and retrieving multicast filters.
const (
	SIOCSMSFILTER = 0x80000008 | 0x40000000 | (0x74 << 16) | (126 << 8) | 4
	SIOCGMSFILTER = 0x80000008 | 0x40000000 | (0x74 << 16) | (127 << 8) | 4
)

// Query and change notification ioctls for the ideal send backlog size
// for a given connection. Clients should use the wrappers defined in
// ws2tcpip.h rather than using these ioctls directly.
const (
	SIO_IDEAL_SEND_BACKLOG_QUERY  = 0x80000007 | 0x20000000 | (0x74 << 16) | (123 << 8) | 4
	SIO_IDEAL_SEND_BACKLOG_CHANGE = 0x80000007 | 0x20000000 | (0x74 << 16) | (122 << 8)
)

// Protocol independent multicast source filter options.
const (
	MCAST_JOIN_GROUP         = 41 // Join all sources for a group
	MCAST_LEAVE_GROUP        = 42 // Drop all sources for a group
	MCAST_BLOCK_SOURCE       = 43 // Block IP group/source
	MCAST_UNBLOCK_SOURCE     = 44 // Unblock IP group/source
	MCAST_JOIN_SOURCE_GROUP  = 45 // Join IP group/source
	MCAST_LEAVE_SOURCE_GROUP = 46 // Leave IP group/source
)

// Options to use with [gs]etsockopt at the IPPROTO_IPV6 level.
// These are specified in RFCs 3493 and 3542.
// The values should be consistent with the IPv6 equivalents.
const (
	IPV6_HOPOPTS              = 1                    // Set/get IPv6 hop-by-hop options.
	IPV6_HDRINCL              = 2                    // Header is included with data.
	IPV6_UNICAST_HOPS         = 4                    // IP unicast hop limit.
	IPV6_MULTICAST_IF         = 9                    // IP multicast interface.
	IPV6_MULTICAST_HOPS       = 10                   // IP multicast hop limit.
	IPV6_MULTICAST_LOOP       = 11                   // IP multicast loopback.
	IPV6_ADD_MEMBERSHIP       = 12                   // Add an IP group membership.
	IPV6_JOIN_GROUP           = IPV6_ADD_MEMBERSHIP  //
	IPV6_DROP_MEMBERSHIP      = 13                   // Drop an IP group membership.
	IPV6_LEAVE_GROUP          = IPV6_DROP_MEMBERSHIP //
	IPV6_DONTFRAG             = 14                   // Don't fragment IP datagrams.
	IPV6_PKTINFO              = 19                   // Receive packet information.
	IPV6_HOPLIMIT             = 21                   // Receive packet hop limit.
	IPV6_PROTECTION_LEVEL     = 23                   // Set/get IPv6 protection level.
	IPV6_RECVIF               = 24                   // Receive arrival interface.
	IPV6_RECVDSTADDR          = 25                   // Receive destination address.
	IPV6_CHECKSUM             = 26                   // Offset to checksum for raw IP socket send.
	IPV6_V6ONLY               = 27                   // Treat wildcard bind as AF_INET6-only.
	IPV6_IFLIST               = 28                   // Enable/Disable an interface list.
	IPV6_ADD_IFLIST           = 29                   // Add an interface list entry.
	IPV6_DEL_IFLIST           = 30                   // Delete an interface list entry.
	IPV6_UNICAST_IF           = 31                   // IP unicast interface.
	IPV6_RTHDR                = 32                   // Set/get IPv6 routing header.
	IPV6_GET_IFLIST           = 33                   // Get an interface list.
	IPV6_RECVRTHDR            = 38                   // Receive the routing header.
	IPV6_TCLASS               = 39                   // Packet traffic class.
	IPV6_RECVTCLASS           = 40                   // Receive packet traffic class.
	IPV6_ECN                  = 50                   // IPv6 ECN codepoint.
	IPV6_RECVECN              = 50                   // Receive ECN codepoints in the IPv6 header.
	IPV6_PKTINFO_EX           = 51                   // Receive extended packet information.
	IPV6_WFP_REDIRECT_RECORDS = 60                   // WFP's Connection Redirect Records
	IPV6_WFP_REDIRECT_CONTEXT = 70                   // WFP's Connection Redirect Context
	IPV6_MTU_DISCOVER         = 71                   // Set/get path MTU discover state.
	IPV6_MTU                  = 72                   // Get path MTU.
	IPV6_NRT_INTERFACE        = 74                   // Set NRT interface constraint (outbound).
	IPV6_RECVERR              = 75                   // Receive ICMPv6 errors.
	IPV6_USER_MTU             = 76                   // Set/get app defined upper bound IP layer MTU.
)

// IP_UNSPECIFIED_HOP_LIMIT represents an unspecified hop limit
const IP_UNSPECIFIED_HOP_LIMIT = -1

// IP_PROTECTION_LEVEL is an alias for IPV6_PROTECTION_LEVEL
const IP_PROTECTION_LEVEL = IPV6_PROTECTION_LEVEL

// Values of IPV6_PROTECTION_LEVEL.
const (
	PROTECTION_LEVEL_UNRESTRICTED   = 10       // For peer-to-peer apps.
	PROTECTION_LEVEL_EDGERESTRICTED = 20       // Same as unrestricted. Except for Teredo.
	PROTECTION_LEVEL_RESTRICTED     = 30       // For Intranet apps.
	PROTECTION_LEVEL_DEFAULT        = ^uint(0) // Use UINT_MAX as the default value.
)

// Maximum length of address literals (potentially including a port number)
// generated by any address-to-string conversion routine.  This length can
// be used when declaring buffers used with getnameinfo, WSAAddressToString,
// inet_ntoa, etc.  We just provide one define, rather than one per api,
// to avoid confusion.
//
// The totals are derived from the following data:
//
//	15: IPv4 address
//	45: IPv6 address including embedded IPv4 address
//	11: Scope Id
//	 2: Brackets around IPv6 address when port is present
//	 6: Port (including colon)
//	 1: Terminating null byte
const (
	INET_ADDRSTRLEN  = 22
	INET6_ADDRSTRLEN = 65
)

//
// Options to use with [gs]etsockopt at the IPPROTO_TCP level.
// TCP_NODELAY is defined in ws2def.h for historical reasons.
//

// Offload preferences supported.
const (
	TCP_OFFLOAD_NO_PREFERENCE = 0
	TCP_OFFLOAD_NOT_PREFERRED = 1
	TCP_OFFLOAD_PREFERRED     = 2
)

const (
	TCP_EXPEDITED_1122             = 0x0002
	TCP_KEEPALIVE                  = 3
	TCP_MAXSEG                     = 4
	TCP_MAXRT                      = 5
	TCP_STDURG                     = 6
	TCP_NOURG                      = 7
	TCP_ATMARK                     = 8
	TCP_NOSYNRETRIES               = 9
	TCP_TIMESTAMPS                 = 10
	TCP_OFFLOAD_PREFERENCE         = 11
	TCP_CONGESTION_ALGORITHM       = 12
	TCP_DELAY_FIN_ACK              = 13
	TCP_MAXRTMS                    = 14
	TCP_FASTOPEN                   = 15
	TCP_KEEPCNT                    = 16
	TCP_KEEPIDLE                   = TCP_KEEPALIVE
	TCP_KEEPINTVL                  = 17
	TCP_FAIL_CONNECT_ON_ICMP_ERROR = 18
	TCP_ICMP_ERROR_INFO            = 19
)

//
// Options to use with [gs]etsockopt at the IPPROTO_UDP level.
// UDP_NOCHECKSUM is defined in ws2tcpip.h for historical reasons.
// UDP_CHECKSUM_COVERAGE is defined in ws2tcpip.h for historical reasons.
//

const (
	UDP_SEND_MSG_SIZE           = 2
	UDP_RECV_MAX_COALESCED_SIZE = 3
)

// Control message types at the IPPROTO_UDP level.
const (
	UDP_COALESCED_INFO = 3
)

/*
  Winsock2.h
*/

const WINSOCK_VERSION = 0x0202

const FD_SETSIZE = 64

const (
	FIONREAD = 0x4004667f // Get # bytes to read
	FIONBIO  = 0x8004667e // Set/clear non-blocking i/o
	FIOASYNC = 0x8004667d // Set/clear async i/o
)

// Constants for socket I/O controls
const (
	SIOCSHIWAT = 0x80730000                                // Set high watermark
	SIOCGHIWAT = (0x40000000 | (4 << 16) | ('s' << 8) | 1) // Get high watermark
	SIOCSLOWAT = (0x80000000 | (4 << 16) | ('s' << 8) | 2) // Set low watermark
	SIOCGLOWAT = (0x40000000 | (4 << 16) | ('s' << 8) | 3) // Get low watermark
	SIOCATMARK = (0x40000000 | (4 << 16) | ('s' << 8) | 7) // At oob mark?
)

//
// Constants and structures defined by the internet system,
// Per RFC 790, September 1981, taken from the BSD file netinet/in.h.
// IPv6 additions per RFC 2292.
//

// Link numbers
const (
	IMPLINK_IP        = 155
	IMPLINK_LOWEXPER  = 156
	IMPLINK_HIGHEXPER = 158
)

const (
	WSADESCRIPTION_LEN = 256
	WSASYS_STATUS_LEN  = 128
)

// This is used instead of -1, since the SOCKET type is unsigned.
const INVALID_SOCKET = ^SOCKET(0)
const SOCKET_ERROR = -1

// The following may be used in place of the address family, socket type, or
// protocol in a call to WSASocket to indicate that the corresponding value
// should be taken from the supplied WSAPROTOCOL_INFO structure instead of the
// parameter itself.
const FROM_PROTOCOL_INFO = -1

// Types
const (
	SOCK_STREAM    = 1 // Stream socket
	SOCK_DGRAM     = 2 // Datagram socket
	SOCK_RAW       = 3 // Raw-protocol interface
	SOCK_RDM       = 4 // Reliably-delivered message
	SOCK_SEQPACKET = 5 // Sequenced packet stream
)

// WinSock 2 extension -- new options
const (
	SO_PROTOCOL_INFOA = 0x2004 // WSAProtocolInfoA structure
	SO_PROTOCOL_INFOW = 0x2005 // WSAProtocolInfoW structure
	PVD_CONFIG        = 0x3001 // Configuration info for service provider connection is not ack-ed to the other side until conditional function returns CF_ACCEPT
)

// Protocol families, same as address families for now.
const (
	PF_UNSPEC    = AF_UNSPEC
	PF_UNIX      = AF_UNIX
	PF_INET      = AF_INET
	PF_IMPLINK   = AF_IMPLINK
	PF_PUP       = AF_PUP
	PF_CHAOS     = AF_CHAOS
	PF_NS        = AF_NS
	PF_IPX       = AF_IPX
	PF_ISO       = AF_ISO
	PF_OSI       = AF_OSI
	PF_ECMA      = AF_ECMA
	PF_DATAKIT   = AF_DATAKIT
	PF_CCITT     = AF_CCITT
	PF_SNA       = AF_SNA
	PF_DECnet    = AF_DECnet
	PF_DLI       = AF_DLI
	PF_LAT       = AF_LAT
	PF_HYLINK    = AF_HYLINK
	PF_APPLETALK = AF_APPLETALK
	PF_VOICEVIEW = AF_VOICEVIEW
	PF_FIREFOX   = AF_FIREFOX
	PF_UNKNOWN1  = AF_UNKNOWN1
	PF_BAN       = AF_BAN
	PF_ATM       = AF_ATM
	PF_INET6     = AF_INET6
	PF_BTH       = AF_BTH
	PF_MAX       = AF_MAX
)

// Maximum queue length specifiable by listen.
const SOMAXCONN = 0x7fffffff

const (
	MSG_OOB            = 0x1    // Process out-of-band data
	MSG_PEEK           = 0x2    // Peek at incoming message
	MSG_DONTROUTE      = 0x4    // Send without using routing tables
	MSG_WAITALL        = 0x8    // Do not complete until packet is completely filled
	MSG_PUSH_IMMEDIATE = 0x20   // Do not delay receive request completion if data is available
	MSG_PARTIAL        = 0x8000 // Partial send or recv for message xport
)

// WinSock 2 extension -- new flags for WSASend(), WSASendTo(), WSARecv() and WSARecvFrom()
const (
	MSG_INTERRUPT = 0x10 // Send/Recv in the interrupt context
	MSG_MAXIOVLEN = 16
)

// Define constant based on rfc883, used by GetHostByxxxx() calls.
const MAXGETHOSTSTRUCT = 1024

// WinSock 2 extension -- bit values and indices for FD_XXX network events
const (
	FD_READ_BIT                     = 0
	FD_READ                         = 1 << FD_READ_BIT
	FD_WRITE_BIT                    = 1
	FD_WRITE                        = 1 << FD_WRITE_BIT
	FD_OOB_BIT                      = 2
	FD_OOB                          = 1 << FD_OOB_BIT
	FD_ACCEPT_BIT                   = 3
	FD_ACCEPT                       = 1 << FD_ACCEPT_BIT
	FD_CONNECT_BIT                  = 4
	FD_CONNECT                      = 1 << FD_CONNECT_BIT
	FD_CLOSE_BIT                    = 5
	FD_CLOSE                        = 1 << FD_CLOSE_BIT
	FD_QOS_BIT                      = 6
	FD_QOS                          = 1 << FD_QOS_BIT
	FD_GROUP_QOS_BIT                = 7
	FD_GROUP_QOS                    = 1 << FD_GROUP_QOS_BIT
	FD_ROUTING_INTERFACE_CHANGE_BIT = 8
	FD_ROUTING_INTERFACE_CHANGE     = 1 << FD_ROUTING_INTERFACE_CHANGE_BIT
	FD_ADDRESS_LIST_CHANGE_BIT      = 9
	FD_ADDRESS_LIST_CHANGE          = 1 << FD_ADDRESS_LIST_CHANGE_BIT
	FD_MAX_EVENTS                   = 10
	FD_ALL_EVENTS                   = (1 << FD_MAX_EVENTS) - 1
)

// All Windows Sockets error constants are biased by WSABASEERR from
// the "normal"
const WSABASEERR = 10000

// Windows Sockets definitions of regular Microsoft C error constants
const (
	WSAEINTR  = WSABASEERR + 4
	WSAEBADF  = WSABASEERR + 9
	WSAEACCES = WSABASEERR + 13
	WSAEFAULT = WSABASEERR + 14
	WSAEINVAL = WSABASEERR + 22
	WSAEMFILE = WSABASEERR + 24
)

// Windows Sockets definitions of regular Berkeley error constants
const (
	WSAEWOULDBLOCK     = WSABASEERR + 35
	WSAEINPROGRESS     = WSABASEERR + 36
	WSAEALREADY        = WSABASEERR + 37
	WSAENOTSOCK        = WSABASEERR + 38
	WSAEDESTADDRREQ    = WSABASEERR + 39
	WSAEMSGSIZE        = WSABASEERR + 40
	WSAEPROTOTYPE      = WSABASEERR + 41
	WSAENOPROTOOPT     = WSABASEERR + 42
	WSAEPROTONOSUPPORT = WSABASEERR + 43
	WSAESOCKTNOSUPPORT = WSABASEERR + 44
	WSAEOPNOTSUPP      = WSABASEERR + 45
	WSAEPFNOSUPPORT    = WSABASEERR + 46
	WSAEAFNOSUPPORT    = WSABASEERR + 47
	WSAEADDRINUSE      = WSABASEERR + 48
	WSAEADDRNOTAVAIL   = WSABASEERR + 49
	WSAENETDOWN        = WSABASEERR + 50
	WSAENETUNREACH     = WSABASEERR + 51
	WSAENETRESET       = WSABASEERR + 52
	WSAECONNABORTED    = WSABASEERR + 53
	WSAECONNRESET      = WSABASEERR + 54
	WSAENOBUFS         = WSABASEERR + 55
	WSAEISCONN         = WSABASEERR + 56
	WSAENOTCONN        = WSABASEERR + 57
	WSAESHUTDOWN       = WSABASEERR + 58
	WSAETOOMANYREFS    = WSABASEERR + 59
	WSAETIMEDOUT       = WSABASEERR + 60
	WSAECONNREFUSED    = WSABASEERR + 61
	WSAELOOP           = WSABASEERR + 62
	WSAENAMETOOLONG    = WSABASEERR + 63
	WSAEHOSTDOWN       = WSABASEERR + 64
	WSAEHOSTUNREACH    = WSABASEERR + 65
	WSAENOTEMPTY       = WSABASEERR + 66
	WSAEPROCLIM        = WSABASEERR + 67
	WSAEUSERS          = WSABASEERR + 68
	WSAEDQUOT          = WSABASEERR + 69
	WSAESTALE          = WSABASEERR + 70
	WSAEREMOTE         = WSABASEERR + 71
)

// Extended Windows Sockets error constant definitions
const (
	WSASYSNOTREADY         = WSABASEERR + 91
	WSAVERNOTSUPPORTED     = WSABASEERR + 92
	WSANOTINITIALISED      = WSABASEERR + 93
	WSAEDISCON             = WSABASEERR + 101
	WSAENOMORE             = WSABASEERR + 102
	WSAECANCELLED          = WSABASEERR + 103
	WSAEINVALIDPROCTABLE   = WSABASEERR + 104
	WSAEINVALIDPROVIDER    = WSABASEERR + 105
	WSAEPROVIDERFAILEDINIT = WSABASEERR + 106
	WSASYSCALLFAILURE      = WSABASEERR + 107
	WSASERVICE_NOT_FOUND   = WSABASEERR + 108
	WSATYPE_NOT_FOUND      = WSABASEERR + 109
	WSA_E_NO_MORE          = WSABASEERR + 110
	WSA_E_CANCELLED        = WSABASEERR + 111
	WSAEREFUSED            = WSABASEERR + 112
)

//
// Error return codes from GetHostByName() and GetHostByAddr()
// (when using the resolver). Note that these errors are
// retrieved via WSAGetLastError() and must therefore follow
// the rules for avoiding clashes with error numbers from
// specific implementations or language run-time systems.
// For this reason the codes are based at WSABASEERR+1001.
// Note also that [WSA]NO_ADDRESS is defined only for
// compatibility purposes.
//

// Authoritative Answer: Host not found
const WSAHOST_NOT_FOUND = WSABASEERR + 1001

// Non-Authoritative: Host not found, or SERVERFAIL
const WSATRY_AGAIN = WSABASEERR + 1002

// Non-recoverable errors, FORMERR, REFUSED, NOTIMP
const WSANO_RECOVERY = WSABASEERR + 1003

// Valid name, no data record of requested type
const WSANO_DATA = WSABASEERR + 1004

// Define QOS related error return codes
const (
	WSA_QOS_RECEIVERS          = WSABASEERR + 1005 // At least one Reserve has arrived
	WSA_QOS_SENDERS            = WSABASEERR + 1006 // At least one Path has arrived
	WSA_QOS_NO_SENDERS         = WSABASEERR + 1007 // There are no senders
	WSA_QOS_NO_RECEIVERS       = WSABASEERR + 1008 // There are no receivers
	WSA_QOS_REQUEST_CONFIRMED  = WSABASEERR + 1009 // Reserve has been confirmed
	WSA_QOS_ADMISSION_FAILURE  = WSABASEERR + 1010 // Error due to lack of resources
	WSA_QOS_POLICY_FAILURE     = WSABASEERR + 1011 // Rejected for administrative reasons - bad credentials
	WSA_QOS_BAD_STYLE          = WSABASEERR + 1012 // Unknown or conflicting style
	WSA_QOS_BAD_OBJECT         = WSABASEERR + 1013 // Problem with some part of the filterspec or providerspecific buffer in general
	WSA_QOS_TRAFFIC_CTRL_ERROR = WSABASEERR + 1014 // Problem with some part of the flowspec
	WSA_QOS_GENERIC_ERROR      = WSABASEERR + 1015 // General error
	WSA_QOS_ESERVICETYPE       = WSABASEERR + 1016 // Invalid service type in flowspec
	WSA_QOS_EFLOWSPEC          = WSABASEERR + 1017 // Invalid flowspec
	WSA_QOS_EPROVSPECBUF       = WSABASEERR + 1018 // Invalid provider specific buffer
	WSA_QOS_EFILTERSTYLE       = WSABASEERR + 1019 // Invalid filter style
	WSA_QOS_EFILTERTYPE        = WSABASEERR + 1020 // Invalid filter type
	WSA_QOS_EFILTERCOUNT       = WSABASEERR + 1021 // Incorrect number of filters
	WSA_QOS_EOBJLENGTH         = WSABASEERR + 1022 // Invalid object length
	WSA_QOS_EFLOWCOUNT         = WSABASEERR + 1023 // Incorrect number of flows
	WSA_QOS_EUNKOWNPSOBJ       = WSABASEERR + 1024 // Unknown object in provider specific buffer
	WSA_QOS_EPOLICYOBJ         = WSABASEERR + 1025 // Invalid policy object in provider specific buffer
	WSA_QOS_EFLOWDESC          = WSABASEERR + 1026 // Invalid flow descriptor in the list
	WSA_QOS_EPSFLOWSPEC        = WSABASEERR + 1027 // Inconsistent flow spec in provider specific buffer
	WSA_QOS_EPSFILTERSPEC      = WSABASEERR + 1028 // Invalid filter spec in provider specific buffer
	WSA_QOS_ESDMODEOBJ         = WSABASEERR + 1029 // Invalid shape discard mode object in provider specific buffer
	WSA_QOS_ESHAPERATEOBJ      = WSABASEERR + 1030 // Invalid shaping rate object in provider specific buffer
	WSA_QOS_RESERVED_PETYPE    = WSABASEERR + 1031 // Reserved policy element in provider specific buffer
)

const (
	WSA_IO_PENDING        = 997
	WSA_IO_INCOMPLETE     = 996
	WSA_INVALID_HANDLE    = 6
	WSA_INVALID_PARAMETER = 87
	WSA_NOT_ENOUGH_MEMORY = 8
	WSA_OPERATION_ABORTED = 995
)

const (
	WSA_INVALID_EVENT       = WSAEVENT(0)
	WSA_MAXIMUM_WAIT_EVENTS = 64
	WSA_WAIT_FAILED         = uint32(0xFFFFFFFF)
	WSA_WAIT_EVENT_0        = 0x00000000
	WSA_WAIT_IO_COMPLETION  = 0x00000100
	WSA_WAIT_TIMEOUT        = 0x00000102
	WSA_INFINITE            = 0xFFFFFFFF
)

// WinSock 2 extension -- manifest constants for return values of the condition function
const (
	CF_ACCEPT = 0x0000
	CF_REJECT = 0x0001
	CF_DEFER  = 0x0002
)

// WinSock 2 extension -- manifest constants for Shutdown()
const (
	SD_RECEIVE = 0x00
	SD_SEND    = 0x01
	SD_BOTH    = 0x02
)

// WinSock 2 extension -- data type and manifest constants for socket groups
const (
	SG_UNCONSTRAINED_GROUP = 0x01
	SG_CONSTRAINED_GROUP   = 0x02
)

// WinSock 2 extension -- WSAProtocolInfo structure and associated
// manifest constants
const (
	MAX_PROTOCOL_CHAIN = 7
)

const (
	BASE_PROTOCOL    = 1
	LAYERED_PROTOCOL = 0
	WSAPROTOCOL_LEN  = 255
)

// Flag bit definitions for ProviderFlags
const (
	PFL_MULTIPLE_PROTO_ENTRIES  = 0x00000001
	PFL_RECOMMENDED_PROTO_ENTRY = 0x00000002
	PFL_HIDDEN                  = 0x00000004
	PFL_MATCHES_PROTOCOL_ZERO   = 0x00000008
	PFL_NETWORKDIRECT_PROVIDER  = 0x00000010
)

// Flag bit definitions for ServiceFlags1
const (
	XP1_CONNECTIONLESS           = 0x00000001
	XP1_GUARANTEED_DELIVERY      = 0x00000002
	XP1_GUARANTEED_ORDER         = 0x00000004
	XP1_MESSAGE_ORIENTED         = 0x00000008
	XP1_PSEUDO_STREAM            = 0x00000010
	XP1_GRACEFUL_CLOSE           = 0x00000020
	XP1_EXPEDITED_DATA           = 0x00000040
	XP1_CONNECT_DATA             = 0x00000080
	XP1_DISCONNECT_DATA          = 0x00000100
	XP1_SUPPORT_BROADCAST        = 0x00000200
	XP1_SUPPORT_MULTIPOINT       = 0x00000400
	XP1_MULTIPOINT_CONTROL_PLANE = 0x00000800
	XP1_MULTIPOINT_DATA_PLANE    = 0x00001000
	XP1_QOS_SUPPORTED            = 0x00002000
	XP1_INTERRUPT                = 0x00004000
	XP1_UNI_SEND                 = 0x00008000
	XP1_UNI_RECV                 = 0x00010000
	XP1_IFS_HANDLES              = 0x00020000
	XP1_PARTIAL_MESSAGE          = 0x00040000
	XP1_SAN_SUPPORT_SDP          = 0x00080000
)

const (
	BIGENDIAN    = 0x0000
	LITTLEENDIAN = 0x0001
)

const SECURITY_PROTOCOL_NONE = 0x0000

// WinSock 2 extension -- manifest constants for WSAJoinLeaf()
const (
	JL_SENDER_ONLY   = 0x01
	JL_RECEIVER_ONLY = 0x02
	JL_BOTH          = 0x04
)

// WinSock 2 extension -- manifest constants for WSASocket()
const (
	WSA_FLAG_OVERLAPPED             = 0x01
	WSA_FLAG_MULTIPOINT_C_ROOT      = 0x02
	WSA_FLAG_MULTIPOINT_C_LEAF      = 0x04
	WSA_FLAG_MULTIPOINT_D_ROOT      = 0x08
	WSA_FLAG_MULTIPOINT_D_LEAF      = 0x10
	WSA_FLAG_ACCESS_SYSTEM_SECURITY = 0x40
	WSA_FLAG_NO_HANDLE_INHERIT      = 0x80
	WSA_FLAG_REGISTERED_IO          = 0x100
)

// WinSock 2 extension -- manifest constants for SIO_TRANSLATE_HANDLE ioctl
const (
	TH_NETDEV = 0x00000001
	TH_TAPI   = 0x00000002
)

// Service Install Flags
const SERVICE_MULTIPLE = 0x00000001

// Resolution flags for WSAGetAddressByName().
// Note these are also used by the 1.1 API GetAddressByName, so leave them around.
const (
	RES_UNUSED_1    = 0x00000001
	RES_FLUSH_CACHE = 0x00000002
	RES_SERVICE     = 0x00000004
)

// Well known value names for Service Types
const (
	SERVICE_TYPE_VALUE_IPXPORTA  = "IpxSocket"
	SERVICE_TYPE_VALUE_IPXPORTW  = "IpxSocket"
	SERVICE_TYPE_VALUE_SAPIDA    = "SapId"
	SERVICE_TYPE_VALUE_SAPIDW    = "SapId"
	SERVICE_TYPE_VALUE_TCPPORTA  = "TcpPort"
	SERVICE_TYPE_VALUE_TCPPORTW  = "TcpPort"
	SERVICE_TYPE_VALUE_UDPPORTA  = "UdpPort"
	SERVICE_TYPE_VALUE_UDPPORTW  = "UdpPort"
	SERVICE_TYPE_VALUE_OBJECTIDA = "ObjectId"
	SERVICE_TYPE_VALUE_OBJECTIDW = "ObjectId"
)

// Lookup flags for WSALookupServiceBegin/Next.
const (
	LUP_DEEP                     = 0x00000001
	LUP_CONTAINERS               = 0x00000002
	LUP_NOCONTAINERS             = 0x00000004
	LUP_NEAREST                  = 0x00000008
	LUP_RETURN_NAME              = 0x00000010
	LUP_RETURN_TYPE              = 0x00000020
	LUP_RETURN_VERSION           = 0x00000040
	LUP_RETURN_COMMENT           = 0x00000080
	LUP_RETURN_ADDR              = 0x00000100
	LUP_RETURN_BLOB              = 0x00000200
	LUP_RETURN_ALIASES           = 0x00000400
	LUP_RETURN_QUERY_STRING      = 0x00000800
	LUP_RETURN_ALL               = 0x00000FF0
	LUP_RES_SERVICE              = 0x00008000
	LUP_FLUSHCACHE               = 0x00001000
	LUP_FLUSHPREVIOUS            = 0x00002000
	LUP_NON_AUTHORITATIVE        = 0x00004000
	LUP_SECURE                   = 0x00008000
	LUP_RETURN_PREFERRED_NAMES   = 0x00010000
	LUP_DNS_ONLY                 = 0x00020000
	LUP_RETURN_RESPONSE_FLAGS    = 0x00040000
	LUP_ADDRCONFIG               = 0x00100000
	LUP_DUAL_ADDR                = 0x00200000
	LUP_FILESERVER               = 0x00400000
	LUP_DISABLE_IDN_ENCODING     = 0x00800000
	LUP_API_ANSI                 = 0x01000000
	LUP_EXTENDED_QUERYSET        = 0x02000000
	LUP_SECURE_WITH_FALLBACK     = 0x04000000
	LUP_EXCLUSIVE_CUSTOM_SERVERS = 0x08000000
	LUP_REQUIRE_SECURE           = 0x10000000
	LUP_RETURN_TTL               = 0x20000000
	LUP_FORCE_CLEAR_TEXT         = 0x40000000
	LUP_RESOLUTION_HANDLE        = 0x80000000
)

// Service Address Registration and Deregistration Data Types.
const (
	RNRSERVICE_REGISTER WSAESETSERVICEOP = iota
	RNRSERVICE_DEREGISTER
	RNRSERVICE_DELETE
)

// Return flags
const (
	RESULT_IS_ALIAS   = 0x0001
	RESULT_IS_ADDED   = 0x0010
	RESULT_IS_CHANGED = 0x0020
	RESULT_IS_DELETED = 0x0040
)

// Event flag definitions for WSAPoll().
const (
	POLLRDNORM = 0x0100
	POLLRDBAND = 0x0200
	POLLIN     = POLLRDNORM | POLLRDBAND
	POLLPRI    = 0x0400
	POLLWRNORM = 0x0010
	POLLOUT    = POLLWRNORM
	POLLWRBAND = 0x0020
	POLLERR    = 0x0001
	POLLHUP    = 0x0002
	POLLNVAL   = 0x0004
)

// Socket notification registration events. Supplied during registration.
// Indicates that a notification should be issued for the event if its
// condition holds.
const (
	SOCK_NOTIFY_REGISTER_EVENT_NONE   = 0x00
	SOCK_NOTIFY_REGISTER_EVENT_IN     = 0x01 // Input is available from the socket without blocking.
	SOCK_NOTIFY_REGISTER_EVENT_OUT    = 0x02 // Output can be provided to the socket without blocking.
	SOCK_NOTIFY_REGISTER_EVENT_HANGUP = 0x04 // The socket connection has been terminated.
)
const SOCK_NOTIFY_REGISTER_EVENTS_ALL = SOCK_NOTIFY_REGISTER_EVENT_IN | SOCK_NOTIFY_REGISTER_EVENT_OUT | SOCK_NOTIFY_REGISTER_EVENT_HANGUP

// Socket notification events. These are the events possible when a notification
// is received.
//
// The SOCK_NOTIFY_EVENT_ERR and SOCK_NOTIFY_EVENT_REMOVE events
// may be indicated regardless of registration.
//
// If a SOCK_NOTIFY_EVENT_REMOVE event is indicated, no more notifications will
// be provided.
const (
	SOCK_NOTIFY_EVENT_IN     = SOCK_NOTIFY_REGISTER_EVENT_IN     // Input is available from the socket without blocking.
	SOCK_NOTIFY_EVENT_OUT    = SOCK_NOTIFY_REGISTER_EVENT_OUT    // Output can be provided to the socket without blocking.
	SOCK_NOTIFY_EVENT_HANGUP = SOCK_NOTIFY_REGISTER_EVENT_HANGUP // The socket connection has been terminated.
	SOCK_NOTIFY_EVENT_ERR    = 0x40                              // The socket is in an error state.
	SOCK_NOTIFY_EVENT_REMOVE = 0x80                              // The notification has been deregistered.
)
const SOCK_NOTIFY_EVENTS_ALL = SOCK_NOTIFY_REGISTER_EVENTS_ALL | SOCK_NOTIFY_EVENT_ERR | SOCK_NOTIFY_EVENT_REMOVE

// Socket notification registration operations. One operation must be supplied at
// a time.
//
// A SOCK_NOTIFY_OP_DISABLE operation will not destroy the underlying structures.
//
// A SOCK_NOTIFY_OP_REMOVE operation will cause a SOCK_NOTIFY_REMOVE notification
// to be delivered when the operation completes successfully.
const (
	SOCK_NOTIFY_OP_NONE    = 0x00
	SOCK_NOTIFY_OP_ENABLE  = 0x01 // Enables the registration.
	SOCK_NOTIFY_OP_DISABLE = 0x02 // Disables the registration.
	SOCK_NOTIFY_OP_REMOVE  = 0x04 // Removes the registration.
)

// Socket notification trigger behaviors.
//
// When operation is SOCK_NOTIFY_OP_ENABLE:
//   - One of SOCK_NOTIFY_TRIGGER_PERSISTENT or SOCK_NOTIFY_TRIGGER_ONESHOT must be supplied
//   - One of SOCK_NOTIFY_TRIGGER_LEVEL or SOCK_NOTIFY_TRIGGER_EDGE must be supplied
//
// SOCK_NOTIFY_TRIGGER_PERSISTENT is not compatible with SOCK_NOTIFY_TRIGGER_ONESHOT.
// SOCK_NOTIFY_TRIGGER_LEVEL is not compatible with SOCK_NOTIFY_TRIGGER_EDGE.
//
// Socket notification trigger types.
const (
	SOCK_NOTIFY_TRIGGER_ONESHOT    = 0x01 // The registration will be disabled (not removed) upon delivery of the next notification.
	SOCK_NOTIFY_TRIGGER_PERSISTENT = 0x02 // The registration will remain active until it is explicitly disabled or removed.
	SOCK_NOTIFY_TRIGGER_LEVEL      = 0x04 // The registration is for level-triggered notifications.
	SOCK_NOTIFY_TRIGGER_EDGE       = 0x08 // The registration is for edge-triggered notifications.
)
const SOCK_NOTIFY_TRIGGER_ALL = SOCK_NOTIFY_TRIGGER_ONESHOT | SOCK_NOTIFY_TRIGGER_PERSISTENT | SOCK_NOTIFY_TRIGGER_LEVEL | SOCK_NOTIFY_TRIGGER_EDGE

/*
  MSWSock.h
*/

// Options for connect and disconnect data and options.  Used only by
// non-TCP/IP transports such as DECNet, OSI TP4, etc.
const (
	SO_CONNDATA    = 0x7000
	SO_CONNOPT     = 0x7001
	SO_DISCDATA    = 0x7002
	SO_DISCOPT     = 0x7003
	SO_CONNDATALEN = 0x7004
	SO_CONNOPTLEN  = 0x7005
	SO_DISCDATALEN = 0x7006
	SO_DISCOPTLEN  = 0x7007
)

// Option for opening sockets for synchronous access.
const (
	SO_OPENTYPE             = 0x7008
	SO_SYNCHRONOUS_ALERT    = 0x10
	SO_SYNCHRONOUS_NONALERT = 0x20
)

// Other NT-specific options.
const (
	SO_MAXDG                  = 0x7009
	SO_MAXPATHDG              = 0x700A
	SO_UPDATE_ACCEPT_CONTEXT  = 0x700B
	SO_CONNECT_TIME           = 0x700C
	SO_UPDATE_CONNECT_CONTEXT = 0x7010
)

// TCP options.
const TCP_BSDURGENT = 0x7000

// MS Transport Provider IOCTL to control
// reporting PORT_UNREACHABLE messages
// on UDP sockets via Recv/WSARecv/etc.
// Pass TRUE in input buffer to enable (default if supported),
// FALSE to disable.
const SIO_UDP_CONNRESET = IOC_IN | IOC_VENDOR | 12

// MS Transport Provider IOCTL to request
// notification when a given socket is closed.
// Input buffer must be a pointer to the socket handle.
// Input buffer size must be exactly sizeof(HANDLE).
// Output buffer and output buffer length must be
// NULL and 0 respectively. This IOCTL must always
// be issued with an overlapped structure.
//
// This Ioctl code is available only on WinXP SP2 and Win2k3 SP1.
const SIO_SOCKET_CLOSE_NOTIFY = IOC_IN | IOC_VENDOR | 13

// MS Transport Provider IOCTL to control
// reporting NET_UNREACHABLE (TTL expired) messages
// on UDP sockets via Recv/WSARecv/Etc.
// Pass TRUE in input buffer to enabled (default if supported),
// FALSE to disable.
const SIO_UDP_NETRESET = 0x9800000F

const (
	TF_DISCONNECT         = 0x01
	TF_REUSE_SOCKET       = 0x02
	TF_WRITE_BEHIND       = 0x04
	TF_USE_DEFAULT_WORKER = 0x00
	TF_USE_SYSTEM_THREAD  = 0x10
	TF_USE_KERNEL_APC     = 0x20
)

var WSAID_TRANSMITFILE = GUID{
	Data1: 0xb5367df0,
	Data2: 0xcbac,
	Data3: 0x11cf,
	Data4: [8]byte{0x95, 0xca, 0x00, 0x80, 0x5f, 0x48, 0xa1, 0x92},
}

var WSAID_ACCEPTEX = GUID{
	Data1: 0xb5367df1,
	Data2: 0xcbac,
	Data3: 0x11cf,
	Data4: [8]byte{0x95, 0xca, 0x00, 0x80, 0x5f, 0x48, 0xa1, 0x92},
}

var WSAID_GETACCEPTEXSOCKADDRS = GUID{
	Data1: 0xb5367df2,
	Data2: 0xcbac,
	Data3: 0x11cf,
	Data4: [8]byte{0x95, 0xca, 0x00, 0x80, 0x5f, 0x48, 0xa1, 0x92},
}

const (
	TP_ELEMENT_MEMORY     = 1
	TP_ELEMENT_FILE       = 2
	TP_ELEMENT_EOP        = 4
	TP_DISCONNECT         = TF_DISCONNECT
	TP_REUSE_SOCKET       = TF_REUSE_SOCKET
	TP_USE_DEFAULT_WORKER = TF_USE_DEFAULT_WORKER
	TP_USE_SYSTEM_THREAD  = TF_USE_SYSTEM_THREAD
	TP_USE_KERNEL_APC     = TF_USE_KERNEL_APC
)

var WSAID_TRANSMITPACKETS = GUID{
	Data1: 0xd9689da0,
	Data2: 0x1f90,
	Data3: 0x11d3,
	Data4: [8]byte{0x99, 0x71, 0x00, 0xc0, 0x4f, 0x68, 0xc8, 0x76},
}

var WSAID_CONNECTEX = GUID{
	Data1: 0x25a207b9,
	Data2: 0xddf3,
	Data3: 0x4660,
	Data4: [8]byte{0x8e, 0xe9, 0x76, 0xe5, 0x8c, 0x74, 0x06, 0x3e},
}

var WSAID_DISCONNECTEX = GUID{
	Data1: 0x7fda2e11,
	Data2: 0x8630,
	Data3: 0x436f,
	Data4: [8]byte{0xa0, 0x31, 0xf5, 0x36, 0xa6, 0xee, 0xc1, 0x57},
}

const DE_REUSE_SOCKET = TF_REUSE_SOCKET

// Network-location awareness -- Name registration values for use
// with WSASetService and other structures.
var NLA_NAMESPACE = GUID{
	Data1: 0x6642243a,
	Data2: 0x3ba8,
	Data3: 0x4aa6,
	Data4: [8]byte{0xba, 0xa5, 0x2e, 0x0b, 0xd7, 0x1f, 0xdd, 0x83},
}

var NLA_SERVICE_CLASS = GUID{
	Data1: 0x0037e515,
	Data2: 0xb5c9,
	Data3: 0x4a43,
	Data4: [8]byte{0xba, 0xda, 0x8b, 0x48, 0xa8, 0x7a, 0xd2, 0x39},
}

const NLA_ALLUSERS_NETWORK = 0x00000001

const NLA_FRIENDLY_NAME = 0x00000002

var WSAID_WSARECVMSG = GUID{
	Data1: 0xf689d7c8,
	Data2: 0x6f1f,
	Data3: 0x436b,
	Data4: [8]byte{0x8a, 0x53, 0xe5, 0x4f, 0xe3, 0x51, 0xc3, 0x22},
}

// Ioctl codes for translating socket handles to the base provider handle.
// This is performed to prevent breaking non-IFS LSPs when new Winsock extension
// funtions are added.
const (
	SIO_BSP_HANDLE        = 0x4800001B
	SIO_BSP_HANDLE_SELECT = 0x4800001C
	SIO_BSP_HANDLE_POLL   = 0x4800001D
)

// Ioctl code used to translate a socket handle into the base provider's handle.
// This is not used by any Winsock extension function and should not be intercepted
// by Winsock LSPs.
const (
	SIO_BASE_HANDLE = 0x48000022
)

// Ioctl codes for Winsock extension functions.
const (
	SIO_EXT_SELECT  = 0xC800001E
	SIO_EXT_POLL    = 0xC800001F
	SIO_EXT_SENDMSG = 0xC8000020
)

var WSAID_WSASENDMSG = GUID{
	Data1: 0xa441e712,
	Data2: 0x754f,
	Data3: 0x43ca,
	Data4: [8]byte{0x84, 0xa7, 0x0d, 0xee, 0x44, 0xcf, 0x60, 0x6d},
}

var WSAID_WSAPOLL = GUID{
	Data1: 0x18C76F85,
	Data2: 0xDC66,
	Data3: 0x4964,
	Data4: [8]byte{0x97, 0x2E, 0x23, 0xC2, 0x72, 0x38, 0x31, 0x2B},
}

var WSAID_MULTIPLE_RIO = GUID{
	Data1: 0x8509e081,
	Data2: 0x96dd,
	Data3: 0x4005,
	Data4: [8]byte{0xb1, 0x65, 0x9e, 0x2e, 0xe8, 0xc7, 0x9e, 0x3f},
}
