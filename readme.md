# Golang Weighted Random Choice
This is a simple package for creating Weighted Random Choices.

### Usage
To use WRC, import :
```go
import "github.com/kontoulis/go-weighted-random-choice"
``` 
Instantiate a `WeightedRandomChoice` object by calling `New()`.
`New()` takes optional argument `precision` which is by default `1000` which
is used to calibrate the weights for more precise results.
```go
wrc := WeightedRandomChoice.New()
```
Add elements one by one with `name` and `weight`.
```go
wrc.AddElement("common", 59)
wrc.AddElement("epic", 1)
wrc.AddElement("rare", 10)
wrc.AddElement("green", 30)
```
You can also add multiple elements by passing a `map[string]int` in `AddElements` 
```go
wrc.AddElements(map[string]int{
    "common"    : 59,
    "epic"      : 1,
    "rare"      : 10,
    "green"     : 30,
})
```
Then you can call `GetRandomChoice` to retrieve a random weighted choice
```go
choice := wrc.GetRandomChoice()
```
