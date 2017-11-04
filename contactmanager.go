package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	PORT        = "8000"
	DB_USER     = "ubuntu"
	DB_PASSWORD = "ubuntu"
	DB_NAME     = "contactdb"
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
	
	var contact Contact

	sqlStatement := `SELECT id, first_name, last_name, email, phone, notes FROM contacts WHERE id=$1;`
	row := db.QueryRow(sqlStatement, id)  
	err := row.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone, &contact.Notes)

	return &contact, err
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", 301)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Get all contacts with a db query

	contacts := []Contact{Contact{1, "af", "al", "a", "a", ""}, Contact{2, "bf", "bl", "", "", "b"}}

	// Show contact list
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/home.tmpl")
	t.Execute(w, contacts)
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

	updatedContact := Contact{
		Id:        contactId,
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		Notes:     r.FormValue("notes"),
	}

	sqlStatement := `UPDATE contacts SET first_name = $2, last_name = $3, email = $4, phone = $5, notes = $6 WHERE id = $1 RETURNING id;`
	updatedId := 0
	err := db.QueryRow(sqlStatement, updatedContact.Id, updatedContact.FirstName, updatedContact.LastName, updatedContact.Email, updatedContact.Phone, updatedContact.Notes).Scan(&updatedId)
	if err != nil {
		panic(err)
	}
	if updatedContact.Id != updatedId {
		log.Print("Big problem!")
	}

	t.Execute(w, updatedContact)
	log.Print("Displaying updated contact page for contact " + contactIdString + ".")
}

func newContactHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.tmpl", "templates/add_contact.tmpl")

	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		log.Print("Displaying new contact page.")
		return
	}

	newContact := Contact{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		Notes:     r.FormValue("notes"),
	}

	sqlStatement := `INSERT INTO contacts (first_name, last_name, email, phone, notes) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	newId := 0
	err := db.QueryRow(sqlStatement, newContact.FirstName, newContact.LastName, newContact.Email, newContact.Phone, newContact.Notes).Scan(&newId)
	if err != nil {
		panic(err)
	}

	log.Print("Saved new contact with ID " + strconv.Itoa(newId) + ".")
	t.Execute(w, newContact)
}

func main() {
	r := mux.NewRouter()

	dbinfo := "user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"
	db, _ = sql.Open("postgres", dbinfo)
	defer db.Close()

	err := db.Ping()  
	if err != nil {  
		panic(err)
	}

	r.HandleFunc("/", redirectHandler)
	r.HandleFunc("/contacts", homeHandler)
	r.HandleFunc("/contacts/{contactid:[0-9]+}", viewContactHandler)
	r.HandleFunc("/contacts/{contactid:[0-9]+}/edit", editContactHandler)
	r.HandleFunc("/contacts/new", newContactHandler)

	log.Print("Running server on port " + PORT + ".")
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
