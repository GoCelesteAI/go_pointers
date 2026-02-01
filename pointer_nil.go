package main

import "fmt"

func main() {
  // A pointer with no value is nil
  var p *int

  fmt.Println("=== Nil Pointers ===")
  fmt.Println("Value of p:", p)
  fmt.Println("Is p nil?", p == nil)

  // DANGER: Dereferencing nil causes a panic!
  // Uncomment to see: fmt.Println(*p) // panic!

  // Always check for nil before dereferencing
  if p != nil {
    fmt.Println("Value:", *p)
  } else {
    fmt.Println("p is nil, cannot dereference")
  }

  // Assign a value to make it non-nil
  x := 42
  p = &x

  fmt.Println("\nAfter p = &x:")
  fmt.Println("Value of p:", p)
  fmt.Println("Is p nil?", p == nil)
  fmt.Println("Dereferenced:", *p)

  fmt.Println("\n=== Using new() ===")
  q := new(int)
  fmt.Println("Value of q:", q)
  fmt.Println("*q (zero value):", *q)

  *q = 100
  fmt.Println("After *q = 100:", *q)

  fmt.Println("\n=== No Pointer Arithmetic ===")
  fmt.Println("Go does not allow pointer arithmetic")
  fmt.Println("Use slices for sequential access instead")
}
