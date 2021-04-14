package config_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nil-labs/config"
)

var _ = Describe("Config", func() {

	Describe("Constructing a Config object from reader containing yaml", func(){
		Context("when the yaml is valid", func(){
			It("should succeed", func(){
				yaml:=`
				---
				some:[]`
				r:=strings.NewReader(yaml)
				Expect(config.FromReader(r)).To(Succeed())
			})
		})
	})


})
