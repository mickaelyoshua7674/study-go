package db

import (
    "fmt"
    "database/sql"
    _"modernc.org/sqlite" // '_' indicates that this package will be used, but not directly
    // Saw on Reddit and Medium that this package works better then 'github.com/mattn/go-sqlite3'
)

const dataBaseFileName string = "api.db"
var DB *sql.DB 

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", dataBaseFileName)

    if err!=nil {
        fmt.Println(err)
        panic("could not connect to database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5) //Make shore that will have at least 5 connections open always.

    createTables()
}

func createTables() {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    )`
    _, err := DB.Exec(createUsersTable)
    if err!=nil {
        fmt.Println(err)
        panic("could not create users table")
    }


    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        userId INTEGER,
        FOREIGN KEY(userId) REFERENCES users(id)
    )`

    _, err = DB.Exec(createEventsTable)
    if err!=nil {
        fmt.Println(err)
        panic("could not create events table")
    }
}