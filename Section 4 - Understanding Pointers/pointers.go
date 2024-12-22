package main
import "fmt"

func main() {
    age := 32
/*     agePointer := &age
    fmt.Println("Age Address:", agePointer)
    fmt.Println("Age (dereference):", *agePointer) */
/*     fmt.Println("Age:", age)
    fmt.Println("Adult years:", getAdultYears(&age)) */
    fmt.Println("Age:", age)
    getAdultYears(&age)
    fmt.Println("Adult years:", age)
}

/* func getAdultYears(age *int) int {
    return *age-18
} */
func getAdultYears(age *int) {
    *age-=18
}