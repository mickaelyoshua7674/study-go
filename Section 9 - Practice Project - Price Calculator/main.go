package main
import (
    "fmt"
    "bufio"
    "os"
/*     "strings" */
    "strconv"
    "encoding/json"
)

const pricesFileName string = "prices.txt"
const taxIncludedPricesFileName string = "taxIncludedPrices.json"

func main() {
    taxRates := [4]float64{0, 0.07, 0.1, 0.15}
    result := make(map[string][]float64, len(taxRates))

/*     fmt.Print("Enter the values space separated: ")
    reader := bufio.NewReader(os.Stdin)
    txtValues, err := reader.ReadString('\n')
    errorHandler(err)
    txtValues = strings.TrimSuffix(txtValues, "\n")
    txtValues = strings.TrimSuffix(txtValues, "\r") */

    pricesFile, err := os.Open(pricesFileName)
    errorHandler(err)
    defer pricesFile.Close()

    scanner := bufio.NewScanner(pricesFile)
    prices := []float64{}
    for scanner.Scan() {
        price, err := strconv.ParseFloat(scanner.Text(), 64)
        errorHandler(err)
        prices = append(prices, price)
    }
    errorHandler(scanner.Err())

/*     pricesTxt := strings.Split(txtValues, " ") */

/*     prices, err := strArrayToFloatArray(pricesTxt...)
    errorHandler(err) */

    for _, taxRate := range taxRates {
        taxIncludedPrices := make([]float64, len(prices))
        for j, price := range prices {
            taxIncludedPrices[j] = price*(1+taxRate)
        }
        strKey:= strconv.FormatFloat(taxRate, 'f', 2, 64)
        result[strKey] = taxIncludedPrices
    }

    fmt.Println(result)
    json, err := json.Marshal(result)
    errorHandler(err)
    os.WriteFile(taxIncludedPricesFileName, json, 0644)
}

func errorHandler(err error) {
    if err!=nil {
        panic(err)
    }
}
/* func strArrayToFloatArray(strs ...string) ([]float64, error) {
    var err error
    ints := make([]float64, len(strs))
    for i, v := range strs {
        ints[i], err = strconv.ParseFloat(v, 64)
        if err!=nil {
            return []float64{}, err
        }
    }
    return ints, nil
} */