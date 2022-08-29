package filedirec

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func which_simeple() {
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Please Provide an arguments!")
		os.Exit(1)
	}

	file := flags[0]
	fountIt := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, directory := range pathSlice {
		fullPath := directory + "/" + file
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode.IsRegular() {
				// Is it executable?
				if mode&0111 != 0 {
					fountIt = true
					// if the -s flag is set
					if *minusS == true {
						os.Exit(0)
					}
					// if the -a flag is set
					if *minusA == true {
						fmt.Println(fullPath)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}
	if fountIt == false {
		os.Exit(1)
	}

}

func Use_which() {
	which_simeple()
}
