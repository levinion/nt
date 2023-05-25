# nt

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/levinion/nt)

![GitHub](https://img.shields.io/github/license/levinion/nt)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/levinion/nt)


`nt` is a Go package that provides a simple and flexible way to create and use templates with context variables and functions.

## Installation
To install nt, use the `go get` command:

```sh
go get github.com/levinion/nt
```

## Usage
To use nt, you need to import it in your Go files:

```sh
import "github.com/levinion/nt"
```

Then you can create a template with the `Create` function, which takes an optional name argument. If you specify a name, you can later find the template with the `Find` function.

```go
t := nt.Create("mytemplate")
```

You can register or append template functions with the Join method, which takes a function that accepts a `*nt.Ctx` argument. The `Ctx` type is a struct that holds a map of variables, an error field, and a read-only flag. You can use the `Set`, `SafeSet`, and `Get` methods to manipulate the variables in the context. You can also use the `Error` method to set an error and stop the remaining functions from executing.

```go
t.Join(func(c *nt.Ctx) {
    c.Set("name", "Alice")
    c.Set("age", 25)
})
```

You can call the template with the `Call` method, which also takes a function that accepts a `*nt.Ctx` argument. This function can access the variables and error set by the previous functions, but cannot modify them.

```go
err := t.Call(func(c *nt.Ctx) {
    name := c.Get("name")
    age := c.Get("age")
    fmt.Printf("Hello, %s. You are %d years old.\n", name, age)
})
if err != nil {
    // handle error
}
```

You can also add variables to the context with the `Watch` and `WatchMany` methods, which take a key-value pair and a map of key-value pairs respectively.

```go
t.Watch("city", "New York")
t.WatchMany(map[string]any{
    "country": "USA",
    "state": "NY",
})
```

You can merge two templates with the `Concat` method, which combines their contexts and functions.

```go
t1 := nt.Create()
t1.Join(func(c *nt.Ctx) {
    c.Set("foo", "bar")
})

t2 := nt.Create()
t2.Join(func(c *nt.Ctx) {
    c.Set("baz", "qux")
})

t1.Concat(t2) // t1 now has both foo and baz variables
```

## License
`nt` is licensed under the `MIT` license. See LICENSE for more details.