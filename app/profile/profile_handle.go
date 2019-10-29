package profile

import (
	"net/http"
	"os"
	"sample/app/shared/handler"

	"github.com/gorilla/schema"
	// "github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	User    handler.User
	Message string
}

// Profile page
func (h *HTTPHandler) Profile(w http.ResponseWriter, r *http.Request) {

	session, _ := handler.Store.Get(r, "user-session")

	if session.IsNew {
		http.Redirect(w, r, "/login", 302)

	} else {
		email := session.Values["email"].(string)

		user := handler.GetUserByEmail(email)

		err := h.ResponseHTML(w, r, "profile/profile", templateData{
			User:    user,
			Message: "",
		})

		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	}

}

// ProfileEdit page
func (h *HTTPHandler) ProfileEdit(w http.ResponseWriter, r *http.Request) {

	user := handler.User{}

	// r.ParseForm()
	r.ParseMultipartForm(32 << 20)

	decoder := schema.NewDecoder()
	decoder.Decode(&user, r.PostForm)

	// file, handle, err := r.FormFile("myAvatar")

	// defer func() {
	//     if rcv := recover(); rcv != nil {
	// 		handler.UpdateDataWithoutAvatar(user)
	// 		http.Redirect(w, r, "/profile", 302)
	//     }
	// }()

	// if err != nil {

	// }
	// defer file.Close()

	// f, err := os.OpenFile("static/images/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	// if err != nil {

	// }
	// defer f.Close()

	// io.Copy(f, file)

	session, _ := handler.Store.Get(r, "user-session")
	userSession := handler.GetUserByEmail(session.Values["email"].(string))

	if (len(user.Password) > 0) {
		hash, _ := handler.HashPassword(user.Password)
		user.Password = hash

	} else {
		user.Password = userSession.Password
	}

	fileName := handler.UploadFile(w, r)

	if fileName == "nofilefound" {
		handler.UpdateDataWithoutAvatar(user)

	} else {
		if len(userSession.Avatar) > 0 {
			os.Remove("static/images/" + userSession.Avatar)
		}

		user.Avatar = fileName
		handler.UpdateDataWithAvatar(user)

	}

	http.Redirect(w, r, "/profile", 302)

}

// NewProfileHTTPHandler responses new HTTPHandler instance.
func NewProfileHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
