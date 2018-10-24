// +build !integration

package naturallanguageunderstandingv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNaturalLanguageUnderstandingV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NaturalLanguageUnderstandingV1 Suite")
}
