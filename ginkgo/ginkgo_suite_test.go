package ginkgo

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Suite")
}

var _ = Describe("matcher_test", func() {
	It("should panic if []byte is not used", func() {
		var panicked = false
		defer func() {
			if v := recover(); v != nil {
				panicked = true
			}
			Expect(panicked).To(BeTrue())
		}()
		Expect("hello").To(EqualJson("world"))
	})

	It("should panic if []byte is not used", func() {
		var panicked = false
		defer func() {
			if v := recover(); v != nil {
				panicked = true
			}
			Expect(panicked).To(BeTrue())
		}()
		Expect([]byte("hello")).To(EqualJson("world"))
	})

	It("should return true if two JSONs are equivalent", func() {
		Expect([]byte(`{"hello":"world","a":"b"}`)).To(EqualJson([]byte(`{"a":"b","hello":"world"}`)))
	})

	It("should return true if two JSONs are different", func() {
		Expect([]byte(`{"hello":"world","a":"b"}`)).NotTo(EqualJson([]byte(`{"a":"b","hello":"World"}`)))
	})
})
