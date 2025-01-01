package main
import (
    "fmt"
    "example.com/practice/product"
)

func main() {
    hobbies := [3]string{"dance", "gym", "read"}
    fmt.Println(hobbies)

    fmt.Printf("\nFirst element: %v\nSecond and third element: %v\n\n", hobbies[0], hobbies[1:3])
    //bestHobbies := hobbies[0:2]
    bestHobbies := hobbies[:2]
    fmt.Println(bestHobbies)

    bestHobbies = bestHobbies[1:3]
    fmt.Println(bestHobbies)

    courseGoals := []string{"learn Go", "Make an API"}
    fmt.Println(courseGoals)
    courseGoals[1] = "Build an API"
    courseGoals = append(courseGoals, "Make Go my primary language")
    fmt.Println(courseGoals)

    products := []product.Product{product.New(1, "A Book", 9.99), product.New(2, "A Hat", 5.99)}
    fmt.Println(products)
    products = append(products, product.New(3, "A Shoe", 20.99))
    fmt.Println(products)
}
