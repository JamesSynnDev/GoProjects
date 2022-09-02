package monad

import (


)


func solution(_ arr:[Int]) -> [Int] {
	let min = arr.sort(by: <)[0]
	return arr.count == 1 ? [-1] : arr.compacMap({ return != min ? $0 : nil })

}