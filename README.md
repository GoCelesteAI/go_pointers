# Go Lesson 14: Pointers

Demo code from Go Tutorial for Beginners - Lesson 14.

## Files

- `pointer_basics.go` - Address-of (&) and dereference (*) operators
- `pointer_modify.go` - Modifying values through pointers
- `pointer_functions.go` - Pass by value vs pass by pointer
- `pointer_nil.go` - Nil pointers, new(), and safety

## Running the Examples

```bash
# Pointer basics
go run pointer_basics.go

# Modifying via pointers
go run pointer_modify.go

# Pointers in functions
go run pointer_functions.go

# Nil pointers and safety
go run pointer_nil.go
```

## Key Concepts

### Address-of and Dereference
```go
x := 42
p := &x           // & = address-of
fmt.Println(*p)   // * = dereference (prints 42)
```

### Modify via Pointer
```go
*p = 100          // Changes x to 100
```

### Pass by Pointer
```go
func doublePointer(n *int) {
    *n = *n * 2   // Modifies original
}

num := 10
doublePointer(&num)  // num is now 20
```

### Nil Safety
```go
var p *int        // p is nil
if p != nil {
    fmt.Println(*p)  // Safe to dereference
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #14 - Pointers](https://youtube.com/@CelesteAI)

## License

MIT
