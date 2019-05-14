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

const StaticDir = "web/static/"

var rp = buildRP()

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.PathPrefix("/api").HandlerFunc(api.HandleIncoming)
	router.PathPrefix("/").HandlerFunc(handleHome)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func buildRP() *httputil.ReverseProxy {
	u, err := url.Parse("http://localhost:4200")
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(u)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	rp.ServeHTTP(w, r)
	//w.Write([]byte("1"))
	return
}
