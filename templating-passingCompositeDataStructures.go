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

	sages := map[int]string{1:"Gandhi", 2:"MLK", 3:"Buddha", 4:"Jesus", 5:"Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}