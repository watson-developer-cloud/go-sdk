// +build !integration

package toneanalyzerv3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestToneAnalyzerV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ToneAnalyzerV3 Suite")
}
