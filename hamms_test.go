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

		Convey("When hamms.Listen() called", func() {
			listener := hamms.Listen()

			Convey("Listener Should not be null", func() {
				So(listener, ShouldNotBeNil)
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
	Convey("Given connected to :5505 port", t, func() {
		go ListenAndAnswerWithMalformedStringAfterClientSendsData()
		time.Sleep(3 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5505")
		if err != nil {
			fmt.Println("error is ", err)
		}
		Convey("when sent some random data", func() {
			fmt.Fprintf(conn, "RANDOM_DATA")
			Convey("It should write back space", func() {
				tmp := make([]byte, 64)
				_, err = conn.Read(tmp)
				responseString := string(tmp[:])
				So(responseString, ShouldContainSubstring, "foo bar")
			})
		})

	})

	Convey("Given connected to :5506 port", t, func() {
		go ListenAndAnswerEvery5Seconds()
		time.Sleep(3 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5506")
		if err != nil {
			fmt.Println("error is ", err)
		}

		Convey("It should write back one byte data every 5 seconds", func() {
			tmp := make([]byte, 64)
			maxAssertionCount := 3
			for i := 0; i < maxAssertionCount; i++ {
				t0 := time.Now()
				_, err = conn.Read(tmp)
				responseString := string(tmp[:])
				t1 := time.Now()
				actualTimeElapsed := t1.Sub(t0).Seconds()

				So(actualTimeElapsed, ShouldBeBetween, 4.5, 5.5)
				So(responseString, ShouldContainSubstring, " ")
			}
		})

	})

	Convey("Given connected to :5507 port", t, func() {
	    go ListenAndAnswerEvery30Seconds()	
		time.Sleep(3 * time.Second)
		conn, err := net.Dial("tcp", "localhost:5507")
		if err != nil {
			fmt.Println("error is ", err)
		}

		Convey("It should write back one byte data every 30 seconds", func() {
			tmp := make([]byte, 64)
			maxAssertionCount := 1
			for i := 0; i < maxAssertionCount; i++ {
				t0 := time.Now()
				_, err = conn.Read(tmp)
				responseString := string(tmp[:])
				t1 := time.Now()
				actualTimeElapsed := t1.Sub(t0).Seconds()

				So(actualTimeElapsed, ShouldBeBetween, 29, 31)
				So(responseString, ShouldContainSubstring, " ")
			}
		})

	})
}
