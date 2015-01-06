package main

import (
	"fmt"
)

type UserStatus struct {
	id      int    `json:"id"`
	user_id string `json:"user_id"`
	devices string `json:"devices"`
	status  string `json:"status"`
}

type UserData struct {
	id             int    `json:"id"`
	user_id        string `json:"user_id"`
	user_name      string `json:"user_name"`
	security_token string `json:"security_token"`
	api_version    string `json:"api_version"`
}

type UserCertification struct {
	id       int    `json:"id"`
	user_id  string `json:"user_id"`
	platform string `json:"platform"` //iOS/Android
	receipt  string `json:"receipt"`  //JSON_RECEIPT
}

type ServerDB struct {
	user_db    map[int]*UserData
	seq_user   int
	cert_db    map[int]*UserCertification
	seq_cert   int
	status_db  map[int]*UserStatus
	seq_status int
	//seq     int
}

// The DB interface defines methods to manipulate the albums.
type DB interface {
	//User Data
	AddUser(a *UserData) (int, error)

	GetAllUser() []*UserData
	GetUserByID(id int) *UserData
	GetUserByUserID(user_id string) *UserData

	PrintUsers(total_users []*UserData)

	//Certification Data
	AddCert(a *UserCertification) (int, error)
	// Get(id int) *DataItem
	// GetAll() []*DataItem
	// //Find(band, title string, year int) []*DataItem
	// Add(a *DataItem) (int, error)
	// Update(a *DataItem) error
	// Delete(id int)
}

// The one and only database instance.
var db DB

func init() {
	db = &ServerDB{
		user_db:   make(map[int]*UserData),
		cert_db:   make(map[int]*UserCertification),
		status_db: make(map[int]*UserStatus),
	}
	//m: make(map[int]*DataItem),

	// Fill the database
	fmt.Println("Init DB")
	db.AddUser(&UserData{id: 1, user_id: "test1", user_name: "name1", security_token: "token1", api_version: "api1"})
	db.PrintUsers(db.GetAllUser())
	db.AddUser(&UserData{id: 2, user_id: "test2", user_name: "name2", security_token: "token2", api_version: "api1"})
	db.PrintUsers(db.GetAllUser())
	db.AddUser(&UserData{id: 3, user_id: "test3", user_name: "name3", security_token: "token3", api_version: "api1"})
	db.PrintUsers(db.GetAllUser())
}

func (db *ServerDB) PrintUsers(total_users []*UserData) {
	for _, v := range total_users {
		fmt.Printf("User id=%d, user_id=%s, user_name=%s, security_token=%s, api_version=%s \n", v.id, v.user_id, v.user_name, v.security_token, v.api_version)
	}
}

func (db *ServerDB) AddCert(a *UserCertification) (int, error) {
	db.seq_cert++
	a.id = db.seq_cert
	db.cert_db[a.id] = a
	return a.id, nil
}

func (db *ServerDB) AddUser(a *UserData) (int, error) {
	db.seq_user++
	a.id = db.seq_user
	db.user_db[a.id] = a
	return a.id, nil
}

func (db *ServerDB) GetAllUser() []*UserData {
	if len(db.user_db) == 0 {
		return nil
	}

	alldb := make([]*UserData, len(db.user_db))
	i := 0
	for _, v := range db.user_db {
		alldb[i] = v
		i++
	}
	return alldb
}

func (db *ServerDB) GetUserByID(id int) *UserData {
	fmt.Printf("id=%d \n", id)
	if len(db.user_db) == 0 {
		return nil
	}

	return db.user_db[id]
}

func (db *ServerDB) GetUserByUserID(user_id string) *UserData {
	fmt.Printf("user_id=%s \n", user_id)

	if len(db.user_db) == 0 {
		return nil
	}

	for _, v := range db.user_db {
		if v.user_id == user_id {
			return v
		}
	}

	return nil
}

// GetAll returns all albums from the database.
/*
func (db *ServerDB) GetAll() []*DataItem {
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

func (db *ServerDB) Get(id int) *DataItem {
	fmt.Printf("id=%d \n", id)
	if len(db.m) == 0 {
		return nil
	}

	return db.m[id]
}

func (db *ServerDB) Add(a *DataItem) (int, error) {
	db.seq++
	a.id = db.seq
	db.m[a.id] = a
	return a.id, nil
}

func (db *ServerDB) Delete(id int) {
	delete(db.m, id)
}

func (db *ServerDB) Update(a *DataItem) error {
	db.m[a.id] = a
	return nil
}
*/
