//Demonstrates parsing into a struct
package main

import(
	"fmt"
	"encoding/json"
)

type App struct{
	Id string `json: "id"`
	Title string `json:"title"`
	Age string `json:"age,omitempty"` // if Age string is empty, then it will be left out.
}

func main(){
	data := []byte(`
{
	"id": "k34rAT4",
	"title": "My Awesome App",
	"age": "hello"
}
`)
		
var app App

err := json.Unmarshal(data, &app)
if err != nil{
	fmt.Println(err)
}
fmt.Println(app)

}
//As with all structs in Go, it’s important to remember that only fields with a capital first letter are visible to external programs like the JSON Marshaller.As with all structs in Go, it’s important to remember that only fields with a capital first letter are visible to external programs like the JSON Marshaller.