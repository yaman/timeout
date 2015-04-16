package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const sleepPath = "/sleep/{SleepFor:[0-9]+}"
const sleepForParameterName = "SleepFor"
const respondWithStatusPath = "/status/{statuscode:[0-9]+}"
const respondWithDefaultStatusPath = "/status"
const statusCodeParameterName = "statuscode"

func StartRouter(port string) {
	fmt.Println("Hamms http starts with port:", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	http.ListenAndServe(":"+port, mux)
}

func router(w http.ResponseWriter, r *http.Request) {
	rawQuery := r.URL.RawQuery
	queryString, queryValue, _ := SplitRawQuery(rawQuery)
	switch {
	case queryString == "sleep":
		sleepFor(w, r, queryValue)
	case queryString == "status":
		respondWithStatus(w, r, queryValue)
	}
}

func respondWithStatus(w http.ResponseWriter, r *http.Request, queryValue string) {
	status, err := strconv.Atoi(queryValue)

	if err != nil {
		w.WriteHeader(200)
	} else {

		w.WriteHeader(status)
	}
	w.Write([]byte("done!"))
}

func sleepFor(w http.ResponseWriter, request *http.Request, queryValue string) {
	sleepFor, _ := strconv.Atoi(queryValue)
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}

func SplitRawQuery(rawQuery string) (string, string, error) {
	query := strings.Split(rawQuery, "=")
	switch len(query) {
	case 2:
		return query[0], query[1], nil
	}

	return "", "", errors.New("Malformed rawquery: " + rawQuery)
}
