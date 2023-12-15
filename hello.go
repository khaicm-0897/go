package main
import "fmt"

func main() {
fmt.Printf("Welcome to FOSSLinux Users\n")
n := 2548
fmt.Printf("%x", n)
b := make([]byte, 0)
b = append(b, 255)
b = append(b, 10)
fmt.Println(b)
}
