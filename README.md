# Sorted Set

Thread-safe sorted set implementation.

It maintains both a map for O(1) lookups and a sorted slice for ordered iteration.

## Usage

```go
import "github.com/zkqiang/sortedset"

func main() {
    set := sortedset.New(func(i, j int) bool { return i < j })
    
    set.Add(3)
    set.Add(1)
    set.Add(4)
    set.Add(1) // Duplicate
    
    println(set.Elements()) // [1 3 4]
}
```
