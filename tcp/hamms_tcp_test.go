package tcp

import (
	"fmt"
	"net"
	"time"

	"github.com/abdulkadiryaman/go-hamms/hamms"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hamms", func() {

	Describe("Given hamms created with port :5000", func() {
		hamms := hamms.Hamms{Port: ":5000"}

		Context("When hamms.Listen() called", func() {
			listener := hamms.Listen()
			It("Listener Should not be null", func() {
				Expect(listener).ShouldNot(Equal(nil))
			})
			hamms.Close()
		})

	})

	Describe("Given connected to :5501 port", func() {
		go ListenAndDoNotAnswer()
		time.Sleep(1 * time.Second)
		conn, _ := net.Dial("tcp", "localhost:5501")
		Context("When waited for 3 seconds", func() {
			It("It should not write back any data", func() {
				tmp := make([]byte, 256)
				go func(connection net.Conn) {
					_, _ = connection.Read(tmp)
				}(conn)
				time.Sleep(3 * time.Second)
			})
		})
	})

	Describe("Given connected to :5502 port", func() {
		go ListenAndAnswerWithEmptyString()
		time.Sleep(1 * time.Second)
		conn, _ := net.Dial("tcp", "localhost:5502")
		It("It should write back space immediately", func() {
			tmp := make([]byte, 64)
			_, _ = conn.Read(tmp)

			Expect(tmp[0]).Should(Equal(uint8(32)))
		})
	})
	Describe("Given connected to :5503 port", func() {
		go ListenAndAnswerWithEmptyStringAfterClientSendsData()
		time.Sleep(3 * time.Second)
		conn, _ := net.Dial("tcp", "localhost:5503")
		Context("when sent some random data", func() {
			fmt.Fprintf(conn, "RANDOM_DATA")
			It("It should write back space", func() {
				tmp := make([]byte, 64)
				_, _ = conn.Read(tmp)
				Expect(tmp[0]).Should(Equal(uint8(32)))
			})
		})

	})
	Describe("Given connected to :5504 port", func() {
		go ListenAndAnswerWithMalformedStringImmediately()
		time.Sleep(1 * time.Second)
		conn, _ := net.Dial("tcp", "localhost:5504")
		It("It should write back malformed response immediately", func() {
			tmp := make([]byte, 64)
			_, _ = conn.Read(tmp)
			responseString := string(tmp[:])
			Expect(responseString).Should(ContainSubstring("foo bar"))
		})
	})

	Describe("Given connected to :5505 port", func() {
		go ListenAndAnswerWithMalformedStringAfterClientSendsData()
		time.Sleep(3 * time.Second)
		conn, _ := net.Dial("tcp", "localhost:5505")
		Context("when sent some random data", func() {
			fmt.Fprintf(conn, "RANDOM_DATA")
			It("It should write back space", func() {
				tmp := make([]byte, 64)
				_, _ = conn.Read(tmp)
				responseString := string(tmp[:])
				Expect(responseString).Should(ContainSubstring("foo bar"))
			})
		})

	})

	XDescribe("Given connected to :5506 port", func() {
		go ListenAndAnswerEvery5Seconds()
		time.Sleep(3 * time.Second)

		It("It should write back one byte data every 5 seconds", func() {
			conn, _ := net.Dial("tcp", "localhost:5506")
			tmp := make([]byte, 64)
			maxAssertionCount := 3
			for i := 0; i < maxAssertionCount; i++ {
				t0 := time.Now()
				_, _ = conn.Read(tmp)
				responseString := string(tmp[:])
				t1 := time.Now()
				actualTimeElapsed := t1.Sub(t0).Seconds()
				Expect(actualTimeElapsed).Should(BeNumerically(">=", 4.5))
				Expect(actualTimeElapsed).Should(BeNumerically("<=", 7.5))
				Expect(responseString).Should(ContainSubstring(" "))
			}
		})

	})

	XDescribe("Given connected to :5507 port", func() {
		go ListenAndAnswerEvery30Seconds()
		time.Sleep(2 * time.Second)

		It("It should write back one byte data every 30 seconds", func() {
			conn, _ := net.Dial("tcp", "localhost:5507")
			tmp := make([]byte, 64)
			maxAssertionCount := 1
			for i := 0; i < maxAssertionCount; i++ {
				t0 := time.Now()
				_, _ = conn.Read(tmp)
				responseString := string(tmp[:])
				t1 := time.Now()
				actualTimeElapsed := t1.Sub(t0).Seconds()

				Expect(actualTimeElapsed).Should(BeNumerically(">", 25))
				Expect(actualTimeElapsed).Should(BeNumerically("<", 35))
				Expect(responseString).Should(ContainSubstring(" "))
			}
		})

	})
})
