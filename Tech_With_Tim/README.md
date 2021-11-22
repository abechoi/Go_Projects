<div align="center">
<h1>Golang Tutorial</h1>
<p>By Tech with Tim</p>
</div>

1.  [Hello World](#hello-world)
2.  [Create EXE File](#create-exe-file)
3.  [Data Types](#data-types)
4.  [Printing](#printing)

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

```
# check data types by printing %T
fmt.Printf("%T", number)
```

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

## Printing

### Values and Types

```
# print data types by printing %T
fmt.Printf("%T", number)

# print values by printing %v
fmt.Printf("%v", number)

# print literal "%"
fmt.Printf("%%")

# print boolean by printing %t
fmt.Printf("%t", bl)
```

### Integers and Floats

```
# print integer by base 2, 8, 10, and 16
fmt.Printf("%b", binary)
fmt.Printf("%o", octal)
fmt.Printf("%d", decimal)
fmt.Printf("%x", hexadecimal)

# print floats by scientific notation, decimal no exponent, and large exponents
fmt.Printf("%e", scientificNotation)
fmt.Printf("%f or %F", decimalNoExponent)
fmt.Printf("%g", largeExponents)
```

### Strings

```
# print string by printing %s
fmt.Printf("%s", abe)

# print double quoted string by printing %q
fmt.Printf("%q", "abe")
```

### Width and Precision

```
# print default width and precision by printing %f
fmt.Printf("%f", str)

# print width 9 and default precision by printing %f
fmt.Printf("%9f", str)

# print default width and precision 2 by printing %f
# 1.2345 => 1.23
fmt.Printf("%.2f", str)

# print width 9 and precision 2 by printing %f
fmt.Printf("%9.2f", str)

# print width 9 and precision 0 by printing %f
fmt.Printf("%9.f", str)
```

### Padding

```
# print with padding left
# '    abe'
fmt.Printf("%4s", "abe")

# print with padding right
# 'abe    '
fmt.Printf("%-4s", "abe")

# print with zeroes as padding
# '0000abe'
fmt.Printf("%04s", "abe")

# print with padding of arbitrary length
# '    abe'
fmt.Printf("%*s", 4,"abe")
```

### Sprint

```
# store sprint into a variable to print
var out string = fmt.Sprintf("Save this message, and print later.")
fmt.Println(out)
```