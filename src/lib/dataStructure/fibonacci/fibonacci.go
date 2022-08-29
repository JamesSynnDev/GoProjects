package fibonacci

func Fibo_recursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return Fibo_recursion(n-1) + Fibo_recursion(n-2)
	}
}

func Fibo_simple(n int) int {
	return 0
}

func Fibo_tree(n int) int {
	return 0
}
