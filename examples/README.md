winsock2-go / examples
=====

TCP                                       | &nbsp;
:---                                      | :---
[echo_tcp/ds_server](echo_tcp/ds_server/) | Echo Dual-Stack TCP Server
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

API Function Snippets                                                                                                  | &nbsp;
:---                                                                                                                   | :---
[__WSAFDIsSet](https://github.com/K4rian/winsock2-go/blob/main/ws2/types.go#L58)                                       | &nbsp;
[Accept](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L126)                   | &nbsp;
[Bind](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L100)                     | &nbsp;
[CloseSocket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L87)               | &nbsp;
[Connect](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L78)                   | &nbsp;
[FreeAddrInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L38)                     | &nbsp;
[FreeAddrInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L84)                     | &nbsp;
[GetAddrInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L18)                      | &nbsp;
[GetAddrInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/addrinfo.go#L64)                      | &nbsp;
[GetHostByAddr](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L19)                      | &nbsp;
[GetHostByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L35)                      | &nbsp;
[GetHostNameA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L49)                       | &nbsp;
[GetHostNameW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/gethost.go#L57)                       | &nbsp;
[GetNameInfoA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getnameinfo.go#L17)                   | &nbsp;
[GetNameInfoW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getnameinfo.go#L44)                   | &nbsp;
[GetPeerName](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L180)              | &nbsp;
[GetProtoByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getproto.go#L15)                    | &nbsp;
[GetProtoByNumber](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getproto.go#L35)                  | &nbsp;
[GetServByName](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getserv.go#L15)                      | &nbsp;
[GetServByPort](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/getserv.go#L38)                      | &nbsp;
[GetSockName](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L193)              | &nbsp;
[GetSockOpt](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L203)               | &nbsp;
[Htond](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L20)                             | &nbsp;
[Htonf](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L25)                             | &nbsp;
[Htonl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L30)                             | &nbsp;
[Htonll](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L35)                            | &nbsp;
[Htons](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L40)                             | &nbsp;
[InetAddr](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L22)                              | &nbsp;
[InetNtoa](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L45)                              | &nbsp;
[InetNtop](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L58)                              | &nbsp;
[InetNtopW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L80)                             | &nbsp;
[InetPton](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L102)                             | &nbsp;
[InetPtonW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/inet.go#L120)                            | &nbsp;
[IoctlSocket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L44)               | &nbsp;
[Listen](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L107)                   | &nbsp;
[Ntohl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L45)                             | &nbsp;
[Ntohs](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/htonntoh.go#L50)                             | &nbsp;
[Recv](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L223)                     | &nbsp;
[RecvFrom](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_udp/v6_server/server.go#L121)                 | &nbsp;
[Select](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L137)                   | &nbsp;
[Send](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L234)                     | &nbsp;
[SendTo](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_udp/v6_server/server.go#L135)                   | &nbsp;
[SetSockOpt](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L212)               | &nbsp;
[Shutdown](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L86)                  | &nbsp;
[Socket](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L79)                    | &nbsp;
WSAAccept                                                                                                              | &nbsp;
[WSAAddressToStringA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L21)                | &nbsp;
[WSAAddressToStringW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L61)                | &nbsp;
[WSACleanup](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L60)                | &nbsp;
[WSACloseEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L94)             | &nbsp;
[WSAConnect](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L20)                      | &nbsp;
[WSAConnectByNameA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L58)               | &nbsp;
[WSAConnectByNameW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaconnect.go#L101)              | &nbsp;
[WSACreateEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L89)            | &nbsp;
[WSADuplicateSocketA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L25)              | &nbsp;
[WSADuplicateSocketW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L58)              | &nbsp;
[WSAEnumNameSpaceProvidersA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsanamespace.go#L17)    | &nbsp;
WSAEnumNameSpaceProvidersExA                                                                                           | &nbsp;
WSAEnumNameSpaceProvidersExW                                                                                           | &nbsp;
[WSAEnumNameSpaceProvidersW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsanamespace.go#L40)    | &nbsp;
[WSAEnumNetworkEvents](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L103)     | &nbsp;
[WSAEnumProtocolsA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L28)                  | &nbsp;
[WSAEnumProtocolsW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L68)                  | &nbsp;
[WSAEventSelect](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L98)            | &nbsp;
[WSAGetOverlappedResult](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L91)           | &nbsp;
[WSAHtonl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L19)                       | &nbsp;
[WSAHtons](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L28)                       | &nbsp;
[WSAIoctl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L117)                        | &nbsp;
WSAJoinLeaf                                                                                                            | &nbsp;
[WSALookupServiceBeginA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L60)             | &nbsp;
WSALookupServiceBeginW                                                                                                 | &nbsp;
[WSALookupServiceEnd](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L65)                | &nbsp;
[WSALookupServiceNextA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/btquery.go#L72)              | &nbsp;
WSANSPIoctl                                                                                                            | &nbsp;
[WSANtohl](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L37)                       | &nbsp;
[WSANtohs](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsahtonntoh.go#L46)                       | &nbsp;
WSAPoll                                                                                                                | &nbsp;
WSAProviderConfigChange                                                                                                | &nbsp;
WSARecv                                                                                                                | &nbsp;
WSARecvDisconnect                                                                                                      | &nbsp;
WSARecvFrom                                                                                                            | &nbsp;
WSARemoveServiceClass                                                                                                  | &nbsp;
[WSAResetEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L126)            | &nbsp;
WSASend                                                                                                                | &nbsp;
WSASendDisconnect                                                                                                      | &nbsp;
WSASendMsg                                                                                                             | &nbsp;
WSASendTo                                                                                                              | &nbsp;
[WSASetEvent](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L113)              | &nbsp;
WSASetLastError                                                                                                        | &nbsp;
WSASetServiceA                                                                                                         | &nbsp;
WSASetServiceW                                                                                                         | &nbsp;
[WSASocketA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L192)                      | &nbsp;
[WSASocketW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsasocket.go#L204)                      | &nbsp;
[WSAStartup](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_server/server.go#L31)                | &nbsp;
[WSAStringToAddressA](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L101)               | &nbsp;
[WSAStringToAddressW](https://github.com/K4rian/winsock2-go/blob/main/examples/snippets/wsaaddr.go#L137)               | &nbsp;
[WSAWaitForMultipleEvents](https://github.com/K4rian/winsock2-go/blob/main/examples/echo_tcp/v4_client/client.go#L118) | &nbsp;