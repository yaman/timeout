package tcp

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoHamms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timeout Tcp Test Suite")
}
