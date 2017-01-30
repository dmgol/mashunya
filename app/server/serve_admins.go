package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dmgol/mashunya/config"
	"github.com/dmgol/mashunya/config/admin"
	"github.com/dmgol/mashunya/config/auth"
	"github.com/gorilla/csrf"
	"github.com/qor/qor/utils"
)

func ServeAdmins() {
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	mux.Handle("/auth/", auth.Auth.NewRouter())

	for _, path := range []string{"system", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.AdminsPort)
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
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.AdminsPort), skipCheck(handler)); err != nil {
		panic(err)
	}
}
