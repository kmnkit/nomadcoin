package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/kmnkit/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml")) // 이 안에서 에러가 있을 경우 자동으로 처리해 줌
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
