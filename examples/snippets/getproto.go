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
		log.Printf("ExampleGetProtoByName | The function returns nil for the protocol name '%s'", protoName)
		return
	}
	log.Printf("ExampleGetProtoByName | Name: %s\n", protoName)
	log.Printf("ExampleGetProtoByName | Number: %d\n", proto.PProto)

	aliases := doubleBytePtrToStrSlice(proto.PAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("ExampleGetProtoByName | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("ExampleGetProtoByName | - %s\n", aliases[i])
		}
	}
}

func ExampleGetProtoByNumber() {
	protoNB := int32(ws2.IPPROTO_ESP)
	proto, _ := ws2.GetProtoByNumber(protoNB)

	if proto == nil {
		log.Printf("ExampleGetProtoByNumber | The function returns nil for the protocol number '%d'", protoNB)
		return
	}

	protoName := ws2.BytePtrToString(proto.PName)
	if len(protoName) > 0 {
		log.Printf("ExampleGetProtoByNumber | Name: %s\n", protoName)
	}
	log.Printf("ExampleGetProtoByNumber | Number: %d\n", proto.PProto)

	aliases := doubleBytePtrToStrSlice(proto.PAliases) // _common.go#doubleBytePtrToStrSlice
	if len(aliases) > 0 {
		log.Printf("ExampleGetProtoByNumber | Aliases (%d):\n", len(aliases))
		for i := 0; i < len(aliases); i++ {
			log.Printf("ExampleGetProtoByNumber | - %s\n", aliases[i])
		}
	}
}
