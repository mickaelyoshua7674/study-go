package main
import (
    "fmt"
)

func main() {
    numbers := []int{1, 2, 3}
    double := createTransformer(2)
    triple := createTransformer(3)

    /*
    Anonymous function. defining a function in place, as an argument.
    It's called anonymous because it doesn't have a name.
    */
    transformed := transformNumbers(&numbers, func(number int) int {
        return number*2
    })
    doubled := transformNumbers(&numbers, double)
    tripled := transformNumbers(&numbers, triple)

    fmt.Println(transformed)
    fmt.Println(doubled, tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
    dbNumbers := make([]int, len(*numbers))
    for i, v := range *numbers {
        dbNumbers[i] = transform(v)
    }
    return dbNumbers
}
func createTransformer(multiplier int) func(int) int {
    return func(number int) int {
        return number * multiplier
    }
}