package main

import (
	"github.com/garvit4540/go-mysql-bookstore/pkg/routes"
	"github.com/garvit4540/go-mysql-bookstore/pkg/trace"
	"github.com/gorilla/mux"
	"net/http"
)

const port = "3000"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:"+port, r)
	if err != nil {
		trace.PushTrace(trace.ErrorStartingApp, map[string]interface{}{"error": err})
		panic(err)
	}
	trace.PushTrace(trace.AppStarted, map[string]interface{}{"port": port})
}
