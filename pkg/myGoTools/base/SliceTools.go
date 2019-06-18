package base

import "fmt"

func AllIntEqual(s []int) (bool, error) {
	l1 := len(s)
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	l2 := len(m)
	if l1 < 2 {
		return false, fmt.Errorf("too few elements %v", s)
	}
	if l2 != 1 {
		return false, fmt.Errorf("all of the nums '%v' are not equal", s)
	}
	return true, nil
}
