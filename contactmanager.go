package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id        int
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
	//Find the contact with that id in the database and load it
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

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", 301)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Get all contacts with a db query

	// Show contact list
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/home.tmpl")
	t.Execute(w, nil)
}

func viewContactHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactIdString := vars["contactid"]
	contactId, _ := strconv.Atoi(contactIdString)
	contact, _ := loadContact(contactId)
	
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/contact.tmpl")
	t.Execute(w, contact)
	
	log.Print("Displaying contact " + contactIdString + ".")
}

func editContactHandler(w http.ResponseWriter, r *http.Request) {
	// Edit a contact
	vars := mux.Vars(r)
	contactIdString := vars["contactid"]
	contactId, _ := strconv.Atoi(contactIdString)
	contact, _ := loadContact(contactId)
	
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/edit_contact.tmpl")
	
	if r.Method != http.MethodPost {
		t.Execute(w, contact)
		log.Print("Displaying edit contact page for contact " + contactIdString + ".")
		return
	}
}

func newContactHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/add_contact.tmpl")

	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		log.Print("Displaying new contact page.")
		return
	}
	
	// Create a new contact
}

func main() {
	PORT := "8000"
	r := mux.NewRouter()

	r.HandleFunc("/", redirectHandler)
	r.HandleFunc("/contacts", homeHandler)
	r.HandleFunc("/contacts/{contactid:[0-9]+}", viewContactHandler)
	r.HandleFunc("/contacts/{contactid:[0-9]+}/edit", editContactHandler)
	r.HandleFunc("/contacts/new", newContactHandler)

	log.Print("Running server on port " + PORT + ".")
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
