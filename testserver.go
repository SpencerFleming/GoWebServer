package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Makes new Mux Router, which wraps the webserver to make certain processes more
	// convenient.
	r := mux.NewRouter()

	// Home page template allows us to define our home page dynamically, server side.
	home_template := template.Must(template.ParseFiles("./dynamic/home.html"))

	// Comment section template that will allow use to dynamically fill our home page
	// with a comment section. Kept seperate from the home_template for ease of use.
	comments_template := template.Must(template.ParseFiles("./dynamic/comments.html"))

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

	// Comment section structs
	type comment struct {
		Name    string
		Message string
	}
	type comment_section struct {
		Title       string
		Description string
		Comments    []comment
	}

	// Initializing a comment struct so that it will remain as long as the website is functional.
	// We could utilize a binary file or database to store this data between runtimes if we wanted.
	comments := comment_section{
		Title: "Comments",
		Description: "Comments can be typed in and submitted here. They are processed by Go," +
			" and are completely immune to html injection attacks thanks to Go's handy html" +
			" libraries!",
		Comments: make([]comment, 0), //empty!
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

		if r.Method == http.MethodPost {
			comments.Comments = append(
				comments.Comments, comment{r.FormValue("Name"), r.FormValue("Message")})
			fmt.Fprintf(w, "<span class=success>Submitted!</span>")
		}

		comments_template.Execute(w, comments)
	})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("assets/"))))

	port := os.Getenv("PORT")
	// Default port
	if port == "" {
		port = "80"
	}
	port = ":" + port

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
