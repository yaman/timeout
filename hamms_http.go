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

	http.Handle("/", router)
	http.ListenAndServe(listenAddress, nil)
}

func sleepFor(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	sleepFor, _ := strconv.Atoi(params[sleepForParameterName])
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}
