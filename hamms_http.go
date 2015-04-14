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
const respondWithStatusPath = "/status/{statuscode:[0-9]+}"
const respondWithDefaultStatusPath = "/status"
const statusCodeParameterName = "statuscode"

func StartRouter() {
	router := mux.NewRouter()

	router.HandleFunc(sleepPath, sleepFor).Methods("GET")
	router.HandleFunc(respondWithStatusPath, respondWithStatus).Methods("GET")
	router.HandleFunc(respondWithDefaultStatusPath, respondWithStatus).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(listenAddress, nil)
}

func respondWithStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status, err := strconv.Atoi(params[statusCodeParameterName])

	if err != nil {
		w.WriteHeader(200)
	} else {

		w.WriteHeader(status)
	}
	w.Write([]byte("done!"))
}

func sleepFor(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	sleepFor, _ := strconv.Atoi(params[sleepForParameterName])
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}
