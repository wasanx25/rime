package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	var jsonPath = flag.String("path", "", "select json file path")
	var port = flag.Int("port", 8080, "select port number")
	flag.Parse()

	if len(*jsonPath) == 0 {
		log.Fatal("You must set json file path.")
	}

	bytes, err := ioutil.ReadFile(*jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	urlPath := filepath.Join("/", filepath.Base(*jsonPath))

	log.Printf("http://localhost:%d\n", *port)
	log.Printf("See browse -> http://localhost:%d%s", *port, urlPath)
	http.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(bytes))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
