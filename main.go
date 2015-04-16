package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/abdulkadiryaman/go-hamms/http"
	"github.com/abdulkadiryaman/go-hamms/tcp"
)

func main() {
	fmt.Println("Running Go hamms...")
	port := flag.String("port", "80", "port number")
	flag.Parse()

	go tcp.ListenAndDoNotAnswer()
	go tcp.ListenAndAnswerWithEmptyString()
	go tcp.ListenAndAnswerWithEmptyStringAfterClientSendsData()
	go tcp.ListenAndAnswerWithMalformedStringImmediately()
	go tcp.ListenAndAnswerWithMalformedStringAfterClientSendsData()
	go tcp.ListenAndAnswerEvery5Seconds()
	go tcp.ListenAndAnswerEvery30Seconds()
	go http.StartRouter(*port)

	for {
		time.Sleep(10 * time.Second)
	}

}
