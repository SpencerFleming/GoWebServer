package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Makes new Mux Router, which wraps the webserver to make certain processes more
	// convenient.
	r := mux.NewRouter()

	// Home page template allows us to define our home page dynamically, server side.
	home_template := template.Must(template.ParseFiles("./assets/html/home.html"))

	// Structs are used to pass relevant data to our templates. Here we use two so we
	// can have a slice of home_data_examples, instead of a static number of them.
	type home_data_examples struct {
		Title string
		URL   string
	}
	type home_data struct {
		Title       string
		Description string
		Examples    []home_data_examples
	}

	// The handle function defines how we handle requests to the root of our website.
	// It can be set to handle other pages too, but right now it's set to handle "/",
	// per the first argument.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := home_data{
			Title: "Go webserver is running!",
			Description: "This page is dynamically generated when you request data from the " +
				"root directory of the webserver. Below are some static webpages composed of " +
				"HTML and css to show how compatible Go's server is with regular static webpages.",
			Examples: []home_data_examples{{Title: "Gallery",
				URL: "/static/html/gallery.html"},
				{Title: "Imbedded page",
					URL: "/static/html/embedded.html"},
				{Title: "JQuery example",
					URL: "/static/html/funbutton.html"}},
		}
		home_template.Execute(w, data)
	})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("assets/"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
