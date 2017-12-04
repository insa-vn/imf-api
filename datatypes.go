package main

import (
	"imf-api/imfdb"
)


type AppContext struct {
	db		*imfdb.ImfDB
	conf	Config
}


type Net struct {
	Port 		string 	`json:"port"`
	QueryPath 	string 	`json:"queryPath"`
	QueryMethod	string	`json:"queryMethod"`
}


type QueryParam struct {
	Character	string	`json:"character"`
	NbImgs		string	`json:"nbimgs"`
}


type Config struct {
	Db		imfdb.Config	`json:"database"`
	Net		Net				`json:"net"`
	Query	QueryParam		`json:"query"`
}


type RestResult struct {
	Meta	interface{}
	Data	[]string
}