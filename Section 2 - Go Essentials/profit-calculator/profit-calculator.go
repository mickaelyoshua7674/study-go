package main
import "fmt"

func main() {
    revenue := getUserInput("Revenue: ")
    expenses := getUserInput("Expenses: ")
    taxRate := getUserInput("Tax Rate: ")

    ebt, profit, ratio := calculateFinantials(revenue, expenses, taxRate)

    fmt.Printf("\nEBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

func getUserInput(infoText string) (v float64) {
    fmt.Print(infoText)
    fmt.Scan(&v)
    return v
}
func calculateFinantials(r, e, t float64) (ebt, profit, ratio float64) {
    ebt = r - e
    profit = ebt * (1-t/100)
    ratio = ebt / profit
    return ebt, profit, ratio
}