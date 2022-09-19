package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	listen := flag.String("listen", "localhost:1337", "listen address")
	dir := flag.String("dir", "public", "directory to serve")
	flag.Parse()

	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
