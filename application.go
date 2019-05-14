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

const StaticDir = "web/static/"

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/api").HandlerFunc(api.HandleIncoming)
	router.PathPrefix("/").HandlerFunc(buildHomeRouter())
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
