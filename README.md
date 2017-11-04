This is a contact manager implemented using Go.

## Running

Simply `vagrant up` to start the server. Visit `localhost:8000` to view the website.

To view the server log, `vagrant ssh` into the virtual machine, `sudo su` to switch to the root user, and `tmux attach -t server` to open the session that the log is writing to.

## Bugs/fixes

* Still need to implement proper 404 pages if people visit /contact/non-existent-id.
* Definition list CSS should be cleaned up, navigation should be made nicer, and the contact list could look better.
* Some SQL errors still need to be properly handled in the code.
* Templates have a bit of repetition that could be reduced.
* Required fields could be added.
* Notes field should allow new lines.
* Giant submit button looks kinda juvenile.
* Postgres login should be more secure.
* CSRF field needs to be added: github.com/gorilla/csrf
