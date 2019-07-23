package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	auth "github.com/abbot/go-http-auth"
)

const (
	digestAuthUser     = "user"
	digestAuthPassword = "pass"
)

func secret(user, realm string) string {
	if user == digestAuthUser {
		// return fmt.Sprintf("%x", md5.Sum([]byte(digestAuthPassword)))
		return "1a1dc91c907325c69271ddf0c944bc72"
	}
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func main() {
	s := fmt.Sprintf("%x", md5.Sum([]byte(digestAuthPassword)))
	fmt.Println(s)

	authenticator := auth.NewDigestAuthenticator("Secret Zone", secret)
	http.HandleFunc("/", authenticator.Wrap(handle))
	log.Println(http.ListenAndServe(":8080", nil))
}
