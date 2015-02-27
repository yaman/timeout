package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListenTcp(t *testing.T) {

	Convey("Given connected to :5501 port", t, func() {
		go ListenAndDoNotAnswer()

		conn, err := net.Dial("tcp", "localhost:5501")
		if err != nil {
			fmt.Println("error is ", err)
		}

		Convey("When waited for 3 seconds", func() {
			Convey("It should not write back any data", func() {
				tmp := make([]byte, 256)
				go func(connection net.Conn) {
					_, err := connection.Read(tmp)
					if err != nil {
					}
					t.FailNow()
				}(conn)
				time.Sleep(3 * time.Second)
			})
		})
	})

	Convey("Given connected to :5502 port", t, func() {
		go ListenAndAnswerWithEmptyString()
		Convey("When waited for 1 seconds", func() {
			time.Sleep(1 * time.Second)
			Convey("It should write back space immediately", func() {
				conn, err := net.Dial("tcp", "localhost:5502")
				if err != nil {
					fmt.Println("error is ", err)
				}
				tmp := make([]byte, 64)
				_, err = conn.Read(tmp)
				So(tmp[0], ShouldEqual, 32)
			})
		})
	})
}
