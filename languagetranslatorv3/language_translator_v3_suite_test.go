// +build !integration

package languagetranslatorv3_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLanguageTranslatorV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LanguageTranslatorV3 Suite")
}
