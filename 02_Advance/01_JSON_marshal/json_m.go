//JSON -> JAVASCRIPT object notation
// like we all speak english so we have common format a,b,c
// in data world we have ccommon language to spread data i.e json
//XMl and json and more other but common one is json and most popular
// advantage - common language / easy to retrieve data and manuplate it // transfer data// and most important
// it provide perticular data like we fill form for pu fees when we click submit it passes json data to backend
// when at the end you click download form , you see your information that comes from backend to frontend as json
//formate {Name:"shubham",Age:22,Marks:55} //first letter should be capital//
//in development side we use structure and in passing data to frontend we use json , so we have to convert struct to json and
//when we take data back as json we again convert it in struct
// convertion of struct to json -> json Marshel // struct -> slice -> json
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	shubham := person{"shubham", 24}
	akash := person{"akash", 22}
	var sl = []person{shubham, akash}
	fmt.Println(sl)
	bs, err := json.Marshal(sl)

	fmt.Println(string(bs))
	fmt.Println(err)

}
