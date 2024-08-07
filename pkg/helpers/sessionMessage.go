package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	hashKey  = []byte("L0B8W1/5cNf+vK8Re4c2wQ==") // 32 bytes when decoded
	blockKey = []byte("Rm9iQzZJR2Z2S3EzMkN/3Q==") // 32 bytes when decoded
)

var store = sessions.NewCookieStore(hashKey, blockKey)

// GetFlashMessage retrieves the flash message from the session.
func GetFlashMessage(w http.ResponseWriter, r *http.Request) string {
	session, err := store.Get(r, "flash")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	message := session.Flashes("message")
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	if len(message) > 0 {
		return message[0].(string)
	}
	return ""
}

// SetFlashMessage sets a flash message in the session.
func SetFlashMessage(w http.ResponseWriter, r *http.Request, message string) {
	// Get a session.
	session, err := store.Get(r, "flash")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(message, "message")
	session.Save(r, w)

}
