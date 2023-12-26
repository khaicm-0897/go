// package-init/rules-illustration/currency/currency.go
package currency

import "fmt"

var f = func() string {
    fmt.Println("variable f initialized")
    return "test"
}()

var d = func() string {
    fmt.Println("variable d initialized", a)
    return "value of c"
}()

var c = func() string {
    fmt.Println("variable c initialized", b)
    return "value of c"
}()

var a = func() string {
    fmt.Println("variable a initialized")
    return "value of a"
}()

var b = func() string {
    fmt.Println("variable b initialized", a)
    return "value of b"
}()

func init() {
    fmt.Println("currency init")
}

func EuroSymbol() string {
    return "EUR"
}
