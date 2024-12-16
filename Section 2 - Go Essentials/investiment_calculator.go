package main

import (
    "fmt"
    "math"
)

func main() {
    //var investmentAmount float64 = 1000
    //var years float64 = 10

    //var investmentAmount, years float64 = 1000, 10
    //investmentAmount, years, expectedReturnRate := 1000., 10., 5.5
    /*
        ':=' infer the type to the variable. Create and infer the type.
        That's the recomended short way to assign a variable.
    */

    const inflationRate = 2.5

    var investmentAmount, years, expectedReturnRate float64
    fmt.Print("Investment Amount: ")
    fmt.Scan(&investmentAmount)
    /*
        the Scan function works with pointers
    */

    fmt.Print("Expected Return Rate: ")
    fmt.Scan(&expectedReturnRate)

    fmt.Print("Years: ")
    fmt.Scan(&years)



    futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
    fmt.Println(futureValue)
    fmt.Println(futureRealValue)
}