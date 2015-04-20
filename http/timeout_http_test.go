package http

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Timeout", func() {

	Describe("Router", func() {
		go StartRouter("5508")
		time.Sleep(1 * time.Second)

		Describe("SleepFor", func() {
			Context("Http request sent to /anyresourcepath?sleep=5", func() {
				It("Response should not be nil", func() {

					resp, err := http.Get("http://localhost:5508/anyresourcepath?sleep=5")
					defer resp.Body.Close()
					Expect(err).Should(BeNil())
					Expect(resp).ShouldNot(BeNil())

				})
				It("Should respond in 5 seconds", func() {
					start := time.Now()
					resp, _ := http.Get("http://localhost:5508/anyresourcepath?sleep=5")
					defer resp.Body.Close()
					elapsed := time.Since(start).Seconds()

					Expect(elapsed).Should(BeNumerically(">=", 4.5))
					Expect(elapsed).Should(BeNumerically("<=", 5.5))

				})

			})

		})

		Describe("ReturnStatus", func() {
			Context("Http request sent to /anyresource?status=400", func() {
				It("Should respond with the status code 400", func() {
					resp, _ := http.Get("http://localhost:5508/anyresource?status=400")
					defer resp.Body.Close()

					Expect(resp.StatusCode).Should(Equal(400))
				})
			})

			Context("Http request sent to /mypreciuosresource?status=500", func() {
				It("Should respond with the status code 500", func() {
					resp, _ := http.Get("http://localhost:5508/mypreciuosresource?status=500")
					defer resp.Body.Close()

					Expect(resp.StatusCode).Should(Equal(500))
				})
			})

			Context("Http request sent to /status", func() {
				It("Should respond with the status code 200", func() {
					resp, _ := http.Get("http://localhost:5508/status")
					defer resp.Body.Close()

					Expect(resp.StatusCode).Should(Equal(200))
				})
			})
		})

		Describe("FakeTheSize", func() {
			Context("Http request sent to /anyresource?fakethesize=3&for=1024", func() {
				It("Should respond with the status code 200", func() {
					resp, _ := http.Get("http://localhost:5508/anyresource?fakethesize=3&for=1024")
					defer resp.Body.Close()

					Expect(resp.StatusCode).Should(Equal(200))
				})

				It("Should respond with content-lenght=3", func() {
					resp, _ := http.Get("http://localhost:5508/anyresource?fakethesize=3&for=1024")
					defer resp.Body.Close()

					Expect(resp.Header.Get("Content-Length")).Should(Equal("3"))
				})

			})
		})
	})

	Describe("SplitRawQuery", func() {
		It("Should split query name for sleep", func() {
			rawQuery := "sleep=5"
			actual, _, _ := SplitRawQuery(rawQuery)

			Expect(actual).Should(Equal("sleep"))
		})
		It("Should split query field for sleep", func() {
			rawQuery := "sleep=5"
			_, actual, _ := SplitRawQuery(rawQuery)

			Expect(actual).Should(Equal("5"))
		})
		It("Should split query name status", func() {
			rawQuery := "status=500"
			actual, _, _ := SplitRawQuery(rawQuery)

			Expect(actual).Should(Equal("status"))
		})
		It("Should split query field for status", func() {
			rawQuery := "status=500"
			_, actual, _ := SplitRawQuery(rawQuery)

			Expect(actual).Should(Equal("500"))
		})
		It("Should return error if format is wrong", func() {
			rawQuery := "status500"
			_, _, actualErr := SplitRawQuery(rawQuery)

			Expect(actualErr).ShouldNot(Equal(nil))
		})

	})

	Describe("SplitRawQueryIntoMap", func() {
		It("Should return a map with size 2", func() {
			rawQuery := "fakethesize=3&for=1024"
			actual, _ := SplitRawQueryIntoMap(rawQuery)

			Expect(len(actual)).Should(Equal(2))
		})
		It("Should return a map which contains fakethesize and for", func() {
			rawQuery := "fakethesize=3&for=1024"
			actual, _ := SplitRawQueryIntoMap(rawQuery)
			_, fakeExists := actual["fakethesize"]
			_, forExists := actual["for"]
			Expect(fakeExists).Should(Equal(true))
			Expect(forExists).Should(Equal(true))
		})
		It("Should return a map with fakethesize 3 and for 1024", func() {
			rawQuery := "fakethesize=3&for=1024"
			actual, _ := SplitRawQueryIntoMap(rawQuery)
			fakethesizeValue, _ := actual["fakethesize"]
			forValue, _ := actual["for"]
			Expect(fakethesizeValue).Should(Equal("3"))
			Expect(forValue).Should(Equal("1024"))
		})
		It("Should return error if the format is not right for one query parameter", func() {
			rawQuery := "fakethesize1024"
			_, err := SplitRawQueryIntoMap(rawQuery)
			Expect(err).ShouldNot(Equal(nil))
		})
		It("Should return error if the format is not right for multiple query parameters", func() {
			rawQuery := "fakethesize1024&asdasdasda00123"
			_, err := SplitRawQueryIntoMap(rawQuery)
			Expect(err).ShouldNot(Equal(nil))
		})

	})

})
