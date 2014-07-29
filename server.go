package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"os"
)

func main() {
	fmt.Println("listening...")
	m := martini.Classic()
	m.Get("/", func() string {
		response, _ := getJSONResponse()
		return string(response)
	})
	m.Get("/albums/:id", getAlbumID)
	http.ListenAndServe(":"+os.Getenv("PORT"), m)
}

func getAlbumID(params martini.Params) string {
	return "Hello " + "id = " + params["id"] + " from func"
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
	fruits := make(map[string]int)
	fruits["Apple"] = 2
	fruits["Tomato"] = 3

	vagetables := make(map[string]int)
	vagetables["Pappers"] = 5

	d := Data{fruits, vagetables}
	p := Payload{d}
	return json.MarshalIndent(p, "", " ")
}
