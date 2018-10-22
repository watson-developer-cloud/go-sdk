// +build !integration

package assistantv2_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAssistantV2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AssistantV2 Suite")
}
