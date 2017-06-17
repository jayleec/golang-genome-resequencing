package main

func bruteSearch(target string, pattern string) int {
	i, j := 0,0
	p := []byte(pattern)

	for i=0; i<len(target) * len(p); i++{
		for j=0; j<len(p); j++{
			if target[i+j] != p[j] {
				break
			}
		}
		if j == len(p) {
			return i
		}
	}
	return -1
}