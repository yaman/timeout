package tcp

import (
	"fmt"
	"net"
	"time"

	"github.com/yaman/timeout/timeout"
)

func ListenAndDoNotAnswer() {
	timeout := timeout.Timeout{Port: ":5501"}
	listener := timeout.Listen()

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
	}
}

func ListenAndAnswerWithEmptyString() {
	timeout := timeout.Timeout{Port: ":5502"}
	ln := timeout.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Fprintf(conn, " ")
	}
}

func ListenAndAnswerWithMalformedStringImmediately() {
	timeout := timeout.Timeout{Port: ":5504"}
	ln := timeout.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		fmt.Fprintf(conn, "foo bar")
	}
}

func ListenAndAnswerWithEmptyStringAfterClientSendsData() {
	timeout := timeout.Timeout{Port: ":5503"}
	ln := timeout.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		tmp := make([]byte, 256)
		go func(connection net.Conn) {
			_, _ = connection.Read(tmp)
			fmt.Fprintf(connection, " ")
			connection.Close()
		}(conn)
	}
}

func ListenAndAnswerWithMalformedStringAfterClientSendsData() {
	timeout := timeout.Timeout{Port: ":5505"}
	ln := timeout.Listen()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		tmp := make([]byte, 256)
		go func(connection net.Conn) {
			_, _ = connection.Read(tmp)
			fmt.Fprintf(connection, "foo bar")
			connection.Close()
		}(conn)
	}
}

func ListenAndAnswerEvery5Seconds() {
	timeout := timeout.Timeout{Port: ":5506"}

	ln := timeout.Listen()
	conn, _ := ln.Accept()
	defer conn.Close()

	for {
		time.Sleep(5 * time.Second)
		go func(connection net.Conn) {
			fmt.Fprintf(connection, " ")
		}(conn)
	}
}

func ListenAndAnswerEvery30Seconds() {
	timeout := timeout.Timeout{Port: ":5507"}

	ln := timeout.Listen()
	conn, _ := ln.Accept()
	defer conn.Close()

	for {
		time.Sleep(30 * time.Second)
		go func(connection net.Conn) {
			fmt.Fprintf(connection, " ")
		}(conn)

	}

}
