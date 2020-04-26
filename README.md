![Go](https://github.com/artheus/go-events/workflows/Go/badge.svg?branch=master)

# Go Events

Multi-purpose generic pub/sub event system for go applications.
Built on golang channels.

## Usage

Simple example code:
```go
package main

import (
    "fmt"
    "github.com/artheus/go-events"
    "github.com/artheus/go-events/types"
    "time"
)

const channelSize = 10

func main() {
    var err error
    var pipe types.Pipe

    // Create event pipe for managing publishers and subscribers
    if pipe, err = events.Pipe(channelSize); err != nil {
        panic(err)
    }
    defer pipe.Close()

    // Create a publisher and a subscriber from the pipe
    var pub, sub = pipe.Publisher(), pipe.Subscriber()
    defer pub.Close()
    defer sub.Close()

    // Publish a string event to the pipe
    var evt = events.Event(time.Now(), "test object")
    if err = pub.Publish(evt); err != nil {
        panic(err)
    }
    
    // Fetch the published string event from subscriber
    fetchedEvent := <-sub.Get()
    if fetchedEvent == nil {
        panic("event pipe channel closed")
    }

    fmt.Printf("%+v", fetchedEvent)
}
```

## Contribution

Please contribute by creating Pull-requests or Issues here on Github.
