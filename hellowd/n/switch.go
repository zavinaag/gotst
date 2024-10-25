package n

import (
	"fmt"
)

func Sw() {
	var i int
	fmt.Print("Type a number: ")
	fmt.Scan(&i)

	switch {
	case i == 1:
		fmt.Println("You enter 1")
	case i == 2 || i == 3:
		fmt.Println("You enter 2 or 3")
	case i > 3 && i <= 10:
		fmt.Println("You enter ew")
	default:
		fmt.Println("You enter more than 4")
	}
}
