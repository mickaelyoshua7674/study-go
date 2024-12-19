package main
import "fmt"

func main() {
    accountBalance := 1000.

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