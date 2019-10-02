// +build !integration

package visualrecognitionv4_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVisualRecognitionV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VisualRecognitionV4 Suite")
}
