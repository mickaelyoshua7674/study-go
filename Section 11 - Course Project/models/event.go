package models

import (
    "database/sql"
    "time"
    "example.com/rest-api/db"
)

type Event struct {
    Id int64
    Name string `binding:"required"` //Struct tag to tell Gin that this field is required
    Description string `binding:"required"`
    Location string `binding:"required"`
    DateTime time.Time `binding:"required"`
    UserId int64
}

func (e *Event) Save() error {
    query := `
    INSERT INTO events (name, description, location, dateTime, userId)
    VALUES (?, ?, ?, ?, ?)`
    stmt, err := db.DB.Prepare(query) // Prepare a SQL Statement and in some cases lead to a better performance.
    if err!=nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId) // Tipically used when you want to change the database
    if err!=nil {
        return err
    }

    id, err := result.LastInsertId()
    e.Id = id
    return err
}

func GetAllEvents() ([]Event, error) {
    query := "SELECT * FROM events"
    rows, err := db.DB.Query(query) //'Query' is tipically used when you have a query that returns many rolls.
    if err!=nil {
        return nil, err
    }
    defer rows.Close()

    var events []Event
    for rows.Next() {
        var event Event
        err := rows.Scan(&event.Id,
            &event.Name,
            &event.Description,
            &event.Location,
            &event.DateTime,
            &event.UserId)
        if err!=nil {
            return nil, err
        }
        events = append(events, event)
    }

    return events, nil
}

func GetRowsNumber(r *sql.Rows) int {
    count := 0
    for r.Next() {
        count++
    }
    return count
}

func GetEventById(id int64) (*Event, error) {
    query := "SELECT * FROM events WHERE id=?"
    row := db.DB.QueryRow(query, id)

    var event Event
        err := row.Scan(&event.Id,
            &event.Name,
            &event.Description,
            &event.Location,
            &event.DateTime,
            &event.UserId)
        if err!=nil {
            return nil, err
        }
        return &event, nil
}

func (e Event) Update() error {
    query := `
    UPDATE events
    SET name=?, description=?, location=?, dateTime=?
    WHERE id=?`

    _, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.Id)
    return err
}

func (e Event) Delele() error {
    query := "DELETE FROM events WHERE id=?"

    _, err := db.DB.Exec(query, e.Id)
    return err
}