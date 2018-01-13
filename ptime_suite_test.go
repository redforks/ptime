package ptime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPtime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ptime Suite")
}
