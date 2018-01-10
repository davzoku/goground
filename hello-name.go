package main

import (
	"fmt"
    "os"
    "time"
)
    
func main() {
	var name string
    
    fmt.Println("hello, world")
    fmt.Println("what is your name?")
    fmt.Scanf("%s", &name)
    fmt.Fprintf(os.Stdout, "hi, %s! ", name)
    fmt.Println("The time is", time.Now())
    
    
}