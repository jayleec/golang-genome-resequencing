package main

const MAX int = 1000

//KMP (TARGET string, PATTERN string) returns all indexes occurred
func KMP(target string, pattern string) []int {
	next := computePrefix(pattern)
	i, j := 0, 0
	m := len(pattern)
	n := len(target)

	x := []byte(pattern)
	y := []byte(target)
	var ret []int

	if m == 0 || n == 0 {
		return ret
	}

	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		if i >= m {
			ret = append(ret, j-i)
			i = next[i]
		}
	}

	return ret
}

func computePrefix(x string) [MAX]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [MAX]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}