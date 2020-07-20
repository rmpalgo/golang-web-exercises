package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>" +
		"<ul>" +
		"<li>This is cool</li>" +
		"<li>This is really cool</li>" +
		"<li>Amazing</li>" +
		"<li>Woot</li>" +
		"<li>Huehuheuhe</li>" +
		"</ul>" )
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div><em>GoLand</em>!</div>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
