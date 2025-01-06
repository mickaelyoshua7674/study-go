package models
import (
    "example.com/rest-api/db"
    "example.com/rest-api/utils"
)

type User struct {
    Id       int64
    Email    string `binding:"required"`
    Password string `binding:"required"`
}

func (u User) Save() error {
    query := "INSERT INTO users(email, password) VALUES (?,?)"
    hashedPassword, err := utils.HashPassword(u.Password)
    if err!=nil {
        return err
    }

    result, err := db.DB.Exec(query, u.Email, hashedPassword)
    if err!= nil{
        return err
    }

    userId, err := result.LastInsertId()
    if err!= nil{
        return err
    }

    u.Id = userId
    return nil
}