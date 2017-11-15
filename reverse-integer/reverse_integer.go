package reverse_integer

import "math"

func reverse(x int) int {
	arr := make([]int, 0)

	for ;x != 0; {
		arr = append(arr, x % 10)
		x = x / 10
	}

	res := 0
	j := 1
	l := len(arr)
	for k := l - 1; k >= 0; k-- {
		res += arr[k] * j
		j *= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}