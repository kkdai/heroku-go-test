package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("listening...")
	m := martini.Classic()

	//REST API
	m.Get("/fruits/:id", getFruitID)
	m.Post("/fruits", addFruit)
	m.Put("/fruits/:id", updateFruit)
	m.Delete("/fruits/:id", delFruits)
	m.Get("/fruits", getFruits)
	m.Get("/", getAll)
	m.Use(render.Renderer())
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

func delFruits(params martini.Params) string {
	fmt.Println("delFruits...")
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return "error input\n"
	}
	db.Delete(id)
	return "Delete" + params["id"] + " \n"
}

func updateFruit(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
	fmt.Println("updateFruit...")
	id_i, err := strconv.Atoi(params["id"])
	if err != nil {
		id_i = 1
	}
	name_s, num_s := r.FormValue("name"), r.FormValue("num")
	num_i, err := strconv.Atoi(num_s)
	if err != nil {
		num_i = 0
	}

	fruit := DataItem{id: id_i, name: name_s, num: num_i}

	db.Update(&fruit)
	w.Header().Set("Location", fmt.Sprintf("/fruits/%d", id_i))
	return http.StatusOK, fmt.Sprintf("update id=%d to ==> name=%s, num=%d\n", id_i, name_s, num_i)
}

func addFruit(w http.ResponseWriter, r *http.Request) (int, string) {
	fmt.Println("addFruit...")

	fmt.Printf("body =%s \n", r.Body)
	decoder := json.NewDecoder(r.Body)
	var t DataItem_in
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println("error happen")
	}
	// name_s, num_s := r.FormValue("name"), r.FormValue("num")
	// fmt.Println(name_s)
	// fmt.Println(num_s)
	fmt.Printf("Post data name=%s, total numer=%d \n", t.name, t.num)
	// num_i, err := strconv.Atoi(num_s)
	// if err != nil {
	// 	num_i = 0
	// }

	fruit := DataItem{id: 0, name: t.name, num: t.num}

	id, _ := db.Add(&fruit)
	w.Header().Set("Location", fmt.Sprintf("/fruits/%d", id))
	return http.StatusOK, fmt.Sprintf("add id=%d, name=%s, num=%d\n", id, t.name, t.num)
}

func getFruitID(params martini.Params, r render.Render) {
	fmt.Println("getFruitID...")
	/*
	* curl -i "http://localhost:5000/fruit/2"
	 */
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//return "error input\n"
	}
	fruit := db.Get(id)
	//fmt.Sprintf("id=%d, name=%s, total numer=%d\n", id, fruit.name, fruit.num)
	r.JSON(200, map[string]interface{}{"id": fruit.id, "name": fruit.name, "num": fruit.num})
}

func getJSONResponse() ([]byte, error) {
	fruits_data := DataItem{1, "Apple", 2}
	return json.MarshalIndent(fruits_data, "", " ")
}
