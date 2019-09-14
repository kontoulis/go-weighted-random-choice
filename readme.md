# Golang Weighted Random Choice
This is a simple package for creating Weighted Random Choices.

### Usage
Instantiate a `WeightedRandomChoice` object by calling `New()`.
`New()` takes optional argument `precision` which is by default `1000`.
which is used to calibrate the weights for more precise results.
Add elements with `name` and `weight` and then call `GetRandomChoice` to retrieve a random weighted choice.
```go
package main

import (
	"WeightedRandomChoice"
)
func main(){
    wrs := WeightedRandomChoice.New()
    // Add elements one by one 
    wrs.AddElement("common", 59)
    wrs.AddElement("epic", 1)
    wrs.AddElement("rare", 10)
    wrs.AddElement("green", 30)
    // OR add elements as a map[string]int
    wrs.AddElements(map[string]int{
        "common"    : 59,
        "epic"      : 1,
        "rare"      : 10,
        "green"     : 30,
    })

    choice := wrs.GetRandomChoice()
}
```