package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"log"
)

func form(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		form, err := ioutil.ReadFile("./form.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(form))
	case "POST":
		// Parse the form
		r.ParseForm()
		// Print the values of the form
		fmt.Fprintf(w, "You said you like: %s and you picked #%s\n", r.FormValue("fruit"),
                                                                     r.FormValue("number"))
	}
}

func main() {
	http.HandleFunc("/form", form)
	http.ListenAndServe(":8080", nil)
}


