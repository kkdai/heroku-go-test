package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("listening...")
	m := martini.Classic()

	//REST API
	m.Get("/", getAllUsers)
	m.Get("/users.json", getAllUsers)
	m.Get("/users.json/:id", getUserByID)
	m.Post("/users.json", addUser)

	// m.Get("/fruits/:id", getFruitID)
	// m.Post("/fruits", addFruit)
	// m.Put("/fruits/:id", updateFruit)
	// m.Delete("/fruits/:id", delFruits)
	// m.Get("/fruits", getFruits)
	// m.Get("/", getAll)
	m.Use(render.Renderer())
	http.ListenAndServe(":"+os.Getenv("PORT"), m)
}

func getUserByID(params martini.Params, r render.Render) {
	fmt.Println("get User By ID...")

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]interface{}{"error_code": "Request user ID not valid type, try integer."})
	}

	v := db.GetUserByID(id)
	r.JSON(200, map[string]interface{}{"id": v.id, "user_id": v.user_id, "user_name": v.user_name, "security_token": v.security_token, "api_version": v.api_version})
}

func getAllUsers(params martini.Params, r render.Render) {
	fmt.Println("get All User...")
	allUsers := db.GetAllUser()

	var totalUsers []map[string]interface{}
	for _, v := range allUsers {
		fmt.Printf(" id = %d userid = %s user_name = %s token=%s \n", v.id, v.user_id, v.user_name, v.security_token)
		perUser := map[string]interface{}{"id": v.id, "user_id": v.user_id, "user_name": v.user_name, "security_token": v.security_token, "api_version": v.api_version}
		totalUsers = append(totalUsers, perUser)
	}
	r.JSON(200, totalUsers)
}

func addUser(w http.ResponseWriter, r *http.Request, render render.Render) {
	fmt.Println("add User...")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read error")
	}
	fmt.Println(body)
	fmt.Println(reflect.TypeOf(body))
	fmt.Println(string(body))

	var dat map[string]interface{}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		//{“result”:”success/fail”, “type”:”new/existing”, “message”:”MESSAGE”}
		render.JSON(400, map[string]interface{}{"result": "fail", "type": "new", "message": "Input user format is not valid JSON format."})
		return
	}

	fmt.Println(dat)

	for key, value := range dat {
		fmt.Println("Key:", key, "Value:", value)
	}

	user_id_s, _ := dat["user_id"].(string)
	if db.GetUserByUserID(user_id_s) != nil {
		render.JSON(400, map[string]interface{}{"result": "fail", "type": "existing", "message": "Input user already existing."})
		return
	}

	user_name_s, _ := dat["user_name"].(string)
	security_token_s, _ := dat["security_token"].(string)
	api_version_s, _ := dat["api_version"].(string)

	user := UserData{id: 0, user_id: user_id_s, user_name: user_name_s, security_token: security_token_s, api_version: api_version_s}

	db.AddUser(&user)
	ret := map[string]interface{}{"result": "success", "type": "new", "message": "User add success."}
	render.JSON(400, ret)
	//ret_json, _ := json.Marshal(ret)
	//return http.StatusOK, string(ret_json)
}

// func getAll(params martini.Params) string {
// 	fmt.Println("getAll...")

// 	//List all items, currently list all fruits first.
// 	return getFruits(params)
// }

// func getFruits(params martini.Params) string {
// 	fmt.Println("getFruits...")
// 	allFruit := db.GetAll()
// 	var response string
// 	for _, v := range allFruit {
// 		fmt.Printf(" id = %d name = %s num=%d \n", v.id, v.name, v.num)
// 		response = fmt.Sprintf("%s \n id = %d name = %s num=%d \n", response, v.id, v.name, v.num)
// 	}
// 	return response
// }

// func delFruits(params martini.Params) string {
// 	fmt.Println("delFruits...")
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		return "error input\n"
// 	}
// 	db.Delete(id)
// 	return "Delete" + params["id"] + " \n"
// }

// func updateFruit(w http.ResponseWriter, r *http.Request, params martini.Params) (int, string) {
// 	fmt.Println("updateFruit...")
// 	id_i, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		id_i = 1
// 	}
// 	name_s, num_s := r.FormValue("name"), r.FormValue("num")
// 	num_i, err := strconv.Atoi(num_s)
// 	if err != nil {
// 		num_i = 0
// 	}

// 	fruit := DataItem{id: id_i, name: name_s, num: num_i}

// 	db.Update(&fruit)
// 	w.Header().Set("Location", fmt.Sprintf("/fruits/%d", id_i))
// 	return http.StatusOK, fmt.Sprintf("update id=%d to ==> name=%s, num=%d\n", id_i, name_s, num_i)
// }

// func addFruit(w http.ResponseWriter, r *http.Request) (int, string) {
// 	fmt.Println("addFruit...")

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("read error")
// 	}
// 	fmt.Println(body)
// 	fmt.Println(reflect.TypeOf(body))
// 	fmt.Println(string(body))

// 	var dat map[string]interface{}

// 	err = json.Unmarshal(body, &dat)
// 	if err != nil {
// 		fmt.Println("Unmarshal error")
// 	}
// 	fmt.Println(dat)

// 	for key, value := range dat {
// 		fmt.Println("Key:", key, "Value:", value)
// 	}
// 	num_s, _ := dat["num"].(string)
// 	name_s, _ := dat["name"].(string)
// 	num_i, _ := strconv.Atoi(num_s)

// 	fruit := DataItem{id: 0, name: name_s, num: num_i}

// 	db.Add(&fruit)

// 	//w.Header().Set("Location", fmt.Sprintf("/fruits/%d", id))
// 	return http.StatusOK, "work" //fmt.Sprintf("add id=%d, name=%s, num=%d\n", id, name_s, num_i)
// }
// func getFruitID(params martini.Params, r render.Render) {
// 	fmt.Println("getFruitID...")
// 	/*
// 	* curl -i "http://localhost:5000/fruit/2"
// 	 */
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		//return "error input\n"
// 	}
// 	fruit := db.Get(id)
// 	//fmt.Sprintf("id=%d, name=%s, total numer=%d\n", id, fruit.name, fruit.num)
// 	r.JSON(200, map[string]interface{}{"id": fruit.id, "name": fruit.name, "num": fruit.num})
// }

// func getJSONResponse() ([]byte, error) {
// 	fruits_data := DataItem{1, "Apple", 2}
// 	return json.MarshalIndent(fruits_data, "", " ")
// }
