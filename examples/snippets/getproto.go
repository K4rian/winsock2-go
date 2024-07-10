package snippets

import (
	"log"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following function:
- GetProtoByName:   Retrieves the protocol information corresponding to a protocol name.
- GetProtoByNumber: Retrieves protocol information corresponding to a protocol number.
*/

func ExampleGetProtoByName() {
	protoName := append([]byte("RDP"), 0)
	proto, _ := ws2.GetProtoByName(&protoName[0])

	if proto == nil {
		log.Printf("GetProtoByName | The function returns nil for the protocol name '%s'", protoName)
		return
	}
	log.Printf("GetProtoByName | Name: %s\n", protoName)
	log.Printf("GetProtoByName | Number: %d\n", proto.PProto)

	aliases := doubleBytePtrToStrSlice(proto.PAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("GetProtoByName | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("GetProtoByName | - %s\n", aliases[i])
		}
	}
}

func ExampleGetProtoByNumber() {
	protoNB := int32(ws2.IPPROTO_ESP)
	proto, _ := ws2.GetProtoByNumber(protoNB)

	if proto == nil {
		log.Printf("GetProtoByNumber | The function returns nil for the protocol number '%d'", protoNB)
		return
	}

	protoName := ws2.BytePtrToString(proto.PName)
	if len(protoName) > 0 {
		log.Printf("GetProtoByNumber | Name: %s\n", protoName)
	}
	log.Printf("GetProtoByNumber | Number: %d\n", proto.PProto)

	aliases := doubleBytePtrToStrSlice(proto.PAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("GetProtoByNumber | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("GetProtoByNumber | - %s\n", aliases[i])
		}
	}
}
