package slice

// Sum gets the ans of slice
func Sum(s []int) int {
	ret := 0
	for _, v := range s {
		ret += v
	}
	return ret
}

// SumAll calucates the ans of all slice
func SumAll(s1, s2 []int) int {
	ret := 0
	for _, v1 := range s1 {
		ret += v1
	}
	for _, v2 := range s2 {
		ret += v2
	}
	return ret
}

// SumToNew makes each slice ans to new slice
func SumToNew(s ...[]int) []int {
	var ret []int
	for i := 0; i < len(s); i++ {
		if len(s) == 0 {
			ret = append(ret, 0)
		} else {
			ret = append(ret, Sum(s[i]))
		}
	}
	return ret
}
