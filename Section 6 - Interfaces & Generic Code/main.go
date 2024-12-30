package main
import (
    "fmt"
)

func main() {
    result1 := add(1,2)
    fmt.Println(result1)

    result2 := add(5.5,6.5)
    fmt.Println(result2)

    result3 := add("Mickael ","Yoshua")
    fmt.Println(result3)
}

func add[T int|float64|string](a, b T) T {
    return a+b
}