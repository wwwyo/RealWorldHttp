package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "server.crt", "server.key", nil)
	log.Println(err)
}

// ルート認証局の証明書付でリクエストすればtls通信できる
// curl --cacert ca.crt  https://localhost:18443
