package main
import (
    "fmt"
)

func main() {
    /*
    This way is created a dynamic array.
    When the size increase, Go automatically will
    create a new array with a new size to contain all the elements
    and delete the previous array.
    */
    prices := []float64{10.99, 8.99}
    fmt.Println(prices[0:1])
    prices[1] = 9.99

    updatedPrices := append(prices, 5.99)
    fmt.Println(updatedPrices, prices)
}

func slicesMain() {
    var productNames [4]string = [4]string{"A Book"}
    productNames[2] = "A Carpet"
    fmt.Println(productNames)

    prices := [4]float64{10.99, 9.99, 45.99, 20.}
    fmt.Println(prices)
    /*
    Start index is included, end index is excluded.
    */
    fmt.Println(prices[1:3])
    fmt.Println(prices[1:])
    fmt.Println(prices[:3])

    /*
    Slices are a window of the array, so when changing the value
    of a slice, the same value in the array is also changed.
    That means that when creating a slice, the data is not copyed,
    it's like an pointer to the original array.
    */
    featuredPrices := prices[1:]
    featuredPrices[0] = 199.99
    fmt.Println(prices)

    /*
    len is the number of elements.
    cap is the capacity of an array/slice/map.
    */
    highlightedPrices := featuredPrices[:1]
    fmt.Println(len(prices), cap(prices))
    fmt.Println(len(featuredPrices), cap(featuredPrices))

    /*
    Slices can always resize more to the right, but not to the left.
    'featuredPrices' has lenght and capacity of 3 because it was sliced from the second
    element to the end.
    'highlightedPrices'has lenght of 1 but capacity 3 because it was sliced from the 
    first element to the second not included, but it still can be expanded to the right.

    Internally Go memorized that exists more content to the right of that slice.
    */
    fmt.Println(highlightedPrices)
    fmt.Println(len(highlightedPrices), cap(highlightedPrices))
    highlightedPrices = highlightedPrices[:3]
    fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
}