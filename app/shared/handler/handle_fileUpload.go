package handler

import (
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) string {
	file, handle, err := r.FormFile("myAvatar")

	if err != nil {
		return "nofilefound"
	}
	defer file.Close()

	f, err := os.OpenFile("static/images/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

	}
	defer f.Close()

	io.Copy(f, file)

	return handle.Filename
}
