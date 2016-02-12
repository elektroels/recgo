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

// Power to
func PWR(x int, x1 int, n int) int {
    if (n == 1) { 
        return x }
    fmt.Print(x)
    fmt.Print(" ")
    return PWR(x1*x, x1, n-1)
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
    fmt.Print(" of Fibonacci numbers: ")
    result = FBN(n, 1, 1)
    fmt.Println(result)

    // PWR
    n = 1
    var x int = 3
    fmt.Print(x)
    fmt.Print(" to the power of ")
    fmt.Print(n)
    fmt.Print(": ")
    result = PWR(x, x, n)
    fmt.Println(result)
    //var result2 int = PWR(3, 3, 4)
    n = 4
    fmt.Print(x)
    fmt.Print(" to the power of ")
    fmt.Print(n)
    fmt.Print(": ")
    result = PWR(x, x, n)
    fmt.Println(result)
}

