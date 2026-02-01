package main

import "fmt"

func main() {
  // A pointer holds the memory address of a value
  x := 42

  // & is the address-of operator
  p := &x

  fmt.Println("Value of x:", x)
  fmt.Println("Address of x:", p)

  // * is the dereference operator
  // It gives us the value at that address
  fmt.Println("Value at address p:", *p)

  // The type of p is *int (pointer to int)
  fmt.Printf("Type of p: %T\n", p)

  // Another example
  name := "Alice"
  namePtr := &name

  fmt.Println("\nName:", name)
  fmt.Println("Address:", namePtr)
  fmt.Println("Dereferenced:", *namePtr)
  fmt.Printf("Type: %T\n", namePtr)
}
