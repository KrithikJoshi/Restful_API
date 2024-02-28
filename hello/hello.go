package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    //setting the properties to predefined logger , including the log entry prefix and a flag to disable printing the time, source file name and line number
    log.SetPrefix("greetings:")
    log.SetFlags(0)
    //requesting a greeting message
    message, err := greetings.Hello("")
    // if an error was returned, print it to the console and exit the program
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(message)
}