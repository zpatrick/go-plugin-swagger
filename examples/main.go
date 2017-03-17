package main

import (
	"github.com/zpatrick/go-plugin-swagger"
	"net/http"
)

func main() {
	swagger := swagger.New()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/apidocs.json", swagger)

	http.ListenAndServe(":9090", nil)
}
