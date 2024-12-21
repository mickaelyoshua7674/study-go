package main
import (
    "fmt"
    "example.com/bank/fileops"
    // To import created packages, must import the full path
    // starting with the module path.
    // New package must be inside its own folder and inside the main folder.
    "github.com/Pallinder/go-randomdata"
    // With 'go get [package path]' it is possible to install third party packages.
    // Those packages are installed globally and the reference to the
    // installed package is on the 'go.mod' file, so others can see
    // and download the same packages in there simply running 'go get'.
)

const balanceFile = "balance.txt"

func main() {
    accountBalance, err := fileops.GetFloatFromFile(balanceFile)
    if err != nil {
        fmt.Println("ERROR")
        fmt.Println(err)
        fmt.Println("----------")
        panic("can't continue")
        // panic() function will stop the program execution in a way that looks like it crashed
    }

    fmt.Printf("\n*** Welcome to Go Bank! ***\nPhone number %v\n\n", randomdata.PhoneNumber())
    for {
        presentOptions()

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
            fileops.WriteFloatToFile(accountBalance, balanceFile)

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
            fileops.WriteFloatToFile(accountBalance, balanceFile)

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