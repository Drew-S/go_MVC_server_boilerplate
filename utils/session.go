package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("test"))

func GetSession(request *http.Request, name string) *sessions.Session {
	session, _ := store.Get(request, name)
	return session
}
