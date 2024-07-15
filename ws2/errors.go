package ws2

import "strings"

// wsaErrPattern holds a pattern and the corresponding error code.
/*
type wsaErrPattern struct {
	pattern *regexp.Regexp
	code    int32
}
*/

/*
  WinSock error messages and codes
*/
var errorMap = map[string]int32{
	"A blocking operation was interrupted by a call to WSACancelBlockingCall.":                          WSAEINTR,        // WSAEINTR
	"The file handle supplied is not valid.":                                                            WSAEBADF,        // WSAEBADF
	"An attempt was made to access a socket in a way forbidden by its access permissions.":              WSAEACCES,       // WSAEACCES
	"The system detected an invalid pointer address in attempting to use a pointer argument in a call.": WSAEFAULT,       // WSAEFAULT
	"An invalid argument was supplied.":                                                                 WSAEINVAL,       // WSAEINVAL
	"Too many open sockets.":                                                                            WSAEMFILE,       // WSAEMFILE
	"A non-blocking socket operation could not be completed immediately.":                               WSAEWOULDBLOCK,  // WSAEWOULDBLOCK
	"A blocking operation is currently executing.":                                                      WSAEINPROGRESS,  // WSAEINPROGRESS
	"An operation was attempted on a non-blocking socket that already had an operation in progress.":    WSAEALREADY,     // WSAEALREADY
	"An operation was attempted on something that is not a socket.":                                     WSAENOTSOCK,     // WSAENOTSOCK
	"A required address was omitted from an operation on a socket.":                                     WSAEDESTADDRREQ, // WSAEDESTADDRREQ
	"A message sent on a datagram socket was larger than the internal message buffer or some other network limit, or the buffer used to receive a datagram into was smaller than the datagram itself.": WSAEMSGSIZE,        // WSAEMSGSIZE
	"A protocol was specified in the socket function call that does not support the semantics of the socket type requested.":                                                                           WSAEPROTOTYPE,      // WSAEPROTOTYPE
	"An unknown, invalid, or unsupported option or level was specified in a getsockopt or setsockopt call.":                                                                                            WSAENOPROTOOPT,     // WSAENOPROTOOPT
	"The requested protocol has not been configured into the system, or no implementation for it exists.":                                                                                              WSAEPROTONOSUPPORT, // WSAEPROTONOSUPPORT
	"The support for the specified socket type does not exist in this address family.":                                                                                                                 WSAESOCKTNOSUPPORT, // WSAESOCKTNOSUPPORT
	"The attempted operation is not supported for the type of object referenced.":                                                                                                                      WSAEOPNOTSUPP,      // WSAEOPNOTSUPP
	"The protocol family has not been configured into the system or no implementation for it exists.":                                                                                                  WSAEPFNOSUPPORT,    // WSAEPFNOSUPPORT
	"An address incompatible with the requested protocol was used.":                                                                                                                                    WSAEAFNOSUPPORT,    // WSAEAFNOSUPPORT
	"Only one usage of each socket address (protocol/network address/port) is normally permitted.":                                                                                                     WSAEADDRINUSE,      // WSAEADDRINUSE
	"The requested address is not valid in its context.":                                                                                                                                               WSAEADDRNOTAVAIL,   // WSAEADDRNOTAVAIL
	"A socket operation encountered a dead network.":                                                                                                                                                   WSAENETDOWN,        // WSAENETDOWN
	"A socket operation was attempted to an unreachable network.":                                                                                                                                      WSAENETUNREACH,     // WSAENETUNREACH
	"The connection has been broken due to keep-alive activity detecting a failure while the operation was in progress.":                                                                               WSAENETRESET,       // WSAENETRESET
	"An established connection was aborted by the software in your host machine.":                                                                                                                      WSAECONNABORTED,    // WSAECONNABORTED
	"An existing connection was forcibly closed by the remote host.":                                                                                                                                   WSAECONNRESET,      // WSAECONNRESET
	"An operation on a socket could not be performed because the system lacked sufficient buffer space or because a queue was full.":                                                                   WSAENOBUFS,         // WSAENOBUFS
	"A connect request was made on an already connected socket.":                                                                                                                                       WSAEISCONN,         // WSAEISCONN
	"A request to send or receive data was disallowed because the socket is not connected and (when sending on a datagram socket using a sendto call) no address was supplied.":                        WSAENOTCONN,        // WSAENOTCONN
	"A request to send or receive data was disallowed because the socket had already been shut down in that direction with a previous shutdown call.":                                                  WSAESHUTDOWN,       // WSAESHUTDOWN
	"Too many references to some kernel object.":                                                                                                                                                       WSAETOOMANYREFS,    // WSAETOOMANYREFS
	"A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.":          WSAETIMEDOUT,       // WSAETIMEDOUT
	"No connection could be made because the target machine actively refused it.":                                                                                                                      WSAECONNREFUSED,    // WSAECONNREFUSED
	"Cannot translate name.":                                           WSAELOOP,        // WSAELOOP
	"Name component or name was too long.":                             WSAENAMETOOLONG, // WSAENAMETOOLONG
	"A socket operation failed because the destination host was down.": WSAEHOSTDOWN,    // WSAEHOSTDOWN
	"A socket operation was attempted to an unreachable host.":         WSAEHOSTUNREACH, // WSAEHOSTUNREACH
	"Cannot remove a directory that is not empty.":                     WSAENOTEMPTY,    // WSAENOTEMPTY
	"A Windows Sockets implementation may have a limit on the number of applications that may use it simultaneously.": WSAEPROCLIM, // WSAEPROCLIM
	"Ran out of quota.":                             WSAEUSERS,  // WSAEUSERS
	"Ran out of disk quota.":                        WSAEDQUOT,  // WSAEDQUOT
	"File handle reference is no longer available.": WSAESTALE,  // WSAESTALE
	"Item is not available locally.":                WSAEREMOTE, // WSAEREMOTE
	"WSAStartup cannot function at this time because the underlying system it uses to provide network services is currently unavailable.": WSASYSNOTREADY,         // WSASYSNOTREADY
	"The Windows Sockets version requested is not supported.":                                                                             WSAVERNOTSUPPORTED,     // WSAVERNOTSUPPORTED
	"Either the application has not called WSAStartup, or WSAStartup failed.":                                                             WSANOTINITIALISED,      // WSANOTINITIALISED
	"Returned by WSARecv or WSARecvFrom to indicate the remote party has initiated a graceful shutdown sequence.":                         WSAEDISCON,             // WSAEDISCON
	"No more results can be returned by WSALookupServiceNext.":                                                                            WSAENOMORE,             // WSAENOMORE / WSA_E_NO_MORE
	"A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled.":                            WSAECANCELLED,          // WSAECANCELLED / WSA_E_CANCELLED
	"The procedure call table is invalid.":                                                                                                WSAEINVALIDPROCTABLE,   // WSAEINVALIDPROCTABLE
	"The requested service provider is invalid.":                                                                                          WSAEINVALIDPROVIDER,    // WSAEINVALIDPROVIDER
	"The requested service provider could not be loaded or initialized.":                                                                  WSAEPROVIDERFAILEDINIT, // WSAEPROVIDERFAILEDINIT
	"A system call has failed.": WSASYSCALLFAILURE, // WSASYSCALLFAILURE
	"No such service is known. The service cannot be found in the specified name space.": WSASERVICE_NOT_FOUND, // WSASERVICE_NOT_FOUND
	"The specified class was not found.":                                                 WSATYPE_NOT_FOUND,    // WSATYPE_NOT_FOUND
	// "No more results can be returned by WSALookupServiceNext.": WSA_E_NO_MORE,                                                                                             // WSA_E_NO_MORE
	// "A call to WSALookupServiceEnd was made while this call was still processing. The call has been canceled.": WSA_E_CANCELLED,                                           // WSA_E_CANCELLED
	"A database query failed because it was actively refused.": WSAEREFUSED,       // WSAEREFUSED
	"No such host is known.":                                   WSAHOST_NOT_FOUND, // WSAHOST_NOT_FOUND
	"This is usually a temporary error during hostname resolution and means that the local server did not receive a response from an authoritative server.": WSATRY_AGAIN, // WSATRY_AGAIN
	"A non-recoverable error occurred during a database lookup.":                                      WSANO_RECOVERY,              // WSANO_RECOVERY
	"The requested name is valid, but no data of the requested type was found.":                       WSANO_DATA,                  // WSANO_DATA
	"At least one reserve has arrived.":                                                               WSA_QOS_RECEIVERS,           // WSA_QOS_RECEIVERS
	"At least one path has arrived.":                                                                  WSA_QOS_SENDERS,             // WSA_QOS_SENDERS
	"There are no senders.":                                                                           WSA_QOS_NO_SENDERS,          // WSA_QOS_NO_SENDERS
	"There are no receivers.":                                                                         WSA_QOS_NO_RECEIVERS,        // WSA_QOS_NO_RECEIVERS
	"Reserve has been confirmed.":                                                                     WSA_QOS_REQUEST_CONFIRMED,   // WSA_QOS_REQUEST_CONFIRMED
	"Error due to lack of resources.":                                                                 WSA_QOS_ADMISSION_FAILURE,   // WSA_QOS_ADMISSION_FAILURE
	"Rejected for administrative reasons - bad credentials.":                                          WSA_QOS_POLICY_FAILURE,      // WSA_QOS_POLICY_FAILURE
	"Unknown or conflicting style.":                                                                   WSA_QOS_BAD_STYLE,           // WSA_QOS_BAD_STYLE
	"Problem with some part of the filterspec or providerspecific buffer in general.":                 WSA_QOS_BAD_OBJECT,          // WSA_QOS_BAD_OBJECT
	"Problem with some part of the flowspec.":                                                         WSA_QOS_TRAFFIC_CTRL_ERROR,  // WSA_QOS_TRAFFIC_CTRL_ERROR
	"General QOS error.":                                                                              WSA_QOS_GENERIC_ERROR,       // WSA_QOS_GENERIC_ERROR
	"An invalid or unrecognized service type was found in the flowspec.":                              WSA_QOS_ESERVICETYPE,        // WSA_QOS_ESERVICETYPE
	"An invalid or inconsistent flowspec was found in the QOS structure.":                             WSA_QOS_EFLOWSPEC,           // WSA_QOS_EFLOWSPEC
	"Invalid QOS provider-specific buffer.":                                                           WSA_QOS_EPROVSPECBUF,        // WSA_QOS_EPROVSPECBUF
	"An invalid QOS filter style was used.":                                                           WSA_QOS_EFILTERSTYLE,        // WSA_QOS_EFILTERSTYLE
	"An invalid QOS filter type was used.":                                                            WSA_QOS_EFILTERTYPE,         // WSA_QOS_EFILTERTYPE
	"An incorrect number of QOS FILTERSPECs were specified in the FLOWDESCRIPTOR.":                    WSA_QOS_EFILTERCOUNT,        // WSA_QOS_EFILTERCOUNT
	"An object with an invalid ObjectLength field was specified in the QOS provider-specific buffer.": WSA_QOS_EOBJLENGTH,          // WSA_QOS_EOBJLENGTH
	"An incorrect number of flow descriptors was specified in the QOS structure.":                     WSA_QOS_EFLOWCOUNT,          // WSA_QOS_EFLOWCOUNT
	"An unrecognized object was found in the QOS provider-specific buffer.":                           WSA_QOS_EUNKOWNPSOBJ,        // WSA_QOS_EUNKOWNPSOBJ
	"An invalid policy object was found in the QOS provider-specific buffer.":                         WSA_QOS_EPOLICYOBJ,          // WSA_QOS_EPOLICYOBJ
	"An invalid QOS flow descriptor was found in the flow descriptor list.":                           WSA_QOS_EFLOWDESC,           // WSA_QOS_EFLOWDESC
	"An invalid or inconsistent flowspec was found in the QOS provider specific buffer.":              WSA_QOS_EPSFLOWSPEC,         // WSA_QOS_EPSFLOWSPEC
	"An invalid FILTERSPEC was found in the QOS provider-specific buffer.":                            WSA_QOS_EPSFILTERSPEC,       // WSA_QOS_EPSFILTERSPEC
	"An invalid shape discard mode object was found in the QOS provider specific buffer.":             WSA_QOS_ESDMODEOBJ,          // WSA_QOS_ESDMODEOBJ
	"An invalid shaping rate object was found in the QOS provider-specific buffer.":                   WSA_QOS_ESHAPERATEOBJ,       // WSA_QOS_ESHAPERATEOBJ
	"A reserved policy element was found in the QOS provider-specific buffer.":                        WSA_QOS_RESERVED_PETYPE,     // WSA_QOS_RESERVED_PETYPE
	"No such host is known securely.":                                                                 WSA_SECURE_HOST_NOT_FOUND,   // WSA_SECURE_HOST_NOT_FOUND
	"Name based IPSEC policy could not be added.":                                                     WSA_IPSEC_NAME_POLICY_ERROR, // WSA_IPSEC_NAME_POLICY_ERROR
}

