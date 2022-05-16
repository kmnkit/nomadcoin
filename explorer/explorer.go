package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/kmnkit/nomadcoin/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {	
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)	
}

func add(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

}

// Start 탐색시작
func Start(){
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))	
}