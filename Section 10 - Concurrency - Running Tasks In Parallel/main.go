package main
import (
    "fmt"
    "time"
)

func main() {
    /*
    Channels are pipes of data for comunication between goroutines.
    */
    done := make(chan bool)

    go greet("Nice to meet you!", done)
    go greet("How are you?", done)
    go slowGreet("How ... are ... you ...?", done)
    go greet("I hope you're liking the course!", done)

    for range done {}
}

func greet(phrase string, doneChan chan bool) {
    fmt.Println("Hello!", phrase)
    doneChan <- true
}
func slowGreet(phrase string, doneChan chan bool) {
    time.Sleep(3*time.Second)
    greet(phrase, doneChan)
    doneChan <- true
    close(doneChan)
}