package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/kmnkit/nomadcoin/explorer"
	"github.com/kmnkit/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n")
	fmt.Printf("Please use the following commands:\n")
	fmt.Printf("-port=4000:		Set the PORT of the Server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set port the server\n")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'\n")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
