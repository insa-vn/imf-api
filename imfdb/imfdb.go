package imfdb

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)


const (
	INSERT_IMG_URL_STATEMENT 	   = "INSERT INTO img_urls VALUES (null, ?, 1, md5(?));"
	INSERT_IMG_CHARACTER_STATEMENT = "INSERT INTO img_characters VALUES (last_insert_id(), ?, 0);"
	SELECT_RAND_IMG_STATEMENT = 
		`
		SELECT iu.url 
		FROM img_urls iu 
			INNER JOIN img_characters ic ON iu.id = ic.img_id
		WHERE ic.char_name = ?
		ORDER BY rand()
		LIMIT ?
		`
)


type CharacterImg struct {
	Url 	  string
	Character string
}


type ImfDB struct {
	config 	*mysql.Config
	db		*sql.DB
}


type Config struct {
	UserName	string	`json:"username"`
	Password	string	`json:"password"`
	Addr 		string	`json:"address"`
	DBName 		string	`json:"dbName"`
}


func (imfdb *ImfDB) Init(config Config) error {

	// Prepare data source in format :
	// username:password@tcp(endpointUrl:port)/databaseName

	imfdb.config 		= mysql.NewConfig()
	imfdb.config.User 	= config.UserName
	imfdb.config.Passwd = config.Password
	imfdb.config.Addr 	= config.Addr
	imfdb.config.DBName = config.DBName
	imfdb.config.Net 	= "tcp"
	
	var err error
	
	// Attention, this actually won't open any connection
	// Connection will be opened when needed (a ping for example)
	imfdb.db, err = sql.Open("mysql", imfdb.config.FormatDSN())
	if err != nil {
		return err
	}

	// Try to ping the server if the connection is established
	err = imfdb.db.Ping()
	if err != nil {
		return err
	}

	return nil
}


func (imfdb *ImfDB) Close() error {
	return imfdb.db.Close()
}


func (imfdb *ImfDB) AddImgUrls(imgs []CharacterImg) error {
	
	// Attention, at the time this code was written, the go's sql package 
	// as well as it's mysql driver don't support multiple statements
	// That sucks, but yeah...

	statement := &sql.Stmt{}
	var err error

	for _, img := range(imgs) {

		// Prepare SQL insert statement
		statement, err = imfdb.db.Prepare(INSERT_IMG_URL_STATEMENT)
		if err != nil {
			return err
		}
		// Execute statement
		_, err = statement.Exec(img.Url, img.Url)
		if err != nil {
			return err
		}

		// Prepare SQL insert statement
		statement, err = imfdb.db.Prepare(INSERT_IMG_CHARACTER_STATEMENT)
		if err != nil {
			return err
		}
		// Execute statement
		_, err = statement.Exec(img.Character)
		if err != nil {
			return err
		}
	}

	if statement != nil {
		statement.Close()
	}

	return nil
}


func (imfdb *ImfDB)	GetRandomImgs(character string, nbImgs int) ([]string, error) {

	// Query the DB for a number of random image urls corresponding to the selected character
	results, err := imfdb.db.Query(SELECT_RAND_IMG_STATEMENT, character, nbImgs)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	// The result format is weird, pass it to a string slice
	var imgUrls []string
	var url string
	for results.Next() {
		if err := results.Scan(&url); err != nil {
			return nil, err
		}
		imgUrls = append(imgUrls, url)
	}

	// If any error occured...
	if err := results.Err(); err != nil {
		return nil, err
	}

	return imgUrls, nil
}