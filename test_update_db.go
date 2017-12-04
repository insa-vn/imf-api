package main

import (
	"imfdb"
	_ "googlesearch"
	"fmt"
)


const (
	API_KEY		  = "AIzaSyBdKG_M0zZ3BAv5NTyjXoo9015D-svGs0o"
	SEARCH_ENGINE = "014331277232571370493:akt9l5xjkye"
)


func check(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	db := &imfdb.ImfDB {}

	err := db.Init(imfdb.Config {
		UserName : "admin",
		Password : "11235813",
		Addr 	 : "insa-vn.cluqt19clm04.eu-west-2.rds.amazonaws.com:3306", 
		DBName   : "ImagesFight",
	})
	check(err)
	defer db.Close()

	// searcher, err := googlesearch.New(API_KEY, SEARCH_ENGINE, true)
	// check(err)
	
	// var imgLinks []imfdb.CharacterImg

	// results, err := searcher.Search("Jon Snow", 10, 1)
	// check(err)
	
	// for _, item := range(results.Items) {
	// 	imgLinks = append(imgLinks, imfdb.CharacterImg { Url: item.Link, Character: "Jon Snow" })
	// }
	// err = db.AddImgUrls(imgLinks)
	// check(err)

	// results, err = searcher.Search("Tyrion", 10, 1)
	// check(err)
	
	// imgLinks = nil
	// for _, item := range(results.Items) {
	// 	imgLinks = append(imgLinks, imfdb.CharacterImg { Url: item.Link, Character: "Tyrion" })
	// }
	// err = db.AddImgUrls(imgLinks)
	// check(err)

	fmt.Println("Hello")
}