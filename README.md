# Assert

[![GoDoc](https://godoc.org/github.com/tidwall/assert?status.svg)](https://godoc.org/github.com/tidwall/assert)


This package provides an assert function for Go. 
It's designed to work like [assert](https://man7.org/linux/man-pages/man3/assert.3.html) in C.

## Example

```go
package my_test

import "github.com/tidwall/assert"

func TestMyThing(t *testing.T) {
    assert.Assert("hello" == "jello")
}
```

This will print the following message and abort the program.

```
Assertion failed: ("hello" == "jello"), function TestMyThing, file my_test.go, line 6.
```
