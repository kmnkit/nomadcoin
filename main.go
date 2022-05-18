package main

import (
	"github.com/kmnkit/nomadcoin/cli"
	"github.com/kmnkit/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
