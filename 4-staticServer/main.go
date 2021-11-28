package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("assets"))))
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", "/resources/me_dog.png")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
