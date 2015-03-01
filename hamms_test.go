package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListenTcp(t *testing.T) {

	Convey("Given hamms created with port :5000", t, func() {
		hamms := Hamms{":5000"}
		
		Convey("When hamms.Listen() called",func(){
			listener := hamms.Listen()

			Convey("Listener Should not be null",func(){
				So(listener,ShouldNotBeNil)
			})
		})
	})

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
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5502")
		if err != nil {
			fmt.Println("error is ", err)
		}
		Convey("It should write back space immediately", func() {
			tmp := make([]byte, 64)
			_, err = conn.Read(tmp)
			So(tmp[0], ShouldEqual, 32)
		})
	})
	Convey("Given connected to :5503 port", t, func() {
		go ListenAndAnswerWithEmptyStringAfterClientSendsData()
		time.Sleep(3 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5503")
		if err != nil {
			fmt.Println("error is ", err)
		}
		Convey("when sent some random data", func() {
			fmt.Fprintf(conn, "RANDOM_DATA")
			Convey("It should write back space", func() {
				tmp := make([]byte, 64)
				_, err = conn.Read(tmp)
				So(tmp[0], ShouldEqual, 32)
			})
		})

	})
	Convey("Given connected to :5504 port", t, func() {
		go ListenAndAnswerWithMalformedStringImmediately()
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5504")
		if err != nil {
			fmt.Println("error is ", err)
		}
		Convey("It should write back malformed response immediately", func() {
			tmp := make([]byte, 64)
			_, err = conn.Read(tmp)
			responseString := string(tmp[:])
			So(responseString, ShouldContainSubstring, "foo bar")
		})
	})

}
