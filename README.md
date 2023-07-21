# easyweb

like streamlit, can quickly generate a web

## how work

1. Server and client interact through websocket
2. The elements of the page are dynamically generated by jquery
3. The user opens the page, and the page connects to the server through websocket
4. The server executes the code, dynamically creates elements, and sends them to the client
5. The elements(First-level) have a unique id, through which to create, update, delete
6. websocket + jquery + bootstrap

## exampel

```golang
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/lengzhao/easyweb"
    "github.com/lengzhao/easyweb/e"
)

func main() {
    easyweb.NewPage(easyweb.IndexPage, func(page easyweb.Page) {
        page.Title("MyWeb")
        page.Write("this is my first ui.")
    })
    http.ListenAndServe(":8080", nil)
}
```

For more examples, please see the example folder
