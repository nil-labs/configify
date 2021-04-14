package config

import (
	"io"
	"io/ioutil"
	"reflect"
	"gopkg.in/yaml.v2"
)

//Config file
type Config struct {
	value reflect.Value 
}

//FromReader produces a YAML
func FromReader(reader io.Reader) (*Config,error){
	bytes,err:= ioutil.ReadAll(reader)
	if err!=nil{
		return nil,err
	}
	var yml interface{}
	err=yaml.Unmarshal(bytes, &yml)
	if err !=nil {
		return nil, err
	}
	return &Config{value: reflect.ValueOf(yml)},nil
}