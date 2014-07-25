package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// RESTful
type Resource interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

type (
	GetNotSupported    struct{}
	PostNotSupported   struct{}
	PutNotSupported    struct{}
	DeleteNotSupported struct{}
)

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

func (PostNotSupported) Post(values url.Values) (int, interface{}) {
	return 405, ""
}

func (PutNotSupported) Put(values url.Values) (int, interface{}) {
	return 405, ""
}

func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
	return 405, ""
}

type API struct{}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		var data interface{}
		var code int

		method := request.Method // Get HTTP Method (string)
		request.ParseForm()      // Populates request.Form
		values := request.Form

		switch method {
		case "GET":
			code, data = resource.Get(values)
		case "POST":
			code, data = resource.Post(values)
		case "PUT":
			code, data = resource.Put(values)
		case "DELETE":
			code, data = resource.Delete(values)
		}

		content, err := json.Marshal(data)
		if err != nil {
			api.Abort(rw, 500)
		}

		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (api *API) AddResource(resource Resource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (api *API) Start() {
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

// JSON data
type Payload struct {
	Staff Data
}

type Data struct {
	Fruit   Fruits
	Vaggies Vagetables
}

type Fruits map[string]int
type Vagetables map[string]int

type HelloResource struct {
	PostNotSupported
	PutNotSupported
	DeleteNotSupported
}

func (HelloResource) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"hello": "world"}
	return 200, data
}

func main() {

	helloResource := new(HelloResource)

	var api = new(API)
	api.AddResource(helloResource, "/hello")
	api.Start()

	//http.HandleFunc("/", requestHandler)
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
