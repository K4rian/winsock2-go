package snippets

import (
	"log"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- GetServByName: Retrieves service information corresponding to a service name and protocol.
- GetServByPort: Retrieves service information corresponding to a port and protocol.
*/

func ExampleGetServByName() {
	servName := append([]byte("http"), 0)
	servProto := append([]byte("tcp"), 0)
	serv, _ := ws2.GetServByName(&servName[0], &servProto[0])

	if serv == nil {
		log.Printf("ExampleGetServByName | The function returns nil for the service name '%s' and protocol name '%s'", servName, servProto)
		return
	}

	log.Printf("ExampleGetServByName | Name:  %s\n", ws2.BytePtrToString(serv.SName))
	log.Printf("ExampleGetServByName | Proto: %s\n", ws2.BytePtrToString(serv.SProto))
	log.Printf("ExampleGetServByName | Port:  %d\n", ws2.Htons(uint16(serv.SPort)))

	aliases := doubleBytePtrToStrSlice(serv.SAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("ExampleGetServByName | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("ExampleGetServByName | - %s\n", aliases[i])
		}
	}
}

func ExampleGetServByPort() {
	servPort := ws2.Htons(uint16(ws2.IPPORT_ECHO))
	servProto := append([]byte("udp"), 0)
	serv, _ := ws2.GetServByPort(uint16(servPort), &servProto[0])

	if serv == nil {
		log.Printf("ExampleGetServByPort | The function returns nil for the service port '%d' and protocol '%s'", servPort, servProto)
		return
	}

	log.Printf("ExampleGetServByPort | Name:  %s\n", ws2.BytePtrToString(serv.SName))
	log.Printf("ExampleGetServByPort | Proto: %s\n", ws2.BytePtrToString(serv.SProto))
	log.Printf("ExampleGetServByPort | Port:  %d\n", ws2.Htons(uint16(serv.SPort)))

	aliases := doubleBytePtrToStrSlice(serv.SAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("ExampleGetServByPort | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("ExampleGetServByPort | - %s\n", aliases[i])
		}
	}
}
