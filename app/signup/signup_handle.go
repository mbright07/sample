package signup

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/schema"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Message string
}

// Signup page
func (h *HTTPHandler) Signup(w http.ResponseWriter, r *http.Request) {

	err := h.ResponseHTML(w, r, "signup/signup", templateData{})

	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// SignupHandle page
func (h *HTTPHandler) SignupHandle(w http.ResponseWriter, r *http.Request) {

	// email := r.FormValue("email")
	// password := r.FormValue("password")

	user := handler.User{}

	r.ParseForm()

	decoder := schema.NewDecoder()
	decoder.Decode(&user, r.PostForm)

	if !handler.CheckSignupUser(user.Email, user.Password) {

		message := ""

		if !handler.CheckEmptyEmail(user.Email) {
			message = "Email must be not empty!"
		} else if !handler.CheckPassword(user.Password) {
			message = "Password must be > 3 characters!"
		} else if !handler.CheckExistedUser(user.Email, user.Password) {
			message = "Email is existed!"
		}

		err := h.ResponseHTML(w, r, "signup/signup", templateData {
			Message: message,
		})

		if err != nil {
			_ = h.StatusServerError(w, r)
		}

	} else {

		// http.Redirect(w, r, "/login", 302)

		hash, _ := handler.HashPassword(user.Password)
		user.Password = hash

		handler.InsertData(user)

		err := h.ResponseHTML(w, r, "signup/signup", templateData{
			Message: "Sign Up successfully, you can login now!",
		})

		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	}
}

// NewSignupHTTPHandler responses new HTTPHandler instance.
func NewSignupHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
