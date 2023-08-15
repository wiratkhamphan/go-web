package views

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/login/login.html"))
	tmpl.Execute(w, nil)
}

func WD() {
	http.HandleFunc("/", Welcome)
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views/"))))
}
