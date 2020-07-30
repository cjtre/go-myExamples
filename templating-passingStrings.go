package main

import (
	"log"
	"os"
	"text/template"	
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

	// //First we parse our files
	// tpl, err := template.ParseFiles("tpl.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// then we execute. 
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "The total distance traversed by a ball rolling down an inclined plain or a body in freefall after starting from rest is directly proportional to the square of the time elapsed.")
	if err != nil {
		log.Fatalln(err)
	}
}