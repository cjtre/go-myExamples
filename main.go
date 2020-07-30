//The purpose of this programme is to show how we can read data in the form of json
//into a struct.
package main

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Countries struct {
	Countries []Country
}

type Country struct {
	Name string
	Capital string
}

func main(){
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var payload Countries{}

	err = json.Unmarshal([]byte(file), &payload)
	
	m := payload.(map[string]interface{})
	fmt.Println(m)

}
