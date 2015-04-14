package main

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hamms", func() {
	go StartRouter()
	time.Sleep(1 * time.Second)
	Describe("SleepFor", func() {
		Context("Http request sent to /sleep", func() {
			It("Should not be nil", func() {

				resp, err := http.Get("http://localhost:5508/sleep/5")
				defer resp.Body.Close()
				Expect(err).Should(BeNil())
				Expect(resp).ShouldNot(BeNil())

			})
			It("Should respond in 5 seconds", func() {
				start := time.Now()
				resp, _ := http.Get("http://localhost:5508/sleep/5")
				defer resp.Body.Close()
				elapsed := time.Since(start).Seconds()

				Expect(elapsed).Should(BeNumerically(">=", 4.5))
				Expect(elapsed).Should(BeNumerically("<=", 5.5))

			})

		})

	})
	Describe("ReturnStatus", func() {
		Context("Http request sent to /status", func() {
			It("Should respond with the status code 400", func() {
				resp, _ := http.Get("http://localhost:5508/status/400")
				defer resp.Body.Close()

				Expect(resp.StatusCode).Should(Equal(400))
			})
		})
	})
})
