package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dmgol/mashunya/config"
	"github.com/dmgol/mashunya/config/admin"
	"github.com/dmgol/mashunya/config/auth"
	_ "github.com/dmgol/mashunya/db/migrations"
	"github.com/gorilla/csrf"
)

func main() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	mux.Handle("/auth/", auth.Auth.NewRouter())

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	skipCheck := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/auth") {
				r = csrf.UnsafeSkipCheck(r)
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	handler := csrf.Protect([]byte("3693f371bf91487c99286a777811bd4e"), csrf.Secure(false))(mux)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), skipCheck(handler)); err != nil {
		panic(err)
	}
}
