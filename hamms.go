package main

import (
	"fmt"
	"net"
	"time"
)

type Hamms struct {
	Port     string
	listener net.Listener
}

func (hamms *Hamms) Listen() net.Listener {
	listener, err := net.Listen("tcp", hamms.Port)
	if err != nil {
		panic("An Error Occured while trying to open port: " + hamms.Port)
	}
	hamms.listener = listener
	return listener
}

func (hamms *Hamms) Close() error {
	return hamms.listener.Close()
}

func main() {
	fmt.Println("Running Go hamms...")

	go ListenAndDoNotAnswer()
	go ListenAndAnswerWithEmptyString()
	go ListenAndAnswerWithEmptyStringAfterClientSendsData()
	go ListenAndAnswerWithMalformedStringImmediately()
	go ListenAndAnswerWithMalformedStringAfterClientSendsData()
	go ListenAndAnswerEvery5Seconds()
	go ListenAndAnswerEvery30Seconds()
	go StartRouter()

	for {
		time.Sleep(10 * time.Second)
	}

}
