package arrays

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The arrays package", func() {
	Context("formatInput", func() {
		When("a string is passed", func() {
			It("should return an array of strings", func() {
				str := "Brave New World"
				strArr := formatInput(&str)
				Expect(len(*strArr)).To(Equal(13))
			})
		})
	})
})
