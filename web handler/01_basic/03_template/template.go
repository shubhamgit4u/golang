package main

import (
	"fmt"
	"os"
	"text/template"
)

type student struct {
	Name   string
	Rollno int
	Marks  int
}

//var index *template.Template
var fm = template.FuncMap{
	"grade": grade,
}

func grade() string {
	return "hello"
}
func main() {
	index := template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	shubham := student{Name: "shubham", Rollno: 1720, Marks: 85}
	shivani := student{Name: "shivani", Rollno: 1721, Marks: 86}
	suraj := student{Name: "rollno", Rollno: 1723, Marks: 87}
	data := []student{shubham, shivani, suraj}
	fmt.Println(shubham)
	fmt.Println("hello2")
	err := index.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}

}
