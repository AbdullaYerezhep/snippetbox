package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) signupUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user signup form...")
}
func (h *Handler) signupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}
func (h *Handler) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}
