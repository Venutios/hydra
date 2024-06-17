package root

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/root/root.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
