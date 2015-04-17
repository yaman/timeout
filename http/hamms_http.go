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
	fmt.Println("Starting Hamms http with port:", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	http.ListenAndServe(":"+port, mux)
}

func router(w http.ResponseWriter, r *http.Request) {
	rawQuery := r.URL.RawQuery
	var queryString, queryValue string
	var queryMap map[string]string
	var fakeSize string
	var actualSize string
	if strings.Contains(rawQuery, "&") {
		queryMap, _ = SplitRawQueryIntoMap(rawQuery)
		var found bool
		fakeSize, found = queryMap["fakethesize"]
		if found {
			actualSize = queryMap["for"]
			queryString = "fakethesize"
		}
	} else {
		queryString, queryValue, _ = SplitRawQuery(rawQuery)
	}

	switch {
	case queryString == "sleep":
		sleepFor(w, r, queryValue)
	case queryString == "status":
		respondWithStatus(w, r, queryValue)
	case queryString == "fakethesize":
		respondWithFakeSize(w, r, fakeSize, actualSize)
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

func respondWithFakeSize(w http.ResponseWriter, r *http.Request, fakeSize string, actualSize string) {
	w.Header().Set("Content-Length", fakeSize)
}

func sleepFor(w http.ResponseWriter, request *http.Request, queryValue string) {
	sleepFor, _ := strconv.Atoi(queryValue)
	time.Sleep(time.Duration(sleepFor) * time.Second)
	w.Write([]byte("done!"))
}

func SplitRawQueryIntoMap(rawQuery string) (map[string]string, error) {
	query := strings.Split(rawQuery, "&")

	queryMap := make(map[string]string)
	for _, element := range query {
		queryString, queryValue, err := SplitRawQuery(element)
		if err != nil {
			return nil, err
		}
		queryMap[queryString] = queryValue
	}
	return queryMap, nil
}

func SplitRawQuery(rawQuery string) (string, string, error) {
	query := strings.Split(rawQuery, "=")
	switch len(query) {
	case 2:
		return query[0], query[1], nil
	}

	return "", "", errors.New("Malformed rawquery: " + rawQuery)
}
