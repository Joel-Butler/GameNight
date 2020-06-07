// +build ignore
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//todo - load stuff into the template...
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, p)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
