package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/kmnkit/nomadcoin/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

const (
	fileName string = "nomadcoin.wallet"
)

var w *wallet

func hasWalletFile() bool {
	// has a wallet already?
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	// key를 받아서 복붙 가능한 형태로 변환함.
	// 그 byte들을 파일로 저장.
	// 이 key를 전에 한 것처럼 16진수 문자열로 만들 필요는 없음
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)

}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			// yes -> restore from file
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		// no -> create private key, save to file
	}
	return w
}
