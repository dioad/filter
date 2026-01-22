# Filter Package

A Go package providing generic filtering utilities for maps and slices with various logical operations.

## Features

- Filter maps and slices with simple predicates
- Combine multiple filters with logical operations (AND, OR, NOT AND, NOT OR)
- Select elements from slices (first, last, random)
- Convert slices to filter functions
- Type-safe operations using Go generics

## Installation

```bash
go get github.com/dioad/filter
```

## Usage

### Filtering Slices

```go
package main

import (
    "fmt"
    "github.com/dioad/filter"
)

func main() {
    // Basic filtering
    numbers := []int{1, 2, 3, 4, 5}
    
    // Filter even numbers
    isEven := func(n int) bool {
        return n%2 == 0
    }
    
    evenNumbers := filter.FilterSlice(numbers, isEven)
    fmt.Println(evenNumbers) // Output: [2 4]
    
    // Combining filters with AND
    isGreaterThanThree := func(n int) bool {
        return n > 3
    }
    
    evenAndGreaterThanThree := filter.FilterSliceAnd(numbers, isEven, isGreaterThanThree)
    fmt.Println(evenAndGreaterThanThree) // Output: [4]
    
    // Combining filters with OR
    isLessThanTwo := func(n int) bool {
        return n < 2
    }
    
    evenOrLessThanTwo := filter.FilterSliceOr(numbers, isEven, isLessThanTwo)
    fmt.Println(evenOrLessThanTwo) // Output: [1, 2, 4]
}
```

### Filtering Maps

```go
package main

import (
    "fmt"
    "github.com/dioad/filter"
)

func main() {
    // Map filtering
    users := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Carol": 22,
        "Dave":  35,
    }
    
    // Filter users older than 25
    isOlderThan25 := func(age int) bool {
        return age > 25
    }
    
    olderUsers := filter.FilterMap(users, isOlderThan25)
    fmt.Println(olderUsers) // Output: map[Bob:30 Dave:35]
}
```

### Selecting Elements from Slices

```go
package main

import (
    "fmt"
    "github.com/dioad/filter"
)

func main() {
    fruits := []string{"apple", "banana", "cherry", "date"}
    
    // Select first element
    first := filter.SliceSelectFirst(fruits)
    fmt.Println(*first) // Output: apple
    
    // Select last element
    last := filter.SliceSelectLast(fruits)
    fmt.Println(*last) // Output: date
    
    // Select random element
    random := filter.SliceSelectRandom(fruits)
    fmt.Println(*random) // Output: (random fruit)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.