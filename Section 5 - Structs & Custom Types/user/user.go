package user
import (
    "fmt"
    "time"
    "errors"
)

/*
For export a structure is necessarily, not only export the structure itself (by starting the name with an upper case),
but also all the fields inside the structure.
It is possible to expose only the structure and not the fields, depends on the case.
*/
type User struct {
    firstName string
    lastName string
    birthDate string
    createdAt time.Time
}
/*
The constructor of a structure is only a convension, is not built in Golang.
The convention name is 'new[structure name]'. It is not mandatory to be that way.
*/
func New(fn, ln, bd string) (*User, error) {
    if fn=="" || ln=="" || bd=="" {
        return nil, errors.New("first name, last name and birthdate are required")
    }
    return &User{
        firstName: fn,
        lastName: ln,
        birthDate: bd,
        createdAt: time.Now(),
    }, nil

}

/*
This '(u User)' links the function to the structure, now it is possible for the function to 
access the declared structure.
That's called a method to the structure.
The '(u User)' is calles a Receiver argument.
*/
func (u User) OutputUserDetails() {
    fmt.Printf("\nFirst Name: %v\nLast Name: %v\nBirth Date: %v\n", u.firstName, u.lastName, u.birthDate)
}

/*
Here is used a User Pointer because the Receiver argument is just like any common argument, without the
pointer the value would be passed as a copy and the original value would stay the same.
*/
func (u *User) ClearUserName() {
    u.firstName = ""
    u.lastName = ""
}