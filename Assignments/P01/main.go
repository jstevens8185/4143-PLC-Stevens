//	FileName:	main.go
//	Author:		Jered Stevens
//	Date:		09/06/2023
//	Class:		4143 - PLC - Griffin

// Driver code for program 1
//
//	 Prints the mascot name for MSU
//		by calling a function from
//		a file in another folder
package main

import (
	"example/user/P01/mascot"
	"fmt"
)

func main() {
	fmt.Println(mascot.BestMascot())

}
