package snippets

import (
	"log"

	"github.com/K4rian/winsock2-go/ws2"
)

/*
Demonstrate the usage of the following functions:
- WSAEnumNameSpaceProvidersA: Retrieves information on available namespace providers.
- WSAEnumNameSpaceProvidersW: Retrieves information on available namespace providers.
*/

// WSAStartup() has to be called before using these functions.

func ExampleWSAEnumNameSpaceProvidersA() {
	ns := make([]ws2.WSANameSpaceInfoA, 20)
	bufLen := uint32(4096)

	spaces, err := ws2.WSAEnumNameSpaceProvidersA(&bufLen, &ns[0])
	if spaces <= 0 {
		log.Printf("ExampleWSAEnumNameSpaceProvidersA | Error: %v\n", err)
		return
	}
	log.Printf("ExampleWSAEnumNameSpaceProvidersA | %d namespaces available\n", spaces)

	for i := 0; i < int(spaces); i++ {
		e := ns[i]

		log.Printf("\n")
		log.Printf("- NSProviderId: %v\n", ws2.GUIDToString(e.NSProviderId))
		log.Printf("- Namespace:    %v\n", e.NameSpace)
		log.Printf("- Active:       %v\n", e.Active)
		log.Printf("- Version:      %v\n", e.Version)
		log.Printf("- Identifier:   %s\n", ws2.BytePtrToString(e.Identifier))
	}
}

func ExampleWSAEnumNameSpaceProvidersW() {
	ns := make([]ws2.WSANameSpaceInfoW, 20)
	bufLen := uint32(4096)

	spaces, err := ws2.WSAEnumNameSpaceProvidersW(&bufLen, &ns[0])
	if spaces <= 0 {
		log.Printf("ExampleWSAEnumNameSpaceProvidersW | Error: %v\n", err)
		return
	}
	log.Printf("ExampleWSAEnumNameSpaceProvidersW | %d namespaces available\n", spaces)

	for i := 0; i < int(spaces); i++ {
		e := ns[i]
		id, _ := ws2.UTF16PtrToString(e.Identifier)

		log.Printf("\n")
		log.Printf("- NSProviderId: %v\n", ws2.GUIDToString(e.NSProviderId))
		log.Printf("- Namespace:    %v\n", e.NameSpace)
		log.Printf("- Active:       %v\n", e.Active)
		log.Printf("- Version:      %v\n", e.Version)
		log.Printf("- Identifier:   %s\n", id)
	}
}
