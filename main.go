package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

func main() {
	var dirPath = flag.StringP("dir", "d", "./", "Please select json files directory.")
	var port = flag.IntP("port", "p", 8080, "Please select if you want to use other port number except for 8080.")
	var noExt = flag.Bool("noext", false, "Please select if you need not to extenstion(.json).")
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
