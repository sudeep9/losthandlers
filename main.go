package main

import (
	"log"
	"losthandlers/html"
	"net/http"

	"github.com/gorilla/mux"
	ds "github.com/starfederation/datastar/sdk/go"
)

const (
	ServerAddr = "127.0.0.1:8000"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html.TestHome().Render(r.Context(), w)
	}).Methods("GET")

	r.HandleFunc("/alltabs", func(w http.ResponseWriter, r *http.Request) {
		sse := ds.NewSSE(w, r)
		sse.MergeFragmentTempl(html.Loadtabs(),
			ds.WithSelectorID("app"),
			ds.WithMergeMode(ds.FragmentMergeModeInner))
	}).Methods("GET")

	r.HandleFunc("/other", func(w http.ResponseWriter, r *http.Request) {
		sse := ds.NewSSE(w, r)
		sse.MergeFragmentTempl(html.Div("Some other content"),
			ds.WithSelectorID("app"),
			ds.WithMergeMode(ds.FragmentMergeModeInner))
	}).Methods("GET")
}

func main() {

	router := mux.NewRouter()
	registerRoutes(router)

	srv := &http.Server{
		Handler: router,
		Addr:    ServerAddr,
	}

	log.Printf("starting server at %s ...", ServerAddr)
	srv.ListenAndServe()
}
