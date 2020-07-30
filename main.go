package main

import(
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Person struct {
	Name string
	Age int
}

func main(){

	p1 := Person{
		Name: "Wang Yu",
		Age: 33,
	}

	p2 := Person{
		Name: "Christopher Tregear",
		Age: 35,
	}

	p3 := Person{
		Name: "Ethan Tregear",
		Age: 1,
	}

	people := []Person{p1, p2, p3}

	err := tpl.Execute(os.Stdout, people)
	if err != nil{
		log.Fatalln(err)
	}

	
}

