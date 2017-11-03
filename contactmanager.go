package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Contact struct {
	Id        int
	Name      string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Notes     []byte
}

func (c *Contact) save() error {
	filename := c.Name + ".txt"
	return ioutil.WriteFile(filename, c.Notes, 0600)
}

func loadContact(name string) (*Contact, error) {
	filename := name + ".txt"
	notes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Contact{Name: name, Notes: notes}, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	
}

func viewContactHandler(w http.ResponseWriter, r *http.Request) {
	contact := r.URL.Path[len("/contact/"):]
	c, _ := loadContact(contact)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", c.Name, c.Notes)
}

func newContactHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	PORT := "8000"
	log.Print("Running server on port " + PORT)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact/", viewContactHandler)
	http.HandleFunc("/new", newContactHandler)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
