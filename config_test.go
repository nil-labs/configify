package config_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nil-labs/config"
)

var _ = Describe("Config", func() {

		var yaml string
		var cfg *config.Config
		var err error
		BeforeEach(func(){
			r:=strings.NewReader(yaml)
			cfg,err=config.FromReader(r)
		})
	Describe("Constructing a Config object from reader containing yaml", func(){
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
	Describe("Reading string property", func(){
		
		JustBeforeEach(func(){
			yaml=`
			---
			lorem:
			  ipsum: sit amet
			`
		})
		It("should succeed",func(){
			Expect(cfg.String("lorem.ipsum")).To(Equal("sit amet"))
			Expect(cfg.StringE("lorem.ipsum")).ToNot(HaveOccurred())
		})
 
	})

})
