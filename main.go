package main

import (
	"fmt"
	"time"

	"github.com/abdulkadiryaman/go-hamms/http"
	"github.com/abdulkadiryaman/go-hamms/tcp"
)

func main() {
	fmt.Println("Running Go hamms...")

	go tcp.ListenAndDoNotAnswer()
	go tcp.ListenAndAnswerWithEmptyString()
	go tcp.ListenAndAnswerWithEmptyStringAfterClientSendsData()
	go tcp.ListenAndAnswerWithMalformedStringImmediately()
	go tcp.ListenAndAnswerWithMalformedStringAfterClientSendsData()
	go tcp.ListenAndAnswerEvery5Seconds()
	go tcp.ListenAndAnswerEvery30Seconds()
	go http.StartRouter()

	for {
		time.Sleep(10 * time.Second)
	}

}
