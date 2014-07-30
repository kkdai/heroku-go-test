package main

import ()

// JSON data
type DataItem struct {
	id   int
	name string
	num  int
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
	//Update(a *DataItem) error
	//Delete(id int)
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
	if len(db.m) == 0 {
		return nil
	}

	return db.m[id]
}

func (db *FruitDB) Add(a *DataItem) (int, error) {
	db.seq++
	db.m[db.seq] = a
	return db.seq, nil
}