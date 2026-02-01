package main

import "fmt"

// Pass by value - receives a COPY
func doubleValue(n int) {
  n = n * 2
  fmt.Println("  Inside doubleValue: n =", n)
}

// Pass by pointer - receives the ADDRESS
func doublePointer(n *int) {
  *n = *n * 2
  fmt.Println("  Inside doublePointer: *n =", *n)
}

// Practical example: modify a slice header
func addElement(s *[]int, val int) {
  *s = append(*s, val)
}

func main() {
  fmt.Println("=== Pass by Value ===")
  num := 10
  fmt.Println("Before doubleValue: num =", num)
  doubleValue(num)
  fmt.Println("After doubleValue: num =", num)

  fmt.Println("\n=== Pass by Pointer ===")
  num = 10
  fmt.Println("Before doublePointer: num =", num)
  doublePointer(&num)
  fmt.Println("After doublePointer: num =", num)

  fmt.Println("\n=== Modifying Slice via Pointer ===")
  numbers := []int{1, 2, 3}
  fmt.Println("Before:", numbers)
  addElement(&numbers, 4)
  addElement(&numbers, 5)
  fmt.Println("After:", numbers)
}
