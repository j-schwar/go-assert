# go-assert

go-assert is a simple package which implementes convienence testing assertions using go's built in testing package.

## Installation

```
go get github.com/j-schwar/go-assert
```

## Example

The default package name is `assert`.

Functions like `assert.CondError` and `assert.FatalCondError` are useful for table tests where you want to ensure that some cases are errors.

``` go
import (
	"github.com/j-schwar/go-assert"
	"testing"
)

func TestFoo(t *testing.T) {
	tests := []struct{
		name    string
		isError bool
		// ...
	}{
		// ...
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Foo()
			// fail if unexpected error or if no error when expecting one
			assert.FatalCondError(t, test.isError, err)
		})
	}
}
```
