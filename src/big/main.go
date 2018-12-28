package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {
	var z1 big.Int
	z1.SetUint64(12)

	fmt.Println("Big numbers")
	fmt.Println(z1,z1.Mul(big.NewInt(5), big.NewInt(3)))
	fmt.Println(math.Pow(10, 2))
}
