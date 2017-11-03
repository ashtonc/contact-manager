package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Contact struct {
	Id        int
	Name      string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Notes     string
}

func (c *Contact) save() error {
	//Save the contact in the database
	return nil
}

func loadContact(id int) (*Contact, error) {
	//Find the contact with that Id in the database and load it
	contact := Contact{
		Id:        1,
		FirstName: "Ashton",
		LastName:  "Charbonneau",
		Email:     "ashton@ashtonc.ca",
		Phone:     "911",
		Notes:     "",
	}
	return &contact, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	// Show contact list
	w.Write([]byte(message))
}

func viewContactHandler(w http.ResponseWriter, r *http.Request) {
	contactIdString := r.URL.Path[len("/contact/"):]

	contactId, err := strconv.Atoi(contactIdString)
	if err != nil {
		log.Print("Contact " + contactIdString + " not found")
		fmt.Fprintf(w, "Not found.")
		return
	}

	// Load and print the contact page
	contact, _ := loadContact(contactId)

	t, err := template.ParseFiles("templates/base.tmpl", "templates/contact.tmpl")
	t.Execute(w, contact)

	log.Print("Displaying contact " + contactIdString)
}

func newContactHandler(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

func main() {
	PORT := "8000"

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact/", viewContactHandler)
	http.HandleFunc("/contact/new", newContactHandler)

	log.Print("Running server on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
