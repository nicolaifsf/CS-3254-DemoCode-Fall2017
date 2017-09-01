package main

import (
	"net/http"
	"log"
	"time"
	"io/ioutil"
	"fmt"
)

func favoriteCookie(w http.ResponseWriter, r *http.Request) {
	// The name we will be giving our cookie
	const COOKIE_NAME = "favoritecookie"
	switch r.Method {
	// If it is a GET request, we check if the cookie already has been set. If not, we provide a form for them to
	// tell us their favorite cookie
	case http.MethodGet:
		// Try to retrieve the cookie
		favoriteCookie, err := r.Cookie(COOKIE_NAME)
		if err != nil {
			log.Println(err)
		}
		// If the cookie does not exist, we provide a form for the user to let us know their favorite cookie
		if favoriteCookie == nil {
			form, err := ioutil.ReadFile("form.html")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, string(form))
		} else {
			// If the cookie does exist, we simply reply, letting them know what they told us was their favorite cookie
			fmt.Fprintf(w, "You have already told us your favorite cookie is: %s", favoriteCookie.Value)
		}
	case http.MethodPost:
		// First we have to parse the form
		r.ParseForm()
		cookie := http.Cookie {
			Name: 		COOKIE_NAME,								// The name of the cookie
			Value: 		r.PostFormValue("FavoriteCookie"),		// Their favorite cookie
			Expires: 	time.Now().Add(1 * time.Hour),				// Set the expiration date of the cookie
		}
		// Set the cookie
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, "We have logged your choice for favorite cookie!")
	}
}

func main() {
	http.HandleFunc("/cookie", favoriteCookie)
	http.ListenAndServe(":8080", nil)
}


