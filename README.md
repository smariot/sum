# sum

This package can be used to determine the sum of values in a slice.

This is hundreds of times slower than repeated addition, so don't use it unless accuracy is important and
your values have wildly different magnitudes.
```golang
import (
    "fmt"
    "github.com/smariot/sum"
)

fmt.Println("sum:", sum.Slice([]float64{1, 2, 3})
```
