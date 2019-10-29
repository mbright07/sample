package logout

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Email string
	// Message string
}

var store = sessions.NewCookieStore([]byte("secret-key"))

// Logout page
func (h *HTTPHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "user-session")
	email := session.Values["email"].(string)

	err := h.ResponseHTML(w, r, "logout/logout", templateData{
		Email: email,
	})

	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// LogoutHandle page
func (h *HTTPHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/hello", 302)
}

// NewLogoutHTTPHandler responses new HTTPHandler instance.
func NewLogoutHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
