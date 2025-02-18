```go
package main

import "fmt"

func main() {
    var m map[string]int = make(map[string]int)
    // Alternatively, you can use the literal syntax to initialize the map
    // var m = map[string]int{"key": 0} 
    m["key"] = 10
    fmt.Println(m["key"])
}
```