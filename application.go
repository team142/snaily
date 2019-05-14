package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/team142/snaily/api"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		// TODO: host locally
	} else {
		//When running locally - reverse proxy to node js server
		router.PathPrefix("/").HandlerFunc(buildHomeRouter())
	}

	//The server
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func buildHomeRouter() func(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse(*addrToProxy)
	rp := httputil.NewSingleHostReverseProxy(u)
	return func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)
	}
}
