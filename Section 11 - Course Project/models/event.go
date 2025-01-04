package models
import (
    "time"
    "example.com/rest-api/db"
)

type Event struct {
    Id int64
    Name string `binding:"required"` //Struct tag to tell Gin that this field is required
    Description string `binding:"required"`
    Location string `binding:"required"`
    DateTime time.Time `binding:"required"`
    UserId int
}

var events = []Event{}

func (e Event) Save() error {
    query := `
    INSERT INTO events (name, description, location, dateTime, userId)
    VALUES (?, ?, ?, ?, ?)`
    stmt, err := db.DB.Prepare(query) // Prepare a SQL Statement and in some cases lead to a better performance.
    if err!=nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
    if err!=nil {
        return err
    }

    id, err := result.LastInsertId()
    e.Id = id
    return err
}

func GetAllEvents() []Event {
    return events
}