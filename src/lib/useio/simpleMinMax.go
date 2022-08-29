package useio

import "fmt"

func UnameMinMax(x, y int) (int, int) {
	if x > y {
		umin := y
		umax := x
		return umin, umax
	} else {
		umin := x
		umax := y
		return umin, umax
	}

}

/*
func MinMax(x, y int) (mmin, mmax int) {
	if x > y {
		mmin := y
		mmax := x
	} else {
		mmin := x
		mmax := y
	}
	return mmin, mmax
}

func NamedMinMax(x, y int) (nmin, nmax int) {
	if x > y {
		nmin := y
		nmax := x
	} else {
		nmin := x
		nmax := y
	}
	return
}

func Sort(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}
*/

func A1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
	fmt.Println("Func A1 End")
}

func A2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println("Func A2 End")
}

func A3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
	fmt.Println("Func A3 End")
}
