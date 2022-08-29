package useio

import (
	"fmt"
	"os"
)

func UserIO() {

	ar := os.Args
	var a int
	fmt.Scanln(&a)

	fmt.Println("print ar :", ar)
	fmt.Println("print a : ", a)

}
