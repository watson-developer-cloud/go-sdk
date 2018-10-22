// +build !integration

package personalityinsightsv3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPersonalityInsightsV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PersonalityInsightsV3 Suite")
}
