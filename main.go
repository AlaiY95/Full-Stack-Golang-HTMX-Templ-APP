package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Lifts struct to hold information about a Exercise
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

		// Sample data of exercise to be displayed
		exercise := map[string][]Exercise{
			"Exercise": {
				{Movement: "Bench Press", Weight: 80, Sets: 1, Reps: 10},
				{Movement: "Deadlifts", Weight: 100, Sets: 1, Reps: 7},
				{Movement: "Squats", Weight: 100, Sets: 1, Reps: 8},
			},
		}

		// Execute the template with the exercise data
		tmpl.Execute(w, exercise)
	}

	// handler function #2 - handles adding a new exercise and returns a template block as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		// Simulate processing time with a sleep
		time.Sleep(1 * time.Second)

		// Retrieve the fields from the POST request form
		movement := r.PostFormValue("movement")
		weightStr := r.PostFormValue("weight")
		setsStr := r.PostFormValue("sets")
		repsStr := r.PostFormValue("reps")

		// Convert the form values from string to int
		weight, err := strconv.Atoi(weightStr)
		if err != nil {
			http.Error(w, "Invalid weight value", http.StatusBadRequest)
			return
		}
		sets, err := strconv.Atoi(setsStr)
		if err != nil {
			http.Error(w, "Invalid sets value", http.StatusBadRequest)
			return
		}
		reps, err := strconv.Atoi(repsStr)
		if err != nil {
			http.Error(w, "Invalid reps value", http.StatusBadRequest)
			return
		}

		// Prepare a new template for the exercise list element
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Execute the template for the specific block "exercise-list-element" with the new exercise data
		tmpl.ExecuteTemplate(w, "exercise-list-element", Exercise{Movement: movement, Weight: weight, Sets: sets, Reps: reps})
	}

	// Register the handlers for the root and add-exercise paths
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-exercise/", h2)

	// Start the HTTP server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
