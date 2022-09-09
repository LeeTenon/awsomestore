package jsonx

import (
	"awsomestore/common/configmanager"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("../../config/config.yaml")
	if err != nil {
		panic("read file fail")
	}

	result:=&configmanager.BaseConfig{}
	err = Unmarshal(data,result)
	if err != nil {
		panic(fmt.Sprintln("unmarshal fail: ", err))
	}

	js,_:=json.MarshalIndent(result,"","\t")
	fmt.Print(string(js))
}

