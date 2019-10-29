package login

import (
	"net/http"
	"sample/app/shared/handler"

	// "github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Email   string
	Message string
}

// Login page
func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {

	session, _ := handler.Store.Get(r, "user-session")

	if (!session.IsNew) {
		http.Redirect(w, r, "/profile", 302)

	} else {
		err := h.ResponseHTML(w, r, "login/login", templateData{})

		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	}
}

// LoginHandle page
func (h *HTTPHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	if !handler.CheckLoginUser(email, password) {

		message := ""

		if !handler.CheckEmptyEmail(email) {
			message = "Email must be not empty!"
		} else if !handler.CheckPassword(password) {
			message = "Password must be > 3 characters!"
		} else if !handler.CheckExistedUser(email, password) {
			message = "Email is not existed or wrong password!"
		}

		err := h.ResponseHTML(w, r, "login/login", templateData{
			Message: message,
		})

		if err != nil {
			_ = h.StatusServerError(w, r)
		}

	} else {
		session, _ := handler.Store.Get(r, "user-session")
		session.Values["email"] = email
		session.Values["password"] = password
		session.Save(r, w)
		http.Redirect(w, r, "/profile", 302)
	}
}

// NewLoginHTTPHandler responses new HTTPHandler instance.
func NewLoginHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
