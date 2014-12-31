package main

import (
	"fmt"
)

// {“user_id”:”USER ID”,
//  “security_token”:”SECURITY TOKEN” ,
//  “api_version”:”1”,
//  “devices”:[{“mac_address”:“01-23-45-67-89-ab”, “status”:”online/offline”},
//  {“mac_address”:“01-23-45-67-89-ab”, “status”:”online/offline”}]}

// JSON data

type DataItem_in struct {
	name string `json:"name"`
	num  int    `json:"num"`
}
type DataItem struct {
	id   int    `json:"id"`
	name string `json:"name"`
	num  int    `json:"num"`
}

type FruitDB struct {
	m   map[int]*DataItem
	seq int
}

// The DB interface defines methods to manipulate the albums.
type DB interface {
	Get(id int) *DataItem
	GetAll() []*DataItem
	//Find(band, title string, year int) []*DataItem
	Add(a *DataItem) (int, error)
	Update(a *DataItem) error
	Delete(id int)
}

// The one and only database instance.
var db DB

func init() {
	db = &FruitDB{
		m: make(map[int]*DataItem),
	}
	// Fill the database
	db.Add(&DataItem{id: 1, name: "Apple", num: 5})
	db.Add(&DataItem{id: 2, name: "Banana", num: 3})
	db.Add(&DataItem{id: 3, name: "Lemon", num: 2})
}

// GetAll returns all albums from the database.
func (db *FruitDB) GetAll() []*DataItem {
	if len(db.m) == 0 {
		return nil
	}

	alldb := make([]*DataItem, len(db.m))
	i := 0
	for _, v := range db.m {
		alldb[i] = v
		i++
	}
	return alldb
}

func (db *FruitDB) Get(id int) *DataItem {
	fmt.Printf("id=%d \n", id)
	if len(db.m) == 0 {
		return nil
	}

	return db.m[id]
}

func (db *FruitDB) Add(a *DataItem) (int, error) {
	db.seq++
	a.id = db.seq
	db.m[a.id] = a
	return a.id, nil
}

func (db *FruitDB) Delete(id int) {
	delete(db.m, id)
}

func (db *FruitDB) Update(a *DataItem) error {
	db.m[a.id] = a
	return nil
}
