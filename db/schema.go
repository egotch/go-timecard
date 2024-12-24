package db

import (
	"log"
	"os"
)

// "database/sql"
// "os"

// "github.com/mattn/go-sqlite3"

func InitDb(Logger *log.Logger) {

  if os.File
  f, err := os.OpenFile("sqlite3.db", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
  f, err := os.Create("sqlite.db")

  if err != nil {
    Logger.Panicf("Unable to init db: %v", err) 
  }

  f.Close()
  Logger.Println("database initiated")

}

