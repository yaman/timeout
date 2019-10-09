package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"yaman/timeout/http"
	"yaman/timeout/tcp"
)

func main() {
	portParameter := flag.String("port", "8080", "port number to run http from")
	protoParameter := flag.String("proto", "http", "protocol to run timeouts")
	flag.Parse()

	port := os.Getenv("PORT")
	if len(port) == 0 && len(*portParameter) > 0 {
		port = *portParameter
	}
	proto := os.Getenv("PROTO")
	if len(proto) == 0 && len(*protoParameter) > 0 {
		proto = *protoParameter
	}

	switch {
	case strings.Contains(proto, "http"):

		fmt.Println("Running Go HTTP Timeout...")
		go http.StartRouter(port)

	case strings.Contains(proto, "tcp"):

		fmt.Println("Running Go TCP Timeout...")
		go tcp.ListenAndDoNotAnswer()
		go tcp.ListenAndAnswerWithEmptyString()
		go tcp.ListenAndAnswerWithEmptyStringAfterClientSendsData()
		go tcp.ListenAndAnswerWithMalformedStringImmediately()
		go tcp.ListenAndAnswerWithMalformedStringAfterClientSendsData()
		go tcp.ListenAndAnswerEvery5Seconds()
		go tcp.ListenAndAnswerEvery30Seconds()
	}

	select {}
}
