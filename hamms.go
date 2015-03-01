package main

import (
	"fmt"
	"net"
	"time"
)

type Hamms struct {
	Port string
}

func (hamms *Hamms) Listen() net.Listener {
	listener, err := net.Listen("tcp", hamms.Port)
	if err != nil {
		fmt.Println("An Error Occured while trying to open port from ", hamms.Port, err)
	}
	fmt.Println("Listening from ", hamms.Port)
	return listener
}

func ListenAndDoNotAnswer() {
	hamms := Hamms{":5501"}
	listener := hamms.Listen()

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5501")
	}
}

func ListenAndAnswerWithEmptyString() {
	hamms := Hamms{":5502"}
	ln := hamms.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5502")
		fmt.Println("Writing to connection")
		fmt.Fprintf(conn, " ")
	}
}

func ListenAndAnswerWithMalformedStringImmediately() {
	hamms := Hamms{":5504"}
	ln := hamms.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5504")
		fmt.Println("Writing to connection")
		fmt.Fprintf(conn, "foo bar")
	}
}

func ListenAndAnswerWithEmptyStringAfterClientSendsData() {
	hamms := Hamms{":5503"}
	ln := hamms.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5503")

		tmp := make([]byte, 256)
		go func(connection net.Conn) {
			_, _ = connection.Read(tmp)
			fmt.Fprintf(connection, " ")
			connection.Close()
		}(conn)
	}
}

func main() {
	fmt.Println("Running Go Hamms.....")
	go ListenAndDoNotAnswer()
	go ListenAndAnswerWithEmptyString()
	go ListenAndAnswerWithEmptyStringAfterClientSendsData()
	go ListenAndAnswerWithMalformedStringImmediately()

	for {
		time.Sleep(10 * time.Second)
	}
}
