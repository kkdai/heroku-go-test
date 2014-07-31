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

	//REST API
	defer m.Post("/fruits", addFruit)
	m.Get("/fruit/:id", getFruitID)
	m.Get("/fruits", getFruits)
	m.Get("/", getAll)
	http.ListenAndServe(":"+os.Getenv("PORT"), m)
}

func getAll(params martini.Params) string {
	fmt.Println("getAll...")

	//List all items, currently list all fruits first.
	return getFruits(params)
}
func getFruits(params martini.Params) string {
	fmt.Println("getFruits...")
	allFruit := db.GetAll()
	var response string
	for _, v := range allFruit {
		fmt.Printf(" id = %d name = %s num=%d \n", v.id, v.name, v.num)
		response = fmt.Sprintf("%s \n id = %d name = %s num=%d \n", response, v.id, v.name, v.num)
	}
	return response
}

func addFruit(w http.ResponseWriter, r *http.Request) (int, string) {
	fmt.Println("addFruit...")
	name, num_s := r.FormValue("name"), r.FormValue("num")
	num, err := strconv.Atoi(num_s)
	if err != nil {
		num = 0
	}

	fruit := DataItem{id: 0, name: name, num: num}

	id, _ := db.Add(&fruit)
	w.Header().Set("Location", fmt.Sprintf("/albums/%d", id))
	return http.StatusOK, fmt.Sprintf("add id=%d, name=%s, num=%d\n", id, name, num)
}

func getFruitID(params martini.Params) string {
	fmt.Println("getFruitID...")
	/*
	* curl -i "http://localhost:5000/fruit/2"
	 */
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return "error input\n"
	}
	fruit := db.Get(id)
	return fmt.Sprintf("id=%d, name=%s, total numer=%d\n", id, fruit.name, fruit.num)

}

func getJSONResponse() ([]byte, error) {
	fruits_data := DataItem{1, "Apple", 2}
	return json.MarshalIndent(fruits_data, "", " ")
}
