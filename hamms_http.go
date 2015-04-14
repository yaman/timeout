package main

import (
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
)

func SleepFor() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/sleep/{SleepFor:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sleepFor, _ := strconv.Atoi(params["SleepFor"])
		time.Sleep(time.Duration(sleepFor) * time.Second)
		w.Write([]byte("done!"))
	}).Methods("GET")

	http.Handle("/", rtr)

	http.ListenAndServe(":5508", nil)
}
