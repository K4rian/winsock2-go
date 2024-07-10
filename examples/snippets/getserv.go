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
		log.Printf("GetServByName | The function returns nil for the service name '%s' and protocol name '%s'", servName, servProto)
		return
	}

	log.Printf("GetServByName | Name:  %s\n", ws2.BytePtrToString(serv.SName))
	log.Printf("GetServByName | Proto: %s\n", ws2.BytePtrToString(serv.SProto))
	log.Printf("GetServByName | Port:  %d\n", ws2.Htons(uint16(serv.SPort)))

	aliases := doubleBytePtrToStrSlice(serv.SAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("GetServByName | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("GetServByName | - %s\n", aliases[i])
		}
	}
}

func ExampleGetServByPort() {
	servPort := ws2.Htons(uint16(ws2.IPPORT_ECHO))
	servProto := append([]byte("udp"), 0)
	serv, _ := ws2.GetServByPort(uint16(servPort), &servProto[0])

	if serv == nil {
		log.Printf("GetServByPort | The function returns nil for the service port '%d' and protocol '%s'", servPort, servProto)
		return
	}

	log.Printf("GetServByPort | Name:  %s\n", ws2.BytePtrToString(serv.SName))
	log.Printf("GetServByPort | Proto: %s\n", ws2.BytePtrToString(serv.SProto))
	log.Printf("GetServByPort | Port:  %d\n", ws2.Htons(uint16(serv.SPort)))

	aliases := doubleBytePtrToStrSlice(serv.SAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("GetServByPort | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("GetServByPort | - %s\n", aliases[i])
		}
	}
}
