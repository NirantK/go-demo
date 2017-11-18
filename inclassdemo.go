package main

import "fmt"
func main() {
	messages := make(chan string)

	go func() {messages <- "Hello "}()
	go func() {messages <- "World!"}()

	fmt.Print(<-messages)
	fmt.Println(<-messages)	
}