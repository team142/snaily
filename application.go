package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/api"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var addr = flag.String("address", ":8080", "")
var addrToProxy = flag.String("proxy", "http://localhost:4200", "The url of the angular app to reverse proxy")
var container = flag.Bool("container", false, "The url of the angular app to reverse proxy")

func main() {
	flag.Parse()

	router := mux.NewRouter()

	//Handles all API calls
	router.PathPrefix("/api").HandlerFunc(api.HandleIncoming)

	//Handles everything else
	if *container {
		//Serve from web folder
		router.PathPrefix("/").HandlerFunc(staticFileServer)
	} else {
		//When running locally - reverse proxy to node js server
		router.PathPrefix("/").HandlerFunc(buildHomeRouter())
	}

	//The server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func staticFileServer(w http.ResponseWriter, r *http.Request) {
	name := fmt.Sprint("/snaily-web", r.URL.Path)

	if name == "/snaily-web/" {
		name = "/snaily-web/index.html"
	}

	logrus.Println(name)

	if f, err := os.Stat(name); err == nil && !f.IsDir() {
		http.ServeFile(w, r, name)
		return
	}

	logrus.Println(name)
	http.NotFound(w, r)

}

func buildHomeRouter() func(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(*addrToProxy)
	rp := httputil.NewSingleHostReverseProxy(u)
	return func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	}
}
