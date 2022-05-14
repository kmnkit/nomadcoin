package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash) // Sprintf는 프린트를 해주지는 않지만 String으로 반환해준다.
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
