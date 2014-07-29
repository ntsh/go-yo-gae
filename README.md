go-yo-gae
=====

Golang client for the Yo API compatible with Google App Engine

[Read this](https://developers.google.com/appengine/docs/go/gettingstarted/introduction) to get started with Go on GAE


## Usage

With [token](http://yoapi.justyo.co/) in hand, to create a new Yo client simply


    package main

    import (
      "github.com/ntsh/go-yo-gae"
    )

    func main() {
      client := yo.NewClient("my_token")
    }


To send a message to all users who subscribe to you

```
err := client.YoAll(r) // r *http.Request
```

To send a message to a specific user

```
err := client.YoUser("some_user", r)
```

##Example
    package main

    import (
        "fmt"
        "net/http"
         "github.com/ntsh/go-yo-gae"
    )

    func init() {
        http.HandleFunc("/", handler)
    }

    func handler(w http.ResponseWriter, r *http.Request) {
        user := r.URL.Query()["user"][0]
        fmt.Fprint(w,user)
        client := yo.NewClient(TOKEN)
        err := client.YoUser(user, r)
        if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }

To launch example, simply install go-sdk from GAE and follow instructions here:
[https://developers.google.com/appengine/docs/go/gettingstarted/introduction](https://developers.google.com/appengine/docs/go/gettingstarted/introduction)