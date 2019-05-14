package api

import (
	"fmt"
	"net/http"
)

func HandleIncoming(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin:", "http://localhost:8080")
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}

	if r.URL.Path == "/api/register" {
		handleRegisterUser(w, r)
		return
	}

	if r.URL.Path == "/api/login" {
		handleLoginUser(w, r)
		return
	}

	msg := fmt.Sprint("Could not find route for ", r.URL.Path)
	fmt.Println(msg)
	w.Write([]byte(msg))
}
