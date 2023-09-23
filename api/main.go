package main

import (
	"net/http"
)

// simple http server returning a json message returning string param received on url

func main() {
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<p style="font-size:55px;">message: Hello, ` + r.URL.Query().Get("name") + `</br>status: ok</br>Want to learn more about API and coding? Visit our website: <a href="https://scula.io">https://scula.io</a></p>`))
	// w.Write([]byte(`{"message": "Hello, ` + r.URL.Query().Get("name") + `", "status": "ok", "Want to learn more about API and coding? Visit our website": "https://scula.io"}`))
}
