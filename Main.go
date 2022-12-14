package main

import "fmt"

var p1 = fmt.Println

func main() {
    iAge := 88
    if (iAge >= 1) && (iAge <= 18) {
            p1("Important Birthday")
    }   else if (iAge == 21) || (iAge == 50) {
            p1("wel Birthday")
    }   else if iAge >= 65 {
            p1("Geen Birthday")
    }   else {
            p1("Niet Birthday")
    }
    p1("!true =", !true)
}