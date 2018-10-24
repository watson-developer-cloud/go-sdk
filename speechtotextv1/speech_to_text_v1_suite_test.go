// +build !integration

package speechtotextv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpeechToTextV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpeechToTextV1 Suite")
}
