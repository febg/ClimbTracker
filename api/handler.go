package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PostRegisterUser gets new user information from HTTP Post request and registers user in the main user Database
func (c *Control) PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uEmail := v["user_email"]
	uPassword := v["user_password"]
	uName := v["user_name"]

	if uEmail == "" || uPassword == "" || uName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: No User Id was given")
		return
	}
	fmt.Println("Vars: ", v)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"user_id\":\"%s\" \n \"user:password\":\"%s\" \n \"user_name\":\"%s\"}", uEmail, uPassword, uName)
	return
}

// PostLogInUser gets clients credentials from HTTP pPost request, compares it with existing credentials on database and grants or denies access
func (c *Control) PostLogInUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("[LOG] Login in User")
}
