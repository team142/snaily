package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleIncoming(w http.ResponseWriter, r *http.Request) {
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
	logrus.Println(msg)
	w.Write([]byte(msg))
}
