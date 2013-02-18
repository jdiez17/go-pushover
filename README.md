go-pushover
===========

Thin Go wrapper for the Pushover API.

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
		ts := time.Now.Add(time.Minute * 5)
		p := pushover.Pushover{
			"KzGDORePK8gMaC0QOYAMyEEuzJnyUi" /* API key */,
			"pQiRzpo4DXghDmr9QzzfQu27cmVRsG" /* User */,
		}
		n := pushover.Notification{
			Title:     "go-pushover",
			Message:   "Hello from Go!",
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
