# typ

Imagine a `*string` type: Does this type correspond to a `nullable` or an `undefinable` string?

Insecurities about types like the one above are eliminated by using the types of this package.
Distinction between `undefined`, `null` and concrete values is as important as handling the
same data types equally, over the whole application.

Build your own `Undefinable[Foo]`, `Nullable[Foo]` or even `UndefinableNullable[Foo]` type with the provided
generic types of this package. Even union types are supported by using the `AnyOf[AnyOfRuler]` generic type.

Example:

```go
import (
	"fmt"
	"github.com/inkognitro/typ"
)

id1 := typ.NewUuidV4()
id2 := typ.UndefinableUuid{}

// Try to convert an undefinable value to a defined one:
if _, ok := id2.ToUuid(); !ok {
    fmt.Println("could not convert id2 to uuid, because the underlying value is undefined")
}

// Comparison between two values:
if id1.ToUndefinableUuid().Equals(id2) {
    fmt.Println("what the heck is going on here")
}
```
