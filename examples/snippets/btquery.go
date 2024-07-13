package snippets

import (
	"log"
	"unsafe"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- Socket:                 Creates a socket that is bound to a specific transport service provider.
- CloseSocket:            Closes an existing socket.
- GetSockOpt:             Retrieves a socket option.
- WSALookupServiceBeginA: Initiates a client query that is constrained by the information contained within a WSAQuerySetA structure.
- WSALookupServiceEnd:    Called to free the handle after previous calls to WSALookupServiceBegin and WSALookupServiceNext.
- WSALookupServiceNextA:  Called after obtaining a handle from a previous call to WSALookupServiceBegin in order to retrieve the requested service information.
- WSAAddressToStringA:    Converts all components of a sockaddr structure into a human-readable string representation of the address.
*/

/*
 - WSAStartup() has to be called before using the example function.
 - A Bluetooth device is required.
*/

func ExampleBTQuery() {
	const (
		BTHPROTO_RFCOMM = 3
		BTH_SDP_CONNECT = 0x0001001C
	)

	//
	// Create a Bluetooth socket
	socket, err := ws2.Socket(ws2.AF_BTH, ws2.SOCK_STREAM, BTHPROTO_RFCOMM)
	if socket == ws2.INVALID_SOCKET {
		log.Printf("ExampleBTQuery | Socket creation error: %v\n", err)
		return
	}
	defer ws2.CloseSocket(socket)

	//
	// Get Bluetooth protocol info
	protoInfo := ws2.WSAProtocolInfoA{}
	protoInfoSize := int32(unsafe.Sizeof(protoInfo))

	ret, err := ws2.GetSockOpt(socket, ws2.SOL_SOCKET, ws2.SO_PROTOCOL_INFOA, unsafe.Pointer(&protoInfo), &protoInfoSize)
	if ret != 0 {
		log.Printf("ExampleBTQuery | GetSockOpt error (%d): %v\n", ret, err)
		return
	}

	//
	// Initiate a query to the namespace service to obtain a handle
	qa := ws2.WSAQuerySetA{}
	qa.Size = uint32(unsafe.Sizeof(qa))
	qa.NameSpace = ws2.NS_BTH
	hLookup := ws2.HANDLE(0)
	controlFlags := uint32(0x0010 | 0x0002 | 0x0100 | 0x2000 | 0x0020 | 0x0200 | 0x8000)

	ret, err = ws2.WSALookupServiceBeginA(&qa, controlFlags, &hLookup)
	if ret != 0 {
		log.Printf("ExampleBTQuery | WSALookupServiceBeginA error (%d): %v\n", ret, err)
		return
	}
	defer ws2.WSALookupServiceEnd(hLookup)

	//
	// Get the service info
	qResult := ws2.WSAQuerySetA{}
	qBufferLen := uint32(unsafe.Sizeof(qResult))

	ret, err = ws2.WSALookupServiceNextA(hLookup, controlFlags, &qBufferLen, &qResult)
	if ret != 0 {
		log.Printf("ExampleBTQuery | WSALookupServiceNextA error (%d): %v\n", ret, err)
		return
	}
	log.Printf("ExampleBTQuery | Bluetooth Instance Name: %s\n", ws2.BytePtrToString(qResult.ServiceInstanceName))

	//
	// Print the local Bluetooth device address
	csAddrInfo := qResult.CSABuffer
	// pDeviceInfo := qResult.LpBlob // BTH_DEVICE_INFO
	devAddrLen := uint32(256)
	devAddrBuf := make([]byte, devAddrLen)

	ret, err = ws2.WSAAddressToStringA(
		csAddrInfo.LocalAddr.Sockaddr,
		uint32(csAddrInfo.LocalAddr.SockaddrLength),
		&protoInfo,
		&devAddrBuf[0],
		&devAddrLen,
	)
	if ret != 0 {
		log.Printf("ExampleBTQuery | WSAAddressToStringA (1) error (%d): %v\n", ret, err)
		return
	}
	log.Printf("ExampleBTQuery | Bluetooth Local Device Address:  %s\n", string(devAddrBuf))

	//
	// Print the remote Bluetooth device address
	remoteAddrLen := uint32(unsafe.Sizeof(devAddrBuf))

	ret, err = ws2.WSAAddressToStringA(
		csAddrInfo.RemoteAddr.Sockaddr,
		uint32(csAddrInfo.RemoteAddr.SockaddrLength),
		&protoInfo,
		&devAddrBuf[0],
		&remoteAddrLen,
	)
	if ret != 0 {
		log.Printf("ExampleBTQuery | WSAAddressToStringA (2) error (%d): %v\n", ret, err)
		return
	}
	log.Printf("ExampleBTQuery | Bluetooth Remote Device Address: %s\n", string(devAddrBuf))

	//
	// WSALookupServiceNextA can be called again until it returns WSA_E_NO_MORE
	// ...
}
