package auth

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"gopkg.in/authboss.v0"
)

var CookieStore *securecookie.SecureCookie

func init() {
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`Z2hqZHRocmZjZnFuZnlmZGJoZWNz==`)
	CookieStore = securecookie.New(cookieStoreKey, nil)
}

type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	log.Println("NewCookieStorer...")
	return &CookieStorer{w, r}
}

func (s CookieStorer) Get(key string) (string, bool) {
	log.Println("CookieStorer.Get ...", key)
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = CookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

func (s CookieStorer) Put(key, value string) {
	log.Println("CookieStorer.Put ...", key, value)
	encoded, err := CookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

func (s CookieStorer) Del(key string) {
	log.Println("CookieStorer.Del ...", key)
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
