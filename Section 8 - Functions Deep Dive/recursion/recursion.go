package main
import (
    "fmt"
    "errors"
)

func main() {
    result, err := factorial(5)
    if err!=nil {
        panic(err)
    }
    fmt.Println(result)
}

func factorial(number int) (int, error) {
    if number<0 {
        return 0, errors.New("invalid input, value must be greater then -1")
    }
    if number==1 || number==0 {
        return 1, nil
    }
    n, _ := factorial(number-1)
    return number*n, nil
}