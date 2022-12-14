// File and Directory Library
package filedirec

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get pwd and Set Symlinks
func pwdd() {

	arguments := os.Args
	pwd, err := os.Getwd()

	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("Error : ", err)
	}

	if len(arguments) == 1 {
		return
	}
	if arguments[1] != "-P" {
		return
	}

	fileinfo, err := os.Lstat(pwd)

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}

func Pwd_use() {
	pwdd()
}
