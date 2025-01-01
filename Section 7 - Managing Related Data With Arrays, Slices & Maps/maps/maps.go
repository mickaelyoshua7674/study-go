package main

import (
   "fmt"
)

func main() {
    // Inside [] is the type of the key and after is the type of the value.
    // Maps are always dynamic, it's possible to add new keys and values.
    websites := map[string]string{
        "Google":              "google.com",
        "Amazon Web Services": "aws.com",
    }
    fmt.Println(websites)
    fmt.Println(websites["Google"])
    websites["LinkedIn"] = "linkedin.com"
    fmt.Println(websites["LinkedIn"])

    delete(websites, "Google")
    fmt.Println(websites)
}