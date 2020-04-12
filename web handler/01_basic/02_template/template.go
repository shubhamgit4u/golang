package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("index.html")
	fmt.Println(err)
	err = file.Execute(os.Stdout, nil)

}
