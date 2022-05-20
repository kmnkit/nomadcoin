package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/kmnkit/nomadcoin/utils"
)

const (
	signature     string = "03756bdfe86de3586e9bfc64e8d6eb45f68c2d808a0f830f09515e4556fadfa5a804bf57d8df9226a021d6147de7469755eac0dd604403ce8031e08f2309effd%"
	privateKey    string = "3077020101042093124b5fc5d568e73df511f96c03a854d07b6e86888165af542e2a4384040930a00a06082a8648ce3d030107a14403420004021b5921003c788a0c208908ff029f9c76dc0cff89c5b58dce39d0c9fd1ca71b6abc8dc73e0f465b4c280d6f9d68ff3b138b31a95d4106cfd430ec18d51efe65"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	// 비공개키 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	fmt.Printf("%x\n\n\n", keyAsBytes)
	utils.HandleErr(err)

	// 메세지 바이트화
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 메세지에 서명함
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x", signature)
	utils.HandleErr(err)
}
