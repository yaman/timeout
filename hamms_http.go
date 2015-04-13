package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func SleepFor() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/sleep/{SleepFor:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sleepFor := params["SleepFor"]
		w.Write([]byte("done!" + sleepFor))
	}).Methods("GET")

	http.Handle("/", rtr)

	http.ListenAndServe(":5508", nil)
}
