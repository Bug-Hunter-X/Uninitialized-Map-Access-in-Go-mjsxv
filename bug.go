```go
package main

import "fmt"

func main() {
    var m map[string]int
    m["key"] = 10 // This will cause a runtime panic if m is not initialized
    fmt.Println(m["key"])
}
```