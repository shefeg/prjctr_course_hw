// exercise
// 1. Make a roaring tiger, which extends a cat
// implement: mewing and roaring
// 2. Make a trash bin, which have an ability to accept garbage, can be filled or empty

package main

import (
	"fmt"
)

type Bin struct {
	Garbage int
}

func (b *Bin) Fill(garbage int) int {
	b.Garbage += garbage
	return b.Garbage
}

func (b *Bin) Empty() int {
	b.Garbage = 0
	return b.Garbage
}

func main() {
	trash_bin := &Bin{Garbage: 0}
	fmt.Printf("Bin size: %+v\n", trash_bin.Fill(10))
	fmt.Printf("Bin size: %+v\n", trash_bin.Fill(5))
	fmt.Printf("Bin size: %+v\n", trash_bin.Empty())
	fmt.Printf("Bin size: %+v\n", trash_bin.Fill(13))
	fmt.Printf("Bin size: %+v\n", trash_bin.Empty())
}
