package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	var dirPath = flag.String("dir", "./", "select json files directory")
	var port = flag.Int("port", 8080, "select port number")
	var noExt = flag.Bool("ext", false, "select yes or no extension flag")
	flag.Parse()

	if len(*dirPath) == 0 {
		log.Fatal("You must set json file path.")
	}

	jsonFiles := getFiles(*dirPath)

	log.Printf("http://localhost:%d\n", *port)
	var urlPath string
	for _, file := range jsonFiles {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		if *noExt {
			count := len(file[:len(file)-len(filepath.Ext(file))])
			file = file[0:count]
		}
		urlPath = filepath.Join("/", file)
		log.Printf("See browse -> http://localhost:%d%s", *port, urlPath)

		http.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, string(bytes))
		})
	}

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func getFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getFiles(filepath.Join(dir, file.Name()))...)
			continue
		}
		if path.Ext(file.Name()) == ".json" {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}
	return paths
}
