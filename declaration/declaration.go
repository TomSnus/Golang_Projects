package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	i, j int = 1, 2
	c, c1  bool = true, false
)

func main() {

	k := 3
	fmt.Println(i, j, k, c, c1)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//Variable Type with %T
	fmt.Printf("i is of type %T\n", i)

	
}
 /*
 The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
  */