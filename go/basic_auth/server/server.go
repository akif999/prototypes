package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	basicAuthUser     = "user"
	basicAuthPassword = "pass"
)

func main() {
	http.HandleFunc("/basic",
		func(w http.ResponseWriter, r *http.Request) {
			if user, pass, ok := r.BasicAuth(); !ok || user != basicAuthUser || pass != basicAuthPassword {
				w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Not authorized", 401)
				return
			}
			_, err := fmt.Fprintf(w, "Authed\n")
			if err != nil {
				log.Fatal(err)
			}
		},
	)
	http.ListenAndServe(":18888", nil)
}
