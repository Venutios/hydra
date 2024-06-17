package hydra

import (
	"log"
	"net/http"

	root "github.com/Venutios/hydra/internal/hydra/routes/root"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root.RootHandler)
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/"))))
	http.Handle("/", mux)
	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
