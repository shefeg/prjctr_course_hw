// exercise
// make two couples, which have and don't have a child (Couple { ... *Child } )

package main

import (
	"couples/types"
	"fmt"
)

func main() {
	couple_without_child := types.Couple{}
	couple_with_child := types.Couple{
		Child: &types.Child{
			Name: "Erik",
			Age:  2,
		},
	}
	fmt.Printf("Happy couple: %+v\n", couple_with_child.Child)
	fmt.Printf("Sad couple: %+v\n", couple_without_child.Child)
}
