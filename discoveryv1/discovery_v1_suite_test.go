// +build !integration

package discoveryv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiscoveryV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DiscoveryV1 Suite")
}
