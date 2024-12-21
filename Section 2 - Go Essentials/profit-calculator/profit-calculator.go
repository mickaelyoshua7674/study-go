package main
import (
    "fmt"
    "errors"
    "strings"
    "os"
)
const finantialsFileName = "finantials.txt"

func main() {
    revenue, errRevenue := getUserInput("Revenue: ")
    if errRevenue!=nil {
        panic(errRevenue)
    }
    expenses, errExpenses := getUserInput("Expenses: ")
    if errExpenses!=nil {
        panic(errExpenses)
    }
    taxRate, errTaxRate := getUserInput("Tax Rate: ")
    if errTaxRate!=nil {
        panic(errTaxRate)
    }

    ebt, profit, ratio := calculateFinantials(revenue, expenses, taxRate)
    fmt.Printf("\nEBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
    saveFinantials(ebt, profit, ratio)
}

func getUserInput(infoText string) (v float64, err error) {
    fmt.Print(infoText)
    fmt.Scan(&v)
    if v<=0 {
        errorMessage := fmt.Sprintf("%v value must be greater then 0\n", strings.Split(infoText, ":")[0])
        err = errors.New(errorMessage)
    }
    return v, err
}
func calculateFinantials(r, e, t float64) (ebt, profit, ratio float64) {
    ebt = r - e
    profit = ebt * (1-t/100)
    ratio = ebt / profit
    return ebt, profit, ratio
}
func saveFinantials(r, e, t float64) {
    txt := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", r, e, t)
    os.WriteFile(finantialsFileName, []byte(txt), 0644)
}