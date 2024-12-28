package main
import (
    "fmt"
)

type customString string
func (txt customString) log() {
    fmt.Println(txt)
}


func main() {
    var name customString = "Mickael"
    name.log()
}