package main

import "fmt"

func GCD(a int, b int) int {
    // stop condition
    if (a == b) { return a }
    if (a > b) { return GCD(a-b, b) }
    if (a < b) { return GCD(a, b-a)
    } else {
        return 0
    }
}

func main() {
    var a int = 10
    var b int = 4
    var result = GCD(a, b)
    fmt.Print("result of GCD of ")
    fmt.Print(a)
    fmt.Print(" and ")
    fmt.Print(b)
    fmt.Print(": ")
    fmt.Println(result)
}

