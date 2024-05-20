package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Lifts struct to hold information about a film
type Exercise struct {
	Movement string
	Weight   int
	Sets     int
	Reps     int
}

func main() {
	fmt.Println("Go App is starting...")

	// handler function #1 - serves the main page with a list of exercises
	h1 := func(w http.ResponseWriter, r *http.Request) {
		// Parse the index.html template file
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Sample data of films to be displayed
		exercise := map[string][]Exercise{
			"Exercise": {
				{Movement: "Bench Press", Weight: 80, Sets: 1, Reps: 10},
				{Movement: "Deadlifts", Weight: 100, Sets: 1, Reps: 7},
				{Movement: "Squats", Weight: 100, Sets: 1, Reps: 8},
			},
		}

		// Execute the template with the film data
		tmpl.Execute(w, exercise)
	}

	// // handler function #2 - handles adding a new film and returns a template block as an HTMX response
	// h2 := func(w http.ResponseWriter, r *http.Request) {
	// 	// Simulate processing time with a sleep
	// 	time.Sleep(1 * time.Second)

	// 	// Retrieve the title and director from the POST request form
	// 	title := r.PostFormValue("title")
	// 	director := r.PostFormValue("director")

	// 	// Prepare a new template for the film list element
	// 	tmpl := template.Must(template.ParseFiles("index.html"))

	// 	// Execute the template for the specific block "film-list-element" with the new film data
	// 	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	// }

	// Register the handlers for the root and add-film paths
	http.HandleFunc("/", h1)
	// http.HandleFunc("/add-film/", h2)

	// Start the HTTP server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
