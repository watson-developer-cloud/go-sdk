// +build !integration

package comparecomplyv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCompareComplyV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CompareComplyV1 Suite")
}
