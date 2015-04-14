package main

import (
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
)

const listenAddress = ":5508"
const sleepPath = "/sleep/{SleepFor:[0-9]+}"
const sleepForParameterName = "SleepFor"

func StartRouter() {
	router := mux.NewRouter()
	router.HandleFunc(sleepPath, sleepFor).Methods("GET")
	router.HandleFunc("/status/{statuscode:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		status, _ := strconv.Atoi(params["statuscode"])
		w.WriteHeader(status)
		w.Write([]byte("done!"))
	}).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(listenAddress, nil)
}

func sleepFor(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	sleepFor, _ := strconv.Atoi(params[sleepForParameterName])
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}
