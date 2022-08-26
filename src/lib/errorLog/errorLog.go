package errorLog

// https://golangbyexample.com/error-in-golang/XZZ

import (
	"fmt"
	"os"
)

func ErrorLog_1() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}
}
