# Go Getter Generator

Tool to generate getters for your struct in packages

```
go install github.com/fangpinsern/getterGen
```

Inspired/Modified from go tool stringers package.

## Details

Currently only works with pointer types.

```go
type Foo struct {
    Bar *string
}
```

## Future updates

Support non pointer types in struct - should not be too hard. Just a different format

Testing final
