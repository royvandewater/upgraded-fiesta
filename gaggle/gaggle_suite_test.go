package gaggle_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGaggle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gaggle Suite")
}
