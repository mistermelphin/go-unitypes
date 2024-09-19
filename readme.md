# go-unitypes

A Go package providing numeric data types that seamlessly marshal and unmarshal JSON numbers represented as strings or numbers.

## Overview

The `unitypes` package offers custom numeric types for Go that handle JSON fields which can be either a string or a number. This is especially useful when dealing with inconsistent JSON inputs.

Supported types:

- Int
- Int8
- Int16
- Int32
- Int64
- Float32
- Float64

## Installation

```bash
go get github.com/mistermelphin/go-unitypes
```

## Usage

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/mistermelphin/go-unitypes"
)

type MyStruct struct {
    Num       unitypes.Int32 `json:"num"`
    NumString unitypes.Int32 `json:"num_string"`
}

func main() {
    data := []byte(`{"num": 123, "num_string": "123"}`)

    var s MyStruct
    if err := json.Unmarshal(data, &s); err != nil {
        panic(err)
    }

    fmt.Printf("Num: %d, NumString: %d\n", s.Num, s.NumString) // Num: 123, NumString: 123
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
