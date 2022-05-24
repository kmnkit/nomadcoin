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
	address    string // 16진수
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

func restoreKey() (key *ecdsa.PrivateKey) { // named return을 하면 미리 초기화가 돼서 return을 안 해도 됨
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return
}

func aFromK(key *ecdsa.PrivateKey) string {
	//
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			// yes -> restore from file
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		// no -> create private key, save to file
		w.address = aFromK(w.privateKey)
	}
	return w
}
