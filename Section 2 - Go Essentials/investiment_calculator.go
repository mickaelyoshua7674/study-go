package main

import (
    "fmt"
    "math"
)

func main() {
    //var investmentAmount float64 = 1000
    //var years float64 = 10

    //var investmentAmount, years float64 = 1000, 10

    investmentAmount, years, expectedReturnRate := 1000., 10., 5.5
    /*
        ':=' infer the type to the variable. Create and infer the type.
        That's the recomended short way to assign a variable.
    */

    futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    fmt.Println(futureValue)
}