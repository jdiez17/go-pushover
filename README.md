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

    n := pushover.Notification {
        Title: "go-pushover",
        Message: "Hello from Go!",
        Timestamp: time.Now(),
        Priority: 2,
        Retry: 30,
        Expire: 90,
    }

    response, err := p.Notify(n)

    if err != nil {
        if err != pushover.PushoverError {
            panic(err)
        } else {
            fmt.Println(err)
            fmt.Println(response)
        }
    }
}
```
