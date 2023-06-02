// recover from a panic. can stop panic and let it continue with the execution instead
package main

import "fmt"

func mayPanic() {
    panic("a big problem")
}

func madhsauijdhasuiin() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:package main: \n", r)
        }
    }() 
    mayPanic()
    
    fmt.Println("After mayPanic()")
}
