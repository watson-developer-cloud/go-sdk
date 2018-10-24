// +build !integration

package assistantv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAssistantV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AssistantV1 Suite")
}
