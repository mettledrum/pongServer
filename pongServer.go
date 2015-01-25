package main

import (
		"net/http"
		"log"
)

func angularHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

// loosen up access protocol?
func changeHeaderThenServe(h http.Handler) http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) { 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r) 
	}
}

func main() {

	// TODO: make these endpoints
	// angularJS home site -- weird with how file is served :(
	// most current photo
	// POST site to make custom gif from selection and let you download it

	http.HandleFunc("/", angularHandler)
	http.Handle("/pong-pics/", http.StripPrefix("/pong-pics/", changeHeaderThenServe(http.FileServer(http.Dir("./daily_pictures")))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