/*
var errorPatterns = []wsaErrPattern{
	{regexp.MustCompile(`(?i)operation was interrupted by a call\.?`), WSAEINTR},                             // WSAEINTR
	{regexp.MustCompile(`(?i)file handle supplied is not\.?`), WSAEBADF},                                     // WSAEBADF
	{regexp.MustCompile(`(?i)was made to access a socket in a way forbidden by\.?`), WSAEACCES},              // WSAEACCES
	{regexp.MustCompile(`(?i)invalid pointer address in attempting to use a pointer\.?`), WSAEFAULT},         // WSAEFAULT
	{regexp.MustCompile(`(?i)invalid argument was supplied\.?`), WSAEINVAL},                                  // WSAEINVAL
	{regexp.MustCompile(`(?i)many open sockets\.?`), WSAEMFILE},                                              // WSAEMFILE
	{regexp.MustCompile(`(?i)blocking socket operation could not be\.?`), WSAEWOULDBLOCK},                    // WSAEWOULDBLOCK
	{regexp.MustCompile(`(?i)operation is currently executing\.?`), WSAEINPROGRESS},                          // WSAEINPROGRESS
	{regexp.MustCompile(`(?i)operation was attempted on a non-blocking socket\.?`), WSAEALREADY},             // WSAEALREADY
	{regexp.MustCompile(`(?i)operation was attempted on something\.?`), WSAENOTSOCK},                         // WSAENOTSOCK
	{regexp.MustCompile(`(?i)address was omitted from\.?`), WSAEDESTADDRREQ},                                 // WSAEDESTADDRREQ
	{regexp.MustCompile(`(?i)datagram socket was larger than the internal\.?`), WSAEMSGSIZE},                 // WSAEMSGSIZE
	{regexp.MustCompile(`(?i)protocol was specified in the socket\.?`), WSAEPROTOTYPE},                       // WSAEPROTOTYPE
	{regexp.MustCompile(`(?i)unknown, invalid, or unsupported\.?`), WSAENOPROTOOPT},                          // WSAENOPROTOOPT
	{regexp.MustCompile(`(?i)protocol has not been configured\.?`), WSAEPROTONOSUPPORT},                      // WSAEPROTONOSUPPORT
	{regexp.MustCompile(`(?i)support for the specified socket type\.?`), WSAESOCKTNOSUPPORT},                 // WSAESOCKTNOSUPPORT
	{regexp.MustCompile(`(?i)attempted operation is not supported\.?`), WSAEOPNOTSUPP},                       // WSAEOPNOTSUPP
	{regexp.MustCompile(`(?i)protocol family has not been configured\.?`), WSAEPFNOSUPPORT},                  // WSAEPFNOSUPPORT
	{regexp.MustCompile(`(?i)address incompatible with the requested\.?`), WSAEAFNOSUPPORT},                  // WSAEAFNOSUPPORT
	{regexp.MustCompile(`(?i)one usage of each socket address\.?`), WSAEADDRINUSE},                           // WSAEADDRINUSE
	{regexp.MustCompile(`(?i)requested address is not valid\.?`), WSAEADDRNOTAVAIL},                          // WSAEADDRNOTAVAIL
	{regexp.MustCompile(`(?i)socket operation encountered\.?`), WSAENETDOWN},                                 // WSAENETDOWN
	{regexp.MustCompile(`(?i)attempted to an unreachable network\.?`), WSAENETUNREACH},                       // WSAENETUNREACH
	{regexp.MustCompile(`(?i)connection has been broken\.?`), WSAENETRESET},                                  // WSAENETRESET
	{regexp.MustCompile(`(?i)established connection was aborted\.?`), WSAECONNABORTED},                       // WSAECONNABORTED
	{regexp.MustCompile(`(?i)existing connection was forcibly closed\.?`), WSAECONNRESET},                    // WSAECONNRESET
	{regexp.MustCompile(`(?i)socket could not be performed\.?`), WSAENOBUFS},                                 // WSAENOBUFS
	{regexp.MustCompile(`(?i)connect request was made on an already\.?`), WSAEISCONN},                        // WSAEISCONN
	{regexp.MustCompile(`(?i)data was disallowed because the socket is not\.?`), WSAENOTCONN},                // WSAENOTCONN
	{regexp.MustCompile(`(?i)data was disallowed because the socket had already\.?`), WSAESHUTDOWN},          // WSAESHUTDOWN
	{regexp.MustCompile(`(?i)references to some kernel\.?`), WSAETOOMANYREFS},                                // WSAETOOMANYREFS
	{regexp.MustCompile(`(?i)attempt failed because the connected\.?`), WSAETIMEDOUT},                        // WSAETIMEDOUT
	{regexp.MustCompile(`(?i)because the target machine\.?`), WSAECONNREFUSED},                               // WSAECONNREFUSED
	{regexp.MustCompile(`(?i)translate\.?`), WSAELOOP},                                                       // WSAELOOP
	{regexp.MustCompile(`(?i)component or name\.?`), WSAENAMETOOLONG},                                        // WSAENAMETOOLONG
	{regexp.MustCompile(`(?i)operation failed because the destination\.?`), WSAEHOSTDOWN},                    // WSAEHOSTDOWN
	{regexp.MustCompile(`(?i)attempted to an unreachable host\.?`), WSAEHOSTUNREACH},                         // WSAEHOSTUNREACH
	{regexp.MustCompile(`(?i)remove a directory\.?`), WSAENOTEMPTY},                                          // WSAENOTEMPTY
	{regexp.MustCompile(`(?i)implementation may have a limit\.?`), WSAEPROCLIM},                              // WSAEPROCLIM
	{regexp.MustCompile(`(?i)out of quota\.?`), WSAEUSERS},                                                   // WSAEUSERS
	{regexp.MustCompile(`(?i)out of disk quota\.?`), WSAEDQUOT},                                              // WSAEDQUOT
	{regexp.MustCompile(`(?i)handle reference is\.?`), WSAESTALE},                                            // WSAESTALE
	{regexp.MustCompile(`(?i)not available locally\.?`), WSAEREMOTE},                                         // WSAEREMOTE
	{regexp.MustCompile(`(?i)underlying system it uses to provide\.?`), WSASYSNOTREADY},                      // WSASYSNOTREADY
	{regexp.MustCompile(`(?i)version requested is not\.?`), WSAVERNOTSUPPORTED},                              // WSAVERNOTSUPPORTED
	{regexp.MustCompile(`(?i)has not called WSAStartup\.?`), WSANOTINITIALISED},                              // WSANOTINITIALISED
	{regexp.MustCompile(`(?i)WSARecv or WSARecvFrom to indicate\.?`), WSAEDISCON},                            // WSAEDISCON
	{regexp.MustCompile(`(?i)results can be returned\.?`), WSAENOMORE},                                       // WSAENOMORE / WSA_E_NO_MORE
	{regexp.MustCompile(`(?i)WSALookupServiceEnd was made while\.?`), WSAECANCELLED},                         // WSAECANCELLED / WSA_E_CANCELLED
	{regexp.MustCompile(`(?i)call table is\.?`), WSAEINVALIDPROCTABLE},                                       // WSAEINVALIDPROCTABLE
	{regexp.MustCompile(`(?i)requested service provider is\.?`), WSAEINVALIDPROVIDER},                        // WSAEINVALIDPROVIDER
	{regexp.MustCompile(`(?i)requested service provider could\.?`), WSAEPROVIDERFAILEDINIT},                  // WSAEPROVIDERFAILEDINIT
	{regexp.MustCompile(`(?i)system call has\.?`), WSASYSCALLFAILURE},                                        // WSASYSCALLFAILURE
	{regexp.MustCompile(`(?i)such service is known\.?`), WSASERVICE_NOT_FOUND},                               // WSASERVICE_NOT_FOUND
	{regexp.MustCompile(`(?i)specified class was not\.?`), WSATYPE_NOT_FOUND},                                // WSATYPE_NOT_FOUND
	{regexp.MustCompile(`(?i)more results can be returned\.?`), WSA_E_NO_MORE},                               // WSA_E_NO_MORE
	{regexp.MustCompile(`(?i)WSALookupServiceEnd was made while\.?`), WSA_E_CANCELLED},                       // WSA_E_CANCELLED
	{regexp.MustCompile(`(?i)database query failed because\.?`), WSAEREFUSED},                                // WSAEREFUSED
	{regexp.MustCompile(`(?i)\s*No such host is known\.\s*`), WSAHOST_NOT_FOUND},                             // WSAHOST_NOT_FOUND
	{regexp.MustCompile(`(?i)usually a temporary error\.?`), WSATRY_AGAIN},                                   // WSATRY_AGAIN
	{regexp.MustCompile(`(?i)occurred during a database\.?`), WSANO_RECOVERY},                                // WSANO_RECOVERY
	{regexp.MustCompile(`(?i)requested name is valid, but no\.?`), WSANO_DATA},                               // WSANO_DATA
	{regexp.MustCompile(`(?i)least one reserve\.?`), WSA_QOS_RECEIVERS},                                      // WSA_QOS_RECEIVERS
	{regexp.MustCompile(`(?i)least one path\.?`), WSA_QOS_SENDERS},                                           // WSA_QOS_SENDERS
	{regexp.MustCompile(`(?i)are no senders\.?`), WSA_QOS_NO_SENDERS},                                        // WSA_QOS_NO_SENDERS
	{regexp.MustCompile(`(?i)are no receivers\.?`), WSA_QOS_NO_RECEIVERS},                                    // WSA_QOS_NO_RECEIVERS
	{regexp.MustCompile(`(?i)has been confirmed\.?`), WSA_QOS_REQUEST_CONFIRMED},                             // WSA_QOS_REQUEST_CONFIRMED
	{regexp.MustCompile(`(?i)due to lack of\.?`), WSA_QOS_ADMISSION_FAILURE},                                 // WSA_QOS_ADMISSION_FAILURE
	{regexp.MustCompile(`(?i)for administrative reasons\.?`), WSA_QOS_POLICY_FAILURE},                        // WSA_QOS_POLICY_FAILURE
	{regexp.MustCompile(`(?i)or conflicting\.?`), WSA_QOS_BAD_STYLE},                                         // WSA_QOS_BAD_STYLE
	{regexp.MustCompile(`(?i)the filterspec or providerspecific\.?`), WSA_QOS_BAD_OBJECT},                    // WSA_QOS_BAD_OBJECT
	{regexp.MustCompile(`(?i)some part of the flowspec\.?`), WSA_QOS_TRAFFIC_CTRL_ERROR},                     // WSA_QOS_TRAFFIC_CTRL_ERROR
	{regexp.MustCompile(`(?i)QOS error\.?`), WSA_QOS_GENERIC_ERROR},                                          // WSA_QOS_GENERIC_ERROR
	{regexp.MustCompile(`(?i)unrecognized service type was found in the flowspec\.?`), WSA_QOS_ESERVICETYPE}, // WSA_QOS_ESERVICETYPE
	{regexp.MustCompile(`(?i)inconsistent flowspec was found in the QOS structure\.?`), WSA_QOS_EFLOWSPEC},   // WSA_QOS_EFLOWSPEC
	{regexp.MustCompile(`(?i)nvalid QOS provider\.?`), WSA_QOS_EPROVSPECBUF},                                 // WSA_QOS_EPROVSPECBUF
	{regexp.MustCompile(`(?i)invalid QOS filter style\.?`), WSA_QOS_EFILTERSTYLE},                            // WSA_QOS_EFILTERSTYLE
	{regexp.MustCompile(`(?i)invalid QOS filter type\.?`), WSA_QOS_EFILTERTYPE},                              // WSA_QOS_EFILTERTYPE
	{regexp.MustCompile(`(?i)number of QOS FILTERSPECs were specified\.?`), WSA_QOS_EFILTERCOUNT},            // WSA_QOS_EFILTERCOUNT
	{regexp.MustCompile(`(?i)invalid ObjectLength field\.?`), WSA_QOS_EOBJLENGTH},                            // WSA_QOS_EOBJLENGTH
	{regexp.MustCompile(`(?i)flow descriptors was specified\.?`), WSA_QOS_EFLOWCOUNT},                        // WSA_QOS_EFLOWCOUNT
	{regexp.MustCompile(`(?i)unrecognized object was found\.?`), WSA_QOS_EUNKOWNPSOBJ},                       // WSA_QOS_EUNKOWNPSOBJ
	{regexp.MustCompile(`(?i)policy object was found in the QOS\.?`), WSA_QOS_EPOLICYOBJ},                    // WSA_QOS_EPOLICYOBJ
	{regexp.MustCompile(`(?i)flow descriptor was found in the flow descriptor\.?`), WSA_QOS_EFLOWDESC},       // WSA_QOS_EFLOWDESC
	{regexp.MustCompile(`(?i)inconsistent flowspec was found in the QOS provider\.?`), WSA_QOS_EPSFLOWSPEC},  // WSA_QOS_EPSFLOWSPEC
	{regexp.MustCompile(`(?i)FILTERSPEC was found in the QOS provider\.?`), WSA_QOS_EPSFILTERSPEC},           // WSA_QOS_EPSFILTERSPEC
	{regexp.MustCompile(`(?i)shape discard mode object was\.?`), WSA_QOS_ESDMODEOBJ},                         // WSA_QOS_ESDMODEOBJ
	{regexp.MustCompile(`(?i)shaping rate object was found\.?`), WSA_QOS_ESHAPERATEOBJ},                      // WSA_QOS_ESHAPERATEOBJ
	{regexp.MustCompile(`(?i)reserved policy element was found\.?`), WSA_QOS_RESERVED_PETYPE},                // WSA_QOS_RESERVED_PETYPE
	{regexp.MustCompile(`(?i)host is known securely\.?`), WSA_SECURE_HOST_NOT_FOUND},                         // WSA_SECURE_HOST_NOT_FOUND
	{regexp.MustCompile(`(?i)based IPSEC policy\.?`), WSA_IPSEC_NAME_POLICY_ERROR},                           // WSA_IPSEC_NAME_POLICY_ERROR
}
*/

// WSAGetErrorCode returns the corresponding error code of the given WinSocket error message or 0 if not found.
func WSAGetErrorCode(err error) int32 {
	errStr := strings.TrimSpace(err.Error())
	if code, found := errorMap[errStr]; found {
		return code
	}
	return 0
}

// WSAGetErrorCodeEx returns the corresponding error code of the given WinSocket error message using pattern matching or 0 if not found.
/*
func WSAGetErrorCodeEx(err error) int32 {
	errStr := err.Error()
	for _, ep := range errorPatterns {
		if ep.pattern.MatchString(errStr) {
			return ep.code
		}
	}
	return 0
}
*/
