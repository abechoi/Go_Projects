<div align="center">
<h1>Golang Tutorial</h1>
<p>By Tech with Tim</p>
</div>

1.  [Hello World](#hello-world)
2.  [Create EXE File](#create-exe-file)
3.  [Data Types](#data-types)

## Hello World

tutorial.go
```
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```
```
# to run:
go run tutorial.go
```

## Create EXE File

```
# creates tutorial.exe with tutorial.go

# Mac
./tutorial

# Windows
go build tutorial.go
```

## Data Types

### Unsigned Integers (no negatives)

- uint8 / byte (0 to 255)
- uint16 (0 to 65535)
- uint32 (0 to 4294967295)
- uint64 (0 to 18446744073709551615)

### Signed Integers

- int8 (-128 to 127)
- int16 (-32768 to 32767)
- int32 / rune (-2147483648 to 2147483647)
- int64 (-9223372036854775808 to 9223372036854775807)

### Machine Dependent Types

- uint (32 or 64 bits)
- int (same size as uint)
- uintptr (an unsigned integer to store the uninterpreted bits of a pointer value)

### Floats

- float32 (IEEE-754 32-bit floating-point numbers)
- float64 (IEEE-754 64-bit floating-point numbers)

### Strings

- "Hello World"

```
# create a string variable
var name string
# assign "Abe" to name
name = "Abe"
```

### Booleans

- true
- false