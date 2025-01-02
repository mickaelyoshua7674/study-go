package main
import (
    "fmt"
)

type transformIntFn func(int) int

func main() {
    numbers := []int{1, 2, 3, 4}
    fmt.Println(numbers)
/*     dbNumbers := doubleNumbersNewSlice(&numbers)
    fmt.Println(dbNumbers) */
    transformNumbersDirect(&numbers, double)
    fmt.Println(numbers)
    transformNumbersDirect(&numbers, triple)
    fmt.Println(numbers)

}

/*
Passing function as a argument. The type is
func([types of parameters accepted]) [types of values returned]
*/
func transformNumbersDirect(numbers *[]int, transform transformIntFn) {
    for i, v := range *numbers {
        (*numbers)[i] = transform(v)
    }
}
/* func doubleNumbersNewSlice(numbers *[]int) []int {
    dbNumbers := make([]int, len(*numbers))
    for i, v := range *numbers {
        dbNumbers[i] = v*2
    }
    return dbNumbers
} */
func double (n int) int {
    return n*2
}
func triple(n int) int {
    return n*3
}