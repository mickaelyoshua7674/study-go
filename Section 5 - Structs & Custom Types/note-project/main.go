package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "example.com/note/note"
)

func main() {
    title := getuserInput("Note title: ")
    content := getuserInput("Note content: ")

    userNote, err := note.New(title, content)
    if err != nil {
        panic(err)
    }

    userNote.Print()
    err = userNote.Save()
    if err!=nil {
        panic(err)
    }
    fmt.Println("Note saved.")
}

func getuserInput(prompt string) string {
    fmt.Print(prompt)
    /*
        Declaring a input output buffer reader, argument is the source you want to read.
        In this case is the Operation System Standard Input.
    */
    reader := bufio.NewReader(os.Stdin)
    /*
        ReadString will read the input in the command line, the argument is
        the delimiter to were stop reading.
        It will return the input including the delimiter.
    */
    value, err := reader.ReadString('\n')
    if err!=nil {
        return ""
    }

    value = strings.TrimSuffix(value, "\n")
    value = strings.TrimSuffix(value, "\r") //In Windows, sometimes the line break is '\r\n'
    return value
}