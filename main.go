package main

import "fmt"

// Greatest Common Divisor
func GCD(a int, b int) int {
    // stop condition
    if (a == b) { return a }
    if (a > b) { return GCD(a-b, b) }
    if (a < b) { return GCD(a, b-a)
    } else {
        return 0
    }
}

// Fibonacci
func FBN(n int, pn1 int, pn2 int) int {
    if (n == 1) { return pn2; }
    fmt.Print(pn2)
    fmt.Print(" ")
    return FBN(n-1, pn2, (pn1+pn2))
}

func main() {
    // GCD
    var a int = 10
    var b int = 4
    var result = GCD(a, b)
    fmt.Print("result of GCD of ")
    fmt.Print(a)
    fmt.Print(" and ")
    fmt.Print(b)
    fmt.Print(": ")
    fmt.Println(result)

    // FBN
    var n = 5
    fmt.Print("first ")
    fmt.Print(n)
    fmt.Print(" of Fibonacci numbers:  ")
    result = FBN(5, 1, 1)
    fmt.Println(result)
}

