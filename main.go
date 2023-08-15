package main

import (
	"net/http"
	_ "po/config"
	"po/views"
)

func main() {
	// config.BD()
	// http.HandleFunc("/", views.Welcome) // Use Welcome instead of WD
	// http.Handle("/views/style/", http.StripPrefix("/views/style/", http.FileServer(http.Dir("views/style"))))
	views.WD()
	http.ListenAndServe(":8000", nil)
}
