package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var index *template.Template

func main() {
	mux := httprouter.New()
	mux.GET("/", usersHandleFunc1)
	fmt.Println("hello main")
	mux.POST("/login", login)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index := template.Must(template.New("").ParseFiles("login.gohtml"))
	r.ParseForm()
	fmt.Printf("%v%T", r.Form.Get("n1"), r.Form.Get("n1"))
	num1, _ := strconv.Atoi(r.Form.Get("n1"))
	num2, _ := strconv.Atoi(r.Form.Get("n2"))
	index.ExecuteTemplate(w, "login.gohtml", num1+num2)
}
func usersHandleFunc1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hello user")
	r.ParseForm()
	r.Form.Set("n1", "8")
	index := template.Must(template.New("").ParseFiles("index.gohtml"))
	index.ExecuteTemplate(w, "index.gohtml", 3)
}
