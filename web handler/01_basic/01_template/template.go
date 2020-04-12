package main

import "fmt"

func main() {
	var i = "sharma"
	z := `<html>
	 <head>
	  <tittle>hello` + i + `</title>
	 <body> 
	 <p> hello ` + ` ` + i +
		`</body> 
	 </head>
	 <html>`
	fmt.Printf("%T %v", z, z)
}
