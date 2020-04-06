// json ->byte ->make struct in slice-> struct
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	js := `[{"Name":"shubham","Age":24}]`
	fmt.Println(js)
	bc := []byte(js)
	//fmt.Println(bc)
	data := []person{}
	err := json.Unmarshal(bc, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
