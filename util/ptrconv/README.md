# UTIL/PTRCONV

`ptrconv` is utilities for converting an object to object ptr (*object) and vice versa.

See [Documentation](https://github.com/bearaujus/bgdk/blob/master/util/ptrconv/init.go)

# IMPORT

- Package

```go
import "github.com/bearaujus/bgdk/util/ptrconv"
```

- Mock

```go
import "github.com/bearaujus/bgdk/util/ptrconv/mock"
```

&nbsp;

# Examples

- [Convert: Object -> Ptr Object (*object)](#convert-object---ptr-object-object)
- [Convert: Ptr Object (*object) -> Object](#convert-ptr-object-object---object)
- [Convert: Nil Ptr Object (*object) -> Object](#convert-nil-ptr-object-object---object)

### Convert: Object -> Ptr Object (*object)

```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/util/ptrconv"
)

func main() {
	var str = "Hi!"
	ptrStr := ptrconv.Instance().ConvString(str)
	fmt.Println("(object)", str, "->", "(*object)", ptrStr)
}
```

```text
(object) Hi! -> (*object) 0x14000104250
```

&nbsp;

### Convert: Ptr Object (*object) -> Object

```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/util/ptrconv"
)

func main() {
	var dummyStr string = "Hi!"
	var ptrStr *string = &dummyStr
	str := ptrconv.Instance().ConvPtrString(ptrStr)
	fmt.Println("(*object)", ptrStr, "->", "(object)", str)
}
```

```text
(*object) 0x14000010270 -> (object) Hi!
```

&nbsp;

### Convert: Nil Ptr Object (*object) -> Object

```go
package main

import (
	"fmt"
	"github.com/bearaujus/bgdk/util/ptrconv"
)

func main() {
	var ptrInteger *int = nil
	integer := ptrconv.Instance().ConvPtrInt(ptrInteger)
	fmt.Println("(*object)", ptrInteger, "->", "(object)", integer)
}

```

```text
(*object) <nil> -> (object) 0
```

&nbsp;

[Back to top](#utilptrconv) 
