package basic

import "fmt"

func vari() {
	// 문자열 변수 선언 및 초기값 설정
	var s1, s2, s3 string = "first", "second", "third"

	// 여러변수 동시 초기화
	var (
		s = "mystring"
		i = 12
		f = 14.53
	)
	// Type Inference 타입 추론
	ss := "mystring"

	fmt.Println(s1, s2, s3)
	fmt.Println(s, i, f)
	fmt.Println(ss)

}
