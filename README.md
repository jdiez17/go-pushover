go-pushover
===========

Go wrapper for the Pushover API.

Example
=======

```go
package main

import (
    "github.com/jdiez17/go-pushover"
    "fmt"
    "time"
)

func main() {
    p := pushover.Pushover{
        "", /* User key (use your own)*/
        "", /* Application key (use your own) */
    }

    ts := time.Now()
    n := pushover.Notification {
        Title: "go-pushover",
        Message: "Hello from Go!",
        Timestamp: ts,
    }

    response, err := p.Notify(n)
    if err != nil {
        if err != pushover.PushoverError {
            panic(err)
        }
        fmt.Println(response.Message)
    }
}
```
