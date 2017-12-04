package main

import (
	"imf-api/imfdb"

	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"log"
)


func (appCtx *AppContext) Config(configFile string) error {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	} 

	var conf Config
	if json.Unmarshal(file, &conf) != nil {
		return err
	}

	appCtx.conf = conf
	return nil
}


func main() {

	// Create an application context
	appCtx := &AppContext{}

	// Create database connection
	appCtx.db = &imfdb.ImfDB {}

	// Read config file
	err := appCtx.Config("./resources/config.json")
	check(err)
	
	// Configure database connection
	err = appCtx.db.Init(appCtx.conf.Db)
	check(err)
	defer appCtx.db.Close()

	// Create a router
	router := mux.NewRouter()
	router.HandleFunc(appCtx.conf.Net.QueryPath, appCtx.GetImgUrls).Methods(appCtx.conf.Net.QueryMethod)

    log.Fatal(http.ListenAndServe(appCtx.conf.Net.Port, router))
}

func (app *AppContext) GetImgUrls(w http.ResponseWriter, r *http.Request) {

	queries := r.URL.Query()

	// Check params validities
	character := queries.Get(app.conf.Query.Character)
	nbImgs, convErr := strconv.Atoi(queries.Get(app.conf.Query.NbImgs))

	// In case nbimgs is not number, return only one image url
	if convErr != nil {
		nbImgs = 1
	}

	w.Header().Add("Content-Type", "application/json")

	// Query!!!
	imgUrls, _ := app.db.GetRandomImgs(character, nbImgs)
	jsonRet := &RestResult{ Data : imgUrls }
	json.NewEncoder(w).Encode(jsonRet)
}


func check(err error) {
	if err != nil {
		panic(err)
	}
}