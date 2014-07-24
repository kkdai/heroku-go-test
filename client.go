package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Staff Data
}

type Data struct {
	Fruit   Fruits
	Vaggies Vagetables
}

type Fruits map[string]int
type Vagetables map[string]int

func main() {
	url := "http://localhost:5000"
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Staff.Fruit, "\n", p.Staff.Vaggies)

}
