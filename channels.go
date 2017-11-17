// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"
func main() {

    // Create a new channel with `make(chan val-type buffer-size)`.
    // Channels are typed by the values they convey.
    // Here we `make` a channel of strings buffering up to
    // 2 values.
    messages := make(chan string, 2)

    // _Send_ a value into a channel using the `channel <-`
    // syntax. Here we send `"ping"`  to the `messages`
    // channel we made above, from a new goroutine.

    // Because this channel is buffered, we can send these
    // values into the channel without a corresponding
    // concurrent receive.
    go func() { messages <- "first ping" }()
    go func() { messages <- "second ping" }()

    // The `<-channel` syntax _receives_ a value from the
    // channel. Here we'll receive the messages
    // we sent above and print it out.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
