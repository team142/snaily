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

	if r.URL.Path == "/api/register/v1" {
		handleRegisterUser(w, r)
		return
	}

	if r.URL.Path == "/api/login/v1" {
		handleLoginUser(w, r)
		return
	}

	if r.URL.Path == "/api/new-item/v1" {
		handleCreateItem(w, r)
		return
	}

	if r.URL.Path == "/api/get-item/v1" {
		handleGetItem(w, r)
		return
	}

	if r.URL.Path == "/api/my-items/v1" {
		handleGetMyItems(w, r)
		return
	}

	msg := fmt.Sprint("Could not find route for ", r.URL.Path)
	logrus.Println(msg)
	w.Write([]byte(msg))
}
