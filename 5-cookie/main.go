package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", increment)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func increment(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("count")
	var newValue string

	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "count",
			Value: "1",
		})
		newValue = "1"
	} else {
		prevValue, err := strconv.Atoi(cookie.Value)
		if err != nil {
			log.Fatalln(err)
		}
		newValue = strconv.Itoa(prevValue + 1)

		cookie.Value = newValue
		http.SetCookie(w, cookie)
	}
	fmt.Fprintln(w, "Your browser has visited this domain this many times: "+newValue)
}
