package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	tmpl, err := template.ParseFiles("hello.tmpl")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Elias yemate",
	}
	err = tmpl.Execute(os.Stdin, user)

	if err != nil {
		panic(err)
	}
}
