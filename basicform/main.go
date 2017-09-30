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
		// Parse the form
		r.ParseForm()
		// Print the values of the form
		fmt.Fprintf(w, "You said you like: %s and you picked #%s\n", r.PostFormValue("fruit"),
																			r.PostFormValue("number"))
	}
}

func swag(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"swag")
}

func main() {
	http.HandleFunc("/form", form)
    http.HandleFunc("/swag", swag)
	http.ListenAndServe(":8080", nil)
}


