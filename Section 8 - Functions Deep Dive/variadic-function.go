package main
import (
    "fmt"
)

func main() {
    fmt.Println(sum(1,2,3,4))
}

/*
Variadic functions are functions that accept any number of parameters.
The variable will always be a slice of the defined value type after the '...'.
*/
func sum(numbers ...int) int {
    sum := 0
    for _, val := range numbers {
        sum += val
    }
    return sum
}