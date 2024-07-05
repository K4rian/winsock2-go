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
