package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/tszalama/api-servicecatalog/tree/main/internal/config"
)

type Server struct {
	db *sql.DB
}

//Function retrieves the configuration and returns appropiate database connection string
func getConnString() string {

	config := config.GetConfig()

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		config.Server, config.Username, config.Password, config.Port, config.Database)

	return connString
}

//Function sets database connection
func InitDatabase() *Server {
	var err error

	connString := getConnString()

	log.Printf("Setting connection to db with configuration: %s \n", connString)

	server := &Server{}
	server.db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening connection: ", err.Error())
	}

	server.db.SetConnMaxLifetime(time.Minute * 4)

	return server
}

//Function verifies that connection is ok or generates a new one
func (s *Server) getConnection() {

	err := s.db.Ping()
	if err != nil {
		log.Fatal("Could not ping db: ", err.Error())
	}
	log.Println("Ping successful")
}
