package main

import ()

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
