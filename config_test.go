package config_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nil-labs/config"
)

var _ = Describe("Config", func() {

	Describe("Constructing a Config object from reader containing yaml", func(){
		var yaml string
		//var cfg *config.Config
		var err error
		BeforeEach(func(){
			r:=strings.NewReader(yaml)
			_,err=config.FromReader(r)
		})
		Context("when the yaml is valid", func(){
			JustBeforeEach(func(){
				yaml=`
				---
				some:[]`
			})
			It("should succeed", func(){
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("when the yaml is invalid", func(){
			JustBeforeEach(func(){
				yaml=`
				---
				sasasas`
			})
			It("should fail", func(){
				Expect(err).To(HaveOccurred())
			})
		})
	})


})
