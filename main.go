package main

import (
	"fmt"
	"net/http"
)

const httpsPort = ":10443" // change to 443 as root on your actual server
const httpPort = ":8080"

func forceHTTPS() {
	http.ListenAndServe(httpPort, http.RedirectHandler("https://127.0.0.1"+httpsPort, http.StatusFound))
}

func serveSomething(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You've been served."))
}

func main() {
	go forceHTTPS()
	http.HandleFunc("/", serveSomething)
	fmt.Println("fart")
	// generate cert.pem and key.pem by running src/crypto/tls/generate_cert.go
	http.ListenAndServeTLS(httpsPort, "cert.pem", "key.pem", nil)
	fmt.Println("another fart")
}
