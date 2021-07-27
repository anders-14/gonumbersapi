# gonumbersapi

A go module that wraps around [numbersapi](http://numbersapi.com)

## Usage

```go
package main

import (
  "fmt"

  gna "github.com/anders-14/gonumbersapi"
)

func main() {
  // Math fact
  fmt.Println(gna.MathFact(42))

  // Fun fact
  fmt.Println(gna.FunFact(14))

  // What happened on a given date (dd/mm - This ain't murica)
  fmt.Println(gna.DateFact("31/12"))

  // What happened in a given year
  fmt.Println(gna.YearFact(2021))
}
```

