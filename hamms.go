package main

import (
	"fmt"
	"net"
	"time"
)

func ListenAndDoNotAnswer() {
	ln, err := net.Listen("tcp", ":5501")
	if err != nil {
		fmt.Println("An Error Occured while trying to open port from :5501", err)
	}

	fmt.Println("Listening from :5501")

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5501")
		if err != nil {
			// handle error
		}
	}
}

func ListenAndAnswerWithEmptyString() {
	ln, err := net.Listen("tcp", ":5502")
	if err != nil {
		fmt.Println("An Error Occured while trying to open port from :5502", err)
	}

	fmt.Println("Listening from :5502")

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5502")
		if err != nil {
			fmt.Println("Something is wrong with listener : ", err)
		}
		fmt.Println("Writing to connection")
		fmt.Fprintf(conn, " ")
	}
}

func ListenAndAnswerWithEmptyStringAfterClientSendsData() {
	ln, err := net.Listen("tcp", ":5503")
	if err != nil {
		fmt.Println("An Error Occured while trying to open port from :5503", err)
	}

	fmt.Println("Listening from :5503")

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5503")
		if err != nil {
			fmt.Println("Something is wrong with listener : ", err)
		}
		tmp := make([]byte, 256)
		go func(connection net.Conn) {
			_, err := connection.Read(tmp)
			if err != nil {

			}

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

	for {
		time.Sleep(10 * time.Second)
	}
}
