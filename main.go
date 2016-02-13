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

// Factorial
func FCT(x int, n int) int {
    if (n == 1) { return x }
    if (x == 0) { return 1 }
    fmt.Print(x)
    fmt.Print(" ")
    return FCT(x*(n-1), n-1)
}

func Factorial(n int) int {
    if (n == 0) { return 1 }
    return n * Factorial(n-1)
}

// Multiply
func Multiply(a int, b int) int {
    if (a == 1) { return b }
    return a * Multiply(1, b)
}

func main() {
    // GCD
    fmt.Print("Greatest common divisor for 10 and 4: ")
    fmt.Println(GCD(10, 4))


    // FBN
    fmt.Print("first 5 Fibonacci numbers: ")
    fmt.Println(FBN(5, 1, 1))

    // PWR
    fmt.Print("3 to the power of 1: ")
    fmt.Println(PWR(3, 3, 1))

    fmt.Print("4 to the power of 3: ")
    fmt.Println(PWR(4, 4, 3))

    // FCT 
    fmt.Print("Factorial of 6 is: " )
    fmt.Println(FCT(6,6))

    // Factorial
    fmt.Print("Factiorial of 6 is still: ")
    fmt.Println(Factorial(6))

    // Multiply
    fmt.Print("Multiply 4 and 5: ")
    fmt.Println(Multiply(4, 5))
}

