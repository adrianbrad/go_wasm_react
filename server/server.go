package main

import (
	"github.com/rs/cors"
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){http.ServeFile(w, r, "../go/main.wasm")})
	handler := cors.Default().Handler(mux)
	log.Printf("Serving main.wasm on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, handler))
}
