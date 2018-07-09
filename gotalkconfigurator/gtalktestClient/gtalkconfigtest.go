package main

import (
	"fmt"
	"gotalk/gotalkconfigurator"
)

type Confstruct struct {
	TS      string  `xml:"testString"` //`json:"testString"` //`name:"testString"`
	TB      bool    `xml:"testBool"`   //`json:"testBool"`   //`name:"testBool"`
	TF      float64 `xml:"testFloat"`  //`json:"testFloat"`  //`name:"testFloat"`
	TestInt int
}

func main() {
	configstruct := new(Confstruct)
	//gotalkconfigurator.GetConfiguration(gotalkconfigurator.CUSTOM, configstruct, "configfile.conf")
	//gotalkconfigurator.GetConfiguration(gotalkconfigurator.XML, configstruct, "configfile.json")
	gotalkconfigurator.GetConfiguration(gotalkconfigurator.XML, configstruct, "configfile.xml")
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
