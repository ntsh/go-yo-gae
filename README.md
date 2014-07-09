go-yo
=====

Golang client for the Yo API

## Usage

With token in hand, to create a new Yo client simply

```go
package main

import (
  "github.com/sjkaliski/go-yo"
)

func main() {
  client := yo.NewClient("my_token")
}
```

To send a message to all users who subscribe to you

```go
err := client.YoAll()
```

To send a message to a specific user

```
err := client.YoUser("some_user")
```

## Tests

To run tests

`$ go test ./...`