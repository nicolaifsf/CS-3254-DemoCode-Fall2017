package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func form(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		form, err := ioutil.ReadFile("./form.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(form))
	case http.MethodPost:
		// Example of redirecting after the form has been filled out
		http.Redirect(w, r, "/thanks", http.StatusTemporaryRedirect)
	}
}

func thanks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thanks for all the fish!")
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/thanks", thanks)
	http.ListenAndServe(":8080", nil)
}


