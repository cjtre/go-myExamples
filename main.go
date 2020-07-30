//The purpose of this programme is to show how UNSTRUCTURED* data can be decoded into a map.
package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type FruitBasket struct {
    Name    string `json: "fname"`
    Fruit   []string
    Id      int64 `json:"ref"`
}

func main(){
	
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}

	var result []map[string] interface{}

	err = json.Unmarshal([]byte(file), &result)
	if err != nil {
		fmt.Println(err)
	}

	for key, value := range result {
		fmt.Println(key, value)
	}

}