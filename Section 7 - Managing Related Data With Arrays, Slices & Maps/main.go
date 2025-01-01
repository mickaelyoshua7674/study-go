package main
import (
    "fmt"
)

type floatMap map[string]float64

func main() {
    /*
    'make' will pre allocate the memory when creating a slice, map,
    The second argument is the size of the slice created and the third is the capacity,
    that way (in this example), it is possible to directly give values to the first 2 memory
    spaces allocated and append to the other 3 without Go having to allocate more memory.
    */
    userNames := make([]string, 2, 5)
    userNames[0] = "Julie"
    userNames[1] = "Mark"
    fmt.Println(userNames)
    for i, v := range userNames {
        fmt.Println(i, v)
    }
    for i:=0; i<len(userNames); i++ {
        fmt.Println(userNames[i])
    }

    // For maps it is not possible to allocate more then the expected length (the capacity).
    courseRatings := make(floatMap, 3)
    courseRatings["go"] = 4.7
    courseRatings["react"] = 4.8
    courseRatings["angular"] = 4.7
    fmt.Println(courseRatings)
    for k, v := range courseRatings {
        fmt.Println(k, v)
    }
}