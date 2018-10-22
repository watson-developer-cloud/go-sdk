// +build !integration

package naturallanguageclassifierv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNaturalLanguageClassifierV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NaturalLanguageClassifierV1 Suite")
}
