package main
import (
    "fmt"
    "example.com/structs/user"
)

func main() {
    userFirstName := getUserData("Please enter your first name: ")
    userLastName := getUserData("Please enter your last name: ")
    userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

    appUser, err := user.New(userFirstName, userLastName, userBirthDate)
    if err != nil {
        panic(err)
    }
    appUser.OutputUserDetails()
    appUser.ClearUserName()
    appUser.OutputUserDetails()
}

func getUserData(promptText string) (value string) {
    fmt.Print(promptText)
    fmt.Scanln(&value)
    return value
}