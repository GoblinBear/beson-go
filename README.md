# BESON - Binary Extended JSON

[![GoDoc](https://godoc.org/github.com/GoblinBear/beson-go?status.svg)](https://godoc.org/github.com/GoblinBear/beson-go)

Beson library is similar to BSON format used in mongodb. The major difference between beson and bson is that beson allows primitive data to be encoded directly. Beson is designed to transfer or store data in a binary format, not specialized for database storage.

## Table of Contents

- [Features](#Features)
- [Installation](#Installation)
- [Quick start](#Quick-start)
    - [Initialize a big integer number](#Initialize-a-big-integer-number)
    - [Serialize](#Serialize)
    - [Deserialize](#Deserialize)
- [Usage](#Usage)
- [License](#License)

## Features

- Big integer number operations.
    - 128-bit integer：`UInt128` / `Int128`
    - 256-bit integer：`UInt256` / `Int256`
    - 512-bit integer：`UInt512` / `Int512`
    - Variable length integer：`UIntVar` / `IntVar`
- Serialize data to binary sequence.
- Desrialize binary sequence to original data.

## Installation

Download and install it:

```shell
$ go get -u github.com/GoblinBear/beson-go
```

Import it in your code (serialize / deserialize):

```go
import "github.com/GoblinBear/beson-go"
```
Import it in your code (big integer number):

```go
import "github.com/GoblinBear/beson-go/types"
```

## Quick start

### Initialize a big integer number

```go
package main

import (
    "fmt"
    "github.com/GoblinBear/beson-go/types"
)

func main() {
    v1 := types.NewUInt128("1844674407370955161825", 10)
    fmt.Println(v1)
}
```
```shell
&{[206 8 0 0 0 0 0 0 232 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
```

### Serialize

```go
package main

import (
    "fmt"
    "beson"
    "beson/types"
)

func main() {
    v := types.NewUInt32(2568)
    fmt.Println(v)

    ser := beson.Serialize(v)
    fmt.Println(ser)
}
```
```shell
&{2568}
[3 0 8 10 0 0]
```

### Deserialize

```go
package main

import (
    "fmt"
    "beson"
    "beson/types"
)

func main() {
    v := types.NewUInt32(2568)
    ser := beson.Serialize(v)

    data, anchor := beson.Deserialize(ser, 0)
    fmt.Println(data)
    fmt.Println(anchor)
}
```
```shell
6
&{2568}
```

## Usage

- See the wiki page for details：[wiki](https://github.com/GoblinBear/beson-go/wiki)

## License

beson-go source code is licensed under the [Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
