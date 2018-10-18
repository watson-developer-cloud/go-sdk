// +build !integration

package texttospeechv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTextToSpeechV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TextToSpeechV1 Suite")
}
