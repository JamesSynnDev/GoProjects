package commandline

import (
	"flag"
	"fmt"
)

func Use_Flag() {
	minusO := flag.Bool("o", false, "o")
	minusC := flag.Bool("c", false, "c")
	minusK := flag.Int("k", 0, "an int")

	flag.Parse()

	fmt.Println("-o : ", *minusO)
	fmt.Println("-c : ", *minusC)
	fmt.Println("-k : ", *minusK)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)

	}
}

/*
go run useFlag.go
-o : false
-c : false
-k : 0

go run useFlag.go -o a b
-o : true
-c : false
-k : 0
0 : a
1 : b

./useFlag -k
flag needs an argument: -k
Usage of ./useFlag:
	-c c
	-k int
		an int
	-o o

./useFlag -k=abc
invalid value "abc" for flag -k: storconv.ParseInt: parsing "abc": invalid
syntax
Usage of ./useFlag:
	-c c
	-k int
		an int
	-o o

*/

// parse error I can use the 'Errorhandling'
