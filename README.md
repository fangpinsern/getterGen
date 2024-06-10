# Go Getter Generator

Tool to generate getters for your struct in packages

```
go get -u github.com/fangpinsern/getterGen
go install github.com/fangpinsern/getterGen/tools/{insert tool name}
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

```
func (m *demo) GetArr() []*Item {
	if m != nil {
		return m.Item
	}
	return nil
}
```
