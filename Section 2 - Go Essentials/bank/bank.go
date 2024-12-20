package main
import (
    "errors"
    "fmt"
    "os"
    "strconv"
)

const balanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
    data, errReadFile := os.ReadFile(balanceFile)
    if errReadFile != nil {
        return 1000., errors.New("failed to find balance file")
    }

    balance, errParseFloat := strconv.ParseFloat(string(data), 64)
    if errParseFloat != nil {
        return 1000., errors.New("failed to parse stored balance value")
    }
    return balance, nil
}

func writeBalanceToFile(balance float64) {
    balanceText := fmt.Sprint(balance)
    os.WriteFile(balanceFile, []byte(balanceText), 0644)
    // 0644 is a file permission notation where the owner of the file will be
    // able to write and read the file, whereas other users can only read it.
    // https://www.redhat.com/en/blog/linux-file-permissions-explained
}

func main() {
    accountBalance, err := getBalanceFromFile()
    if err != nil {
        fmt.Println("ERROR")
        fmt.Println(err)
        fmt.Println("----------")
        panic("can't continue")
        // panic() function will stop the program execution in a way that looks like it crashed
    }

    fmt.Printf("\n*** Welcome to Go Bank! ***\n\n")
    for {
        fmt.Println("What do you want to do?")
        fmt.Printf("1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit\n\n")

        var choice int
        fmt.Print("Your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            fmt.Printf("Your balance is %.2f\n\n", accountBalance)

        case 2:
            fmt.Print("Deposit amount: ")
            var depositAmount float64
            fmt.Scan(&depositAmount)
            if depositAmount <= 0 {
                fmt.Printf("Invalid amount. Must be greater than 0.\n\n")
                continue
            }
            accountBalance += depositAmount
            fmt.Printf("Your new balance is %.2f\n\n", accountBalance)
            writeBalanceToFile(accountBalance)

        case 3:
            fmt.Print("Withdrawal amount: ")
            var withdraw float64
            fmt.Scan(&withdraw)
            if withdraw <= 0 || withdraw > accountBalance {
                fmt.Printf("Invalid withdraw. Must be greater than 0 and lower than the balance.\n\n")
                continue
            }
            accountBalance -= withdraw
            fmt.Printf("Your new balance is %.2f\n\n", accountBalance)
            writeBalanceToFile(accountBalance)

        default:
            // break -> break inside a 'switch' statement will only break the switch statement, not the loop
            // in this case is necessary use the 'return' ending the entire main function
            fmt.Printf("Goodbye!\n\n")
            return
        }
    }

/*         if choice == 1 {
            fmt.Printf("Your balance is %.2f\n\n", accountBalance)
        } else if choice == 2 {
            fmt.Print("Deposit amount: ")
            var depositAmount float64
            fmt.Scan(&depositAmount)
            if depositAmount <= 0 {
                fmt.Printf("Invalid amount. Must be greater than 0.\n\n")
                continue
            }
            accountBalance += depositAmount
            fmt.Printf("Your new balance is %.2f\n\n", accountBalance)
        } else if choice == 3 {
            fmt.Print("Withdrawal amount: ")
            var withdraw float64
            fmt.Scan(&withdraw)
            if withdraw <= 0 || withdraw > accountBalance {
                fmt.Printf("Invalid withdraw. Must be greater than 0 and lower than the balance.\n\n")
                continue
            }
            accountBalance -= withdraw
            fmt.Printf("Your new balance is %.2f\n\n", accountBalance)
        } else {
            break
        }
    }
    fmt.Printf("Goodbye!\n\n") */
}