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

    if pipe, err = events.Pipe(channelSize); err != nil {
        panic(err)
    }
    defer pipe.Close()

    var pub, sub = pipe.Publisher(), pipe.Subscriber()

    var evt = events.Event(time.Now(), "test object")
    if err = pub.Publish(evt); err != nil {
        panic(err)
    }
    
    fetchedEvent := <-sub.Get()
    if fetchedEvent == nil {
        panic("event pipe channel closed")
    }

    fmt.Printf("%+v", fetchedEvent)
}
```

## Contribution

Please contribute by creating Pull-requests or Issues here on Github.