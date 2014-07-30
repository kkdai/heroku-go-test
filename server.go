package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("listening...")
	m := martini.Classic()
	m.Get("/", func() string {
		allFruit := db.GetAll()
		var response string
		for _, v := range allFruit {
			fmt.Printf(" id = %d name = %s num=%d \n", v.id, v.name, v.num)
			response = fmt.Sprintf("%s \n id = %d name = %s num=%d \n", response, v.id, v.name, v.num)
		}
		return response
	})
	m.Get("/fruit/:id", getFruitID)
	http.ListenAndServe(":"+os.Getenv("PORT"), m)
}

func getFruitID(params martini.Params) string {
	id, _ := strconv.Atoi(params["id"])
	fruit := db.Get(id)
	return "id = " + params["id"] + "name=" + fruit.name + "num=" + string(fruit.num)
}

func serveRest(res http.ResponseWriter, req *http.Request) {
	//Try json response
	response, err := getJSONResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(res, string(response))
}

func getJSONResponse() ([]byte, error) {
	fruits_data := DataItem{1, "Apple", 2}
	return json.MarshalIndent(fruits_data, "", " ")
}
