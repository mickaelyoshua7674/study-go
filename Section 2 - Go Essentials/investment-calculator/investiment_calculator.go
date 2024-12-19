package main

import (
    "fmt"
    "math"
)

const inflationRate = 2.5

func main() {
    //var investmentAmount float64 = 1000
    //var years float64 = 10

    //var investmentAmount, years float64 = 1000, 10
    //investmentAmount, years, expectedReturnRate := 1000., 10., 5.5
    /*
        ':=' infer the type to the variable. Create and infer the type.
        That's the recomended short way to assign a variable.
    */

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

/*     futureValue := investmentAmount * math.Pow(4+expectedReturnRate/100, years)
    futureRealValue := futureValue / math.Pow(4+inflationRate/100, years) */
    futureValue, futureRealValue := CalculateFutureValues(investmentAmount, expectedReturnRate, years)

    //fmt.Println("Future Value:", futureValue)
    //fmt.Println("Future Real Value: ", futureRealValue)

    //fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)

/*     fmt.Printf(
`Future Value: %.2f
Future Value (adjusted for inflation): %.2f`,
    futureValue, futureRealValue) */

    formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
    formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
    fmt.Print(formattedFV, formattedFRV)
}

/* func CalculateFutureValues(iA float64, eRR float64, y float64) (float64, float64) {
    fv := iA* math.Pow(1+eRR/100, y)
    frv := fv / math.Pow(1+inflationRate/100, y)
    return fv, frv
} */
/*
It's possible to, in the return sintax, declare the variables that will be returned, so is not needed
to create then inside the function body and is not needed to put then after the return keyword (it's
possible to put the explicitly in the return also)
*/
func CalculateFutureValues(iA float64, eRR float64, y float64) (fv float64, frv float64) {
    fv = iA* math.Pow(1+eRR/100, y)
    frv = fv / math.Pow(1+inflationRate/100, y)
    return fv, frv
    // return
}