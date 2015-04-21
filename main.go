package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yaman/timeout/http"
	"github.com/yaman/timeout/tcp"
)

func main() {
	port := os.Getenv("PORT")
	portParameter := flag.String("port", "8080", "port number to run http from")
	proto := flag.String("proto", "http", "protocol to run timeouts")

	flag.Parse()

	switch {
	case strings.Contains(*proto, "http"):

		fmt.Println("Running Go HTTP Timeout...")
		if len(port) > 0 {
			go http.StartRouter(port)
		} else if portParameter != nil {
			go http.StartRouter(*portParameter)
		}

	case strings.Contains(*proto, "tcp"):

		fmt.Println("Running Go TCP Timeout...")
		go tcp.ListenAndDoNotAnswer()
		go tcp.ListenAndAnswerWithEmptyString()
		go tcp.ListenAndAnswerWithEmptyStringAfterClientSendsData()
		go tcp.ListenAndAnswerWithMalformedStringImmediately()
		go tcp.ListenAndAnswerWithMalformedStringAfterClientSendsData()
		go tcp.ListenAndAnswerEvery5Seconds()
		go tcp.ListenAndAnswerEvery30Seconds()
	}

	for {
		time.Sleep(10 * time.Second)
	}
}
