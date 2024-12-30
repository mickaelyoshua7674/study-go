package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "example.com/note/note"
    "example.com/note/todo"
)
/*
Interface is a contract where a certain value (typically a struct) has a certain method.
Define the existence of a method, its name and its return.
*/
type saver interface {
    Save() error
}
type outputtable interface {
    saver
    Print()
}

func main() {
    printSomething(1)
    printSomething(1.5)
    printSomething("HI")
    printSomething(true)

    title := getuserInput("Note title: ")
    content := getuserInput("Note content: ")
    todoText := getuserInput("Todo text: ")

    userNote, err := note.New(title, content)
    errHandler(err)
    err = outputData(userNote)
    errHandler(err)

    userTodo, err := todo.New(todoText)
    errHandler(err)
    err = outputData(userTodo)
    errHandler(err)
}

/*
'any' is an alias for 'interface{}'
*/
func printSomething(value any) {
    // This will check if the value is an integer.
    // If it is, will return the value as a type 'int', so will no longer be of type 'any',
    // and will return a boolean value of 'true'.
    // If it is not, will return the null value for the type (in this caso 0 for 'int') and
    // a boolena value of 'false'
    typedVal, ok := value.(int)
    if ok {
        fmt.Println("It is a integer: ", typedVal)
    }

    // This switch will check the type of the variable, not its value.
    switch value.(type) {
    case int:
        fmt.Println("Integer")
    case float64:
        fmt.Println("Float 64")
    case string:
        fmt.Println("String")
    default:
        fmt.Println("Not integer, float64 or string.")
    }
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

func errHandler(e error) {
    if e != nil {
        panic(e)
    }
}

/*
This function will expect a 'data' that have a method called 'Save' and returns an error.
It could be any structure (typically) that match this conditions.
*/
func saveData(data saver) error {
    err := data.Save()
    if err!=nil {
        return err
    }
    fmt.Println("Note saved.")
    return nil
}

func outputData(data outputtable) error {
    data.Print()
    return saveData(data)
}