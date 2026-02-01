package main

import "fmt"

func main() {
  // Pointers let us modify the original value
  x := 10
  p := &x

  fmt.Println("Before modification:")
  fmt.Println("  x =", x)
  fmt.Println("  *p =", *p)

  // Modify via the pointer
  *p = 100

  fmt.Println("\nAfter *p = 100:")
  fmt.Println("  x =", x)
  fmt.Println("  *p =", *p)

  // Both x and *p refer to the same memory location
  x = 200

  fmt.Println("\nAfter x = 200:")
  fmt.Println("  x =", x)
  fmt.Println("  *p =", *p)

  // Practical example: swapping values
  a := 5
  b := 10

  fmt.Println("\nBefore swap: a =", a, ", b =", b)

  temp := a
  a = b
  b = temp

  fmt.Println("After swap: a =", a, ", b =", b)
}
