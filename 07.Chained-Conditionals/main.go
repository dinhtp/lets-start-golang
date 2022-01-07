package main

import "fmt"

/*
@Resource
- Youtube:
    + https://www.youtube.com/watch?v=QvPa8C0y9yc
- Go Documentation:
    + https://go.dev/ref/spec#Operators_and_punctuation
    + https://go.dev/ref/spec#Boolean_types
    + https://go.dev/ref/spec#Logical_operators
*/

func main() {
    a := true
    b := false
    c := true
    d := false

    fmt.Println(fmt.Sprintf("a: %t", a))
    fmt.Println(fmt.Sprintf("b: %t", b))
    fmt.Println(fmt.Sprintf("c: %t", c))
    fmt.Println(fmt.Sprintf("d: %t", d))

    // TODO: evaluate a, b, c, d to make the final value of true

    // TODO: evaluate a, b, c, d to make the final value of false

    // TODO: evaluate a, c to make the final value of false

    // TODO: evaluate b, d to make the final value of true
}
