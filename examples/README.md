winsock2-go / examples
=====

TCP                                       | &nbsp;
:---                                      | :---
[echo_tcp/ds_server](echo_tcp/ds_server/) | Echo Dual-Stack TCP Server&nbsp;&nbsp;
[echo_tcp/v4_client](echo_tcp/v4_client/) | Echo v4 TCP Client
[echo_tcp/v4_server](echo_tcp/v4_server/) | Echo v4 TCP Server
[echo_tcp/v6_client](echo_tcp/v6_client/) | Echo v6 TCP Client
[echo_tcp/v6_server](echo_tcp/v6_server/) | Echo v6 TCP Server

UDP                                       | &nbsp;
:---                                      | :---
[echo_udp/ds_server](echo_udp/ds_server/) | Echo Dual-Stack UDP Server
[echo_udp/v4_client](echo_udp/v4_client/) | Echo v4 UDP Client
[echo_udp/v4_server](echo_udp/v4_server/) | Echo v4 UDP Server
[echo_udp/v6_client](echo_udp/v6_client/) | Echo v6 UDP Client
[echo_udp/v6_server](echo_udp/v6_server/) | Echo v6 UDP Server


| API Function Snippets                                                                                                  |
| :---                                                                                                                   |
| [__WSAFDIsSet](https://github.com/K4rian/winsock2-go/blob/main/ws2/types.go#L58)                                       | 
| [Accept](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L126)                   | 
| [Bind](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L100)                     | 
| [CloseSocket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L87)               | 
| [Connect](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L78)                   | 
| [FreeAddrInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L38)                     | 
| [FreeAddrInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L84)                     | 
| [GetAddrInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L18)                      | 
| [GetAddrInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L64)                      | 
| [GetHostByAddr](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L19)                      | 
| [GetHostByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L35)                      | 
| [GetHostNameA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L49)                       | 
| [GetHostNameW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L57)                       | 
| [GetNameInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getnameinfo.go#L17)                   | 
| [GetNameInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getnameinfo.go#L44)                   | 
| [GetPeerName](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L180)              | 
| [GetProtoByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getproto.go#L15)                    | 
| [GetProtoByNumber](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getproto.go#L35)                  | 
| [GetServByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getserv.go#L15)                      | 
| [GetServByPort](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getserv.go#L38)                      | 
| [GetSockName](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L193)              | 
| [GetSockOpt](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L203)               | 
| [Htond](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L20)                             | 
| [Htonf](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L25)                             | 
| [Htonl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L30)                             | 
| [Htonll](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L35)                            | 
| [Htons](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L40)                             | 
| [InetAddr](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L22)                              | 
| [InetNtoa](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L45)                              | 
| [InetNtop](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L58)                              | 
| [InetNtopW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L80)                             | 
| [InetPton](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L102)                             | 
| [InetPtonW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L120)                            | 
| [IoctlSocket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L44)               | 
| [Listen](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L107)                   | 
| [Ntohl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L45)                             | 
| [Ntohs](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L50)                             | 
| [Recv](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L223)                     | 
| [RecvFrom](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_udp/v6_server/server.go#L121)                 | 
| [Select](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L137)                   | 
| [Send](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L234)                     | 
| [SendTo](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_udp/v6_server/server.go#L135)                   | 
| [SetSockOpt](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L212)               | 
| [Shutdown](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L86)                  | 
| [Socket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L79)                    | 
| WSAAccept                                                                                                              | 
| [WSAAddressToStringA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L23)                | 
| [WSAAddressToStringW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L63)                | 
| [WSACleanup](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L60)                | 
| [WSACloseEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L94)             | 
| [WSAConnect](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L20)                      | 
| [WSAConnectByNameA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L58)               | 
| [WSAConnectByNameW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L101)              | 
| [WSACreateEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L89)            | 
| [WSADuplicateSocketA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L25)              | 
| [WSADuplicateSocketW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L58)              | 
| [WSAEnumNameSpaceProvidersA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsanamespace.go#L17)    | 
| WSAEnumNameSpaceProvidersExA                                                                                           | 
| WSAEnumNameSpaceProvidersExW                                                                                           | 
| [WSAEnumNameSpaceProvidersW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsanamespace.go#L40)    | 
| [WSAEnumNetworkEvents](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L103)     | 
| [WSAEnumProtocolsA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L28)                  | 
| [WSAEnumProtocolsW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L68)                  | 
| [WSAEventSelect](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L98)            | 
| [WSAGetOverlappedResult](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L91)           | 
| [WSAHtonl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L19)                       | 
| [WSAHtons](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L28)                       | 
| [WSAIoctl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L117)                        | 
| WSAJoinLeaf                                                                                                            | 
| [WSALookupServiceBeginA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L60)             | 
| WSALookupServiceBeginW                                                                                                 | 
| [WSALookupServiceEnd](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L65)                | 
| [WSALookupServiceNextA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L72)              | 
| WSANSPIoctl                                                                                                            | 
| [WSANtohl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L37)                       | 
| [WSANtohs](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L46)                       | 
| WSAPoll                                                                                                                | 
| WSAProviderConfigChange                                                                                                | 
| WSARecv                                                                                                                | 
| WSARecvDisconnect                                                                                                      | 
| WSARecvFrom                                                                                                            | 
| WSARemoveServiceClass                                                                                                  | 
| [WSAResetEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L126)            | 
| WSASend                                                                                                                | 
| WSASendDisconnect                                                                                                      | 
| WSASendMsg                                                                                                             | 
| WSASendTo                                                                                                              | 
| [WSASetEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L113)              | 
| WSASetLastError                                                                                                        | 
| WSASetServiceA                                                                                                         | 
| WSASetServiceW                                                                                                         | 
| [WSASocketA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L192)                      | 
| [WSASocketW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L204)                      | 
| [WSAStartup](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L31)                | 
| [WSAStringToAddressA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L103)               | 
| [WSAStringToAddressW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L139)               | 
| [WSAWaitForMultipleEvents](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L118) | 