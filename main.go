//The purpose of this programme is to show how STRUCTURED json can be read into a struct.
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

	var baskets []FruitBasket
	json.Unmarshal([]byte(file), &baskets)
	
	fmt.Printf("baskets: %+v\n", baskets)
	



}