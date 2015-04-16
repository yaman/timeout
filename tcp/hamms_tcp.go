package tcp

import (
	"fmt"
	"net"
	"time"

	"github.com/abdulkadiryaman/go-hamms/hamms"
)

func ListenAndDoNotAnswer() {
	hamms := hamms.Hamms{Port: ":5501"}
	listener := hamms.Listen()

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5501")
	}
}

func ListenAndAnswerWithEmptyString() {
	hamms := hamms.Hamms{Port: ":5502"}
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
	hamms := hamms.Hamms{Port: ":5504"}
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
	hamms := hamms.Hamms{Port: ":5503"}
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

func ListenAndAnswerWithMalformedStringAfterClientSendsData() {
	hamms := hamms.Hamms{Port: ":5505"}
	ln := hamms.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted a connection from :5505")

		tmp := make([]byte, 256)
		go func(connection net.Conn) {
			_, _ = connection.Read(tmp)
			fmt.Fprintf(connection, "foo bar")
			connection.Close()
		}(conn)
	}
}

func ListenAndAnswerEvery5Seconds() {
	hamms := hamms.Hamms{Port: ":5506"}

	ln := hamms.Listen()
	conn, _ := ln.Accept()
	defer conn.Close()

	for {
		fmt.Println("Accepted a connection from :5506")
		time.Sleep(5 * time.Second)
		go func(connection net.Conn) {
			fmt.Fprintf(connection, " ")
		}(conn)
	}
}

func ListenAndAnswerEvery30Seconds() {
	hamms := hamms.Hamms{Port: ":5507"}

	ln := hamms.Listen()
	conn, _ := ln.Accept()
	defer conn.Close()

	for {
		fmt.Println("Accepted a conection from :5507")
		time.Sleep(30 * time.Second)
		go func(connection net.Conn) {
			fmt.Fprintf(connection, " ")
		}(conn)

	}

}
