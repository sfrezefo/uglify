package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func parsej(f interface{}) {
	switch vv := f.(type) {
	case string:
		fmt.Print('"',vv'"')
	case float64:
		fmt.Print(vv)
	case []interface{}:
		fmt.Print("[")
		for i, u := range vv {
			if i != 0 {
				fmt.Print(",")
			}
			parsej(u)
		}
		fmt.Print("]")
	case map[string]interface{}:
		fmt.Print("{")
		l := 0
		for k, v := range vv {
			if l != 0 {
				fmt.Print(",")
			}
			fmt.Print('"',k,'":')
			parsej(v)
			l++
		}
		fmt.Print("}")
	case interface{}:
		fmt.Println(vv)
		fmt.Println("is an interface:")
	default:
		fmt.Println(vv)
		fmt.Println("is of a type I don't know how to handle")
	}
}

func main() {
	var g interface{}

	file, err := ioutil.ReadFile("test1.json")
	s := string(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	err1 := json.Unmarshal(file, &g)

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("unmarshalled : ", g)
	fmt.Println(" ")
	parsej(g)
	fmt.Println(" ")

}
