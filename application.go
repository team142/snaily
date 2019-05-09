package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
)

var addr = flag.String("address", ":8080", "")

const StaticDir = "web/static/"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome)
	router.HandleFunc("/api/", handleAPI)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		logrus.Errorln(err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		logrus.Errorln(err)
		return
	}
}
