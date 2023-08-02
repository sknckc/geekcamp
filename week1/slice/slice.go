package slice

import "fmt"

// 题目一，以整型类型切片进行解答
func SliceDel(a []int, n int) ([]int, int, error) {
	length := len(a)
	tar := a[n]
	if length <= n || length < 0 {
		return a, tar, fmt.Errorf("slice out of range,len is %d, index is %d", length, n)
	}
	init := []int{}

	init = append(init, a[:n]...)
	init = append(init, a[n+1:]...)
	return init, tar, nil
}

//题目二：使用拷贝法进行改进

func SliceDel2(a []int, n int) ([]int, int, error) {
	length := len(a)
	tar := a[n]
	if length <= n || length < 0 {
		return a, tar, fmt.Errorf("slice out of range,len is %d, index is %d", length, n)
	}
	j := 0
	for i, v := range a {

		if i != n {
			a[j] = v
			j++
		}
	}

	return a[:j], tar, nil
}

// 泛型
func SliceDel3[T any](a []T, n int) ([]T, T, error) {
	length := len(a)
	if length <= n || length < 0 {
		return a, a[n], fmt.Errorf("slice out of range,len is %d, index is %d", length, n)
	}
	tar := a[n]
	j := 0
	for i, v := range a {
		if i != n {
			a[j] = v
			j++
		}
	}
	return a[:j], tar, nil
}
