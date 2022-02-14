package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"path"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "",err
	}

	return string(bytes), nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", path.Base(r.URL.EscapedPath()))
	var html, err = loadFile("Desktop/test/" + path.Base(r.URL.EscapedPath()))
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404, Sorry can't find the page")
	}
	fmt.Fprint(w, html)
	//fmt.Fprintf(w, "Hello %s, welcome to your HTTP server", r.URL.Path[1:])
} 


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

}