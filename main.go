package main

import (
	"github.com/kmnkit/nomadcoin/explorer"
	"github.com/kmnkit/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
