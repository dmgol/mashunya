package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/authboss.v0"
)

const sessionCookieName = "mashunya"

var sessionStore *sessions.CookieStore

type SessionStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func init() {
	//sessionStoreKey, _ := base64.StdEncoding.DecodeString(`Z2hqZHRocmZjZnFuZnlmZGJoZWNz`)
	sessionStoreKey := []byte("TestKey")
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
}

func NewSessionStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	log.Println("NewSessionStorer...")
	return &SessionStorer{w, r}
}

func (s SessionStorer) Get(key string) (string, bool) {
	log.Println("SessionStorer.Get ...", key)
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

func (s SessionStorer) Put(key, value string) {
	log.Println("SessionStorer.Put ...", key, value)
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values[key] = value
	session.Save(s.r, s.w)
}

func (s SessionStorer) Del(key string) {
	log.Println("SessionStorer.Del ...", key)
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(session.Values, key)
	session.Save(s.r, s.w)
}
