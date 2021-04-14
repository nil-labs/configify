package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

//Config file
type Config struct {
	data reflect.Value 
}

//FromReader produces a YAML
func FromReader(reader io.Reader) (*Config,error){
	bytes,err:= ioutil.ReadAll(reader)
	if err!=nil{
		return nil,err
	}
	yml:= make(map[interface{}]interface{})
	err=yaml.Unmarshal(bytes, &yml)
	if err !=nil {
		return nil, err
	}
	return &Config{data: reflect.ValueOf(yml)},nil
}


//String returns the string value based on path
// in case the path does not exists, or there is incompatible type nil will be returned
func (c *Config) String(path string) string{
	str,err:= c.StringE(path)
	if err!=nil{
		return ""
	}
	return str
}

//StringE returns the string value based on path
func (c *Config) StringE(path string) (string,error){
	v,err:= c.value(path)
	if err!=nil{
		return "", err
	}
	if v.Addr().IsValid(){
		switch v.Type(){
		case reflect.TypeOf(""):{
			return v.String(),nil
		}
	default:{
			return "", fmt.Errorf("Not implemented")
	}
		}
	}

	return "",nil
}

func (c *Config) value(path string) (reflect.Value,error){
	ref :=c.data
	for _, name:=range strings.Split(path, "."){
		ref=ref.FieldByName(name)
		if ref.IsZero(){
			return reflect.Value{}, fmt.Errorf("invalid path err")
		}
	}
	return ref, nil
}


