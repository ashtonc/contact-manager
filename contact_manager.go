package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contact struct {
	Name string
	//    FirstName string
	//    LastName string
	//    Email string
	//    Phone string
	Notes []byte
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	contact := r.URL.Path[len("/contact/"):]
	c, _ := loadContact(contact)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", c.Name, c.Notes)
}

func main() {
	http.HandleFunc("/contact/", viewHandler)
	http.ListenAndServe(":8000", nil)
}
