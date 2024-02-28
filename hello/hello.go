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
    names := []string{"Gladys", "Samantha", "Darrin"}
    //requesting a greeting message
    messages, err := greetings.Hellos(names)
    // if an error was returned, print it to the console and exit the program
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(messages)
}